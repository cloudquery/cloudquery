package client

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func (c *Client) writeRemote(ctx context.Context, tableName string, content []byte) error {
	uid := uuid.New().String()
	switch c.csvSpec.Backend {
	case BackendTypeGCS:
		fileName := tableName + "_" + uid + ".csv"
		gcpWriter := c.gcpStorageClient.Bucket(c.csvSpec.Directory).Object(fileName).NewWriter(ctx)
		gcpWriter.Write(content)
		gcpWriter.Close()
	case BackendTypeS3:
		if _, err := c.awsUploader.Upload(ctx, &s3.PutObjectInput{
			Bucket: &c.csvSpec.Directory,
			Key:    &tableName,
			Body:   bytes.NewReader(content),
		}); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown backend type: %s", c.csvSpec.Backend)
	}
	return nil
}
