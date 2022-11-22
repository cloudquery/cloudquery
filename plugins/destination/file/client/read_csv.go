package client

import (
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) readCSV(ctx context.Context, table *schema.Table, sourceName string, res chan<- []interface{}) error {
	var reader io.Reader
	var err error

	switch c.csvSpec.Backend {
	case BackendTypeLocal:
		filePath := path.Join(c.csvSpec.Directory, table.Name+".csv")
		f, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer f.Close()
		reader = f
	case BackendTypeS3:
		// This is AWS trying to be smart and download file in parallel and drop it all to the developer instead
		// of exposing normal reader interface.
		writer := manager.NewWriteAtBuffer(make([]byte, 1024))
		c.awsDownloader.Download(ctx, writer, &s3.GetObjectInput{
			Bucket: aws.String(c.bucket),
			Key:    aws.String(c.dir + "/" + table.Name + ".csv"),
		})
		reader = bytes.NewReader(writer.Bytes())
	case BackendTypeGCS:
		reader, err = c.gcpStorageClient.Bucket(c.bucket).Object(c.dir + "/" + table.Name + ".csv").NewReader(ctx)
		if err != nil {
			return err
		}
	default:
		panic("unknown backend type " + c.csvSpec.Backend)
	}


	r := csv.NewReader(reader)
	sourceNameIndex := table.Columns.Index(schema.CqSourceNameColumn.Name)
	if sourceNameIndex == -1 {
		return fmt.Errorf("could not find column %s in table %s", schema.CqSourceNameColumn.Name, table.Name)
	}

	for {
		record, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		if record[sourceNameIndex] != sourceName {
			continue
		}
		values := make([]interface{}, len(record))
		for i, v := range record {
			values[i] = v
		}

		res <- values
	}
	return nil
}
