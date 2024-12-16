package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"regexp"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	awstypes "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/uuid"
)

var reInvalidJSONKey = regexp.MustCompile(`\W`)

// If we don't use a reader that supports seeking, the S3 SDK will allocate a 5MB buffer each time it reads a chunk
// While this can work for large files, it's not optimal for small files, based on our tests we're mostly uploading small files
// And depending on source concurrency and destination batch settings, we can upload quite a bit of small files at the same time
// For example we've seen 220 concurrent uploads where most files are only a few dozen KB, this means we're allocating ~1GB of memory, where we could be allocating only a few MBs
// The memory is only cleared after the upload is finished, which makes it even worse
// Please note that for large files the memory we read is capped by the batch size setting which defaults to 50MB
func (c *Client) adjustReaderIfNeeded(r io.Reader) (io.Reader, error) {
	// We don't use a reader that supports seeking if NoRotate is set since that means we'll be holding the entire file in memory until the table has finished resolving and the file is uploaded.
	// This has the downside of potentially using a lot of memory for tables with small files.
	// But it's a tradeoff we're willing to make for now until https://github.com/aws/aws-sdk-go-v2/issues/2694 is resolved, as it has a workaround that we can use (reducing source concurrency).
	if c.spec.NoRotate {
		return r, nil
	}

	allData, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	readerSeeker := bytes.NewReader(allData)
	return readerSeeker, nil
}

func (c *Client) createObject(ctx context.Context, table *schema.Table, objKey string) (*filetypes.Stream, error) {
	s, err := c.Client.StartStream(table, func(r io.Reader) error {
		adjustedReader, err := c.adjustReaderIfNeeded(r)
		if err != nil {
			return err
		}
		params := &s3.PutObjectInput{
			Bucket:      aws.String(c.spec.Bucket),
			Key:         aws.String(objKey),
			Body:        adjustedReader,
			ContentType: aws.String(c.spec.GetContentType()),
		}

		if c.spec.ACL != "" {
			params.ACL = awstypes.ObjectCannedACL(c.spec.ACL)
		}

		sseConfiguration := c.spec.ServerSideEncryptionConfiguration
		if sseConfiguration != nil {
			params.SSEKMSKeyId = &sseConfiguration.SSEKMSKeyId
			params.ServerSideEncryption = sseConfiguration.ServerSideEncryption
		}

		_, err = manager.NewUploader(c.s3Client).Upload(ctx, params)
		return err
	})
	return s, err
}

func (c *Client) WriteTable(ctx context.Context, msgs <-chan *message.WriteInsert) error {
	var s *filetypes.Stream

	for msg := range msgs {
		if s == nil {
			table := msg.GetTable()
			objKey := c.spec.ReplacePathVariables(table.Name, uuid.NewString(), time.Now().UTC(), c.syncID)
			// if object was already initialized, use the same key
			c.initializedTablesLock.Lock()
			if val, ok := c.initializedTables[table.Name]; ok {
				objKey = val
				delete(c.initializedTables, table.Name)
			}
			c.initializedTablesLock.Unlock()

			var err error
			s, err = c.createObject(ctx, table, objKey)
			if err != nil {
				return err
			}
		}

		if c.spec.Athena {
			var err error
			msg.Record, err = sanitizeRecordJSONKeys(msg.Record)
			if err != nil {
				_ = s.FinishWithError(err)
				return err
			}
		}

		if err := s.Write([]arrow.Record{msg.Record}); err != nil {
			_ = s.FinishWithError(err)
			return err
		}
	}

	return s.Finish()
}

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	return c.writer.Write(ctx, msgs)
}

func (c *Client) MigrateTable(ctx context.Context, ch <-chan *message.WriteMigrateTable) error {
	for msg := range ch {
		if !c.spec.GenerateEmptyObjects {
			continue
		}
		table := msg.GetTable()
		objKey := c.spec.ReplacePathVariables(table.Name, uuid.NewString(), time.Now().UTC(), c.syncID)
		// We don't need any locking here because all messages for the same table are processed sequentially
		c.initializedTables[table.Name] = objKey
		s, err := c.createObject(ctx, table, objKey)
		if err != nil {
			return err
		}
		err = s.Finish()
		if err != nil {
			return err
		}
	}
	return nil
}

// sanitizeRecordJSONKeys replaces all invalid characters in JSON keys with underscores. This is required
// for compatibility with Athena.
func sanitizeRecordJSONKeys(record arrow.Record) (arrow.Record, error) {
	cols := make([]arrow.Array, record.NumCols())
	for i, col := range record.Columns() {
		if arrow.TypeEqual(col.DataType(), types.NewJSONType()) {
			b := types.NewJSONBuilder(memory.DefaultAllocator)
			for r := 0; r < int(record.NumRows()); r++ {
				if col.IsNull(r) {
					b.AppendNull()
					continue
				}
				data, err := sanitizeJSONRawMessage(col.GetOneForMarshal(r).(json.RawMessage))
				if err != nil {
					return nil, err
				}
				b.Append(data)
			}
			cols[i] = b.NewArray()
			continue
		}
		cols[i] = col
	}
	return array.NewRecord(record.Schema(), cols, record.NumRows()), nil
}

func sanitizeJSONRawMessage(rawMessage json.RawMessage) (any, error) {
	var data any
	err := json.Unmarshal(rawMessage, &data)
	if err != nil {
		return nil, err
	}

	return sanitizeJSONKeysForObject(data), nil
}

func sanitizeJSONKeysForObject(data any) any {
	// we only care about objects that have keys: present either while nesting or in arrays
	switch data := data.(type) {
	case map[string]any:
		res := make(map[string]any, len(data))
		for k, v := range data {
			res[reInvalidJSONKey.ReplaceAllString(k, "_")] = sanitizeJSONKeysForObject(v)
		}
		return res
	case []any:
		for i, el := range data {
			data[i] = sanitizeJSONKeysForObject(el)
		}
		return data
	default:
		return data
	}
}
