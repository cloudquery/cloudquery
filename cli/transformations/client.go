package transformations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/ipc"
)

type TransformationClient struct {
	initialSchema *arrow.Schema
	finalSchema   *arrow.Schema
	queryBuilder  *queryBuilder
}

func NewTransformationClient() *TransformationClient {
	c := &TransformationClient{queryBuilder: newQueryBuilder()}
	c.mustEnsureServerRunning()
	return c
}

func (c *TransformationClient) SetInitialSchema(schema *arrow.Schema) {
	if c.initialSchema != nil {
		panic("Cannot set the initial schema more than once")
	}
	if schema == nil {
		panic("Cannot set the initial schema to nil")
	}
	c.initialSchema = schema
	c.finalSchema = schema
}

func (c *TransformationClient) AddTransformations(transformations []Transformation) error {
	if c.finalSchema == nil {
		return fmt.Errorf("Cannot add transformation without setting the initial schema")
	}
	for _, transformation := range transformations {
		sql, err := transformation.SQLFn(c.finalSchema)
		if err != nil {
			return err
		}
		c.queryBuilder.add(sql)
		c.finalSchema, err = transformation.SchemaChangeFn(c.finalSchema)
		if err != nil {
			return err
		}
	}
	for _, sql := range c.queryBuilder.build() {
		err := c.registerSQL(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *TransformationClient) GetTransformedSchema() *arrow.Schema {
	return c.finalSchema
}

func (c *TransformationClient) GetInitialSchema() *arrow.Schema {
	return c.initialSchema
}

func (c *TransformationClient) TransformRecord(batchRecord arrow.Record) (resultRecord arrow.Record, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic occurred applying the expected schema to a record: %v", r)
		}
	}()

	// Send the serialized data to the Python server
	resp, err := c.sendPost("manipulate", "application/octet-stream", serializeRecord(batchRecord))
	if err != nil {
		return nil, err
	}

	var (
		queryResultRecord arrow.Record
		arrays            []arrow.Array
	)
	reader, err := ipc.NewReader(bytes.NewReader(resp))
	if err != nil {
		return nil, err
	}
	for reader.Next() {
		queryResultRecord = reader.Record()
		break
	}
	// A transformation may add a WHERE clause that filters out all rows
	if queryResultRecord == nil {
		return nil, nil
	}

	arrays = recordToArrays(queryResultRecord)

	resultRecord = array.NewRecord(c.finalSchema, arrays, queryResultRecord.NumRows())
	return resultRecord, nil
}

func (c *TransformationClient) mustEnsureServerRunning() {
	_, err := c.sendPost("restart", "application/json", bytes.NewBuffer(nil))
	if err != nil {
		panic(err)
	}
}

func (TransformationClient) sendPost(endpoint string, contentType string, payload io.Reader) ([]byte, error) {
	// Send the POST request to the server
	resp, err := http.Post(fmt.Sprintf("http://localhost:5000/%v", endpoint), contentType, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Server returned HTTP %v", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func (c *TransformationClient) registerSQL(sql string) error {
	payload, _ := json.Marshal(map[string]string{
		"query": sql,
	})
	_, err := c.sendPost("query", "application/json", bytes.NewBuffer(payload))
	return err
}

func serializeRecord(batchRecord arrow.Record) io.Reader {
	buf := new(bytes.Buffer)
	writer := ipc.NewWriter(buf, ipc.WithSchema(batchRecord.Schema()))
	if err := writer.Write(batchRecord); err != nil {
		log.Fatal(err)
	}
	writer.Close()
	return buf
}

func recordToArrays(record arrow.Record) []arrow.Array {
	numCols := int(record.NumCols())
	arrays := make([]arrow.Array, numCols)
	for i := 0; i < numCols; i++ {
		arrays[i] = record.Column(i)
	}
	return arrays
}
