package client

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"path"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *Client) writeJSONResource(ctx context.Context, tableName string, resources <-chan []interface{}) error {
	filePath := path.Join(c.csvSpec.Directory, tableName+".json")
	var writer io.Writer
	var err error
	switch c.csvSpec.Backend {
	case BackendTypeLocal:
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return err
		}
		defer func() {
			err = f.Close()
		}()
		writer = f
	case BackendTypeS3:
		rPipe, wPipe := io.Pipe()
		if _, err := c.awsUploader.Upload(ctx, &s3.PutObjectInput{
			Bucket: &c.csvSpec.Directory,
			Key:    &tableName,
			Body:   rPipe,
		}); err != nil {
			return err
		}
		writer = wPipe
	case BackendTypeGCS:
		fileName := tableName + "_" + uid + ".csv"
		gcpWriter := c.gcpStorageClient.Bucket(c.csvSpec.Directory).Object(fileName).NewWriter(ctx)
		writer = gcpWriter
	default:
		panic("unknown backend type: " + c.csvSpec.Backend)
	}

	for r := range resources {
		b, err := json.Marshal(r)
		if err != nil {
			return err
		}
		b = append(b, '\n')
		if _, err := writer.Write(b); err != nil {
			return err
		}
	}


	return err
}
