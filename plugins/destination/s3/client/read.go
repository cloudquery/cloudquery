package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
)

const maxFileSize = 1024 * 1024 * 20

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	if !c.spec.NoRotate {
		return fmt.Errorf("reading is not supported when no_rotate is false. Table: %q", table.Name)
	}
	if c.spec.PathContainsUUID() {
		return fmt.Errorf("reading is not supported when path contains uuid variable. Table: %q", table.Name)
	}

	name := c.spec.ReplacePathVariables(table.Name, uuid.NewString(), time.Time{}, c.syncID)
	writerAtBuffer := manager.NewWriteAtBuffer(make([]byte, 0, maxFileSize))
	_, err := manager.NewDownloader(c.s3Client).Download(ctx,
		writerAtBuffer,
		&s3.GetObjectInput{
			Bucket: aws.String(c.spec.Bucket),
			Key:    aws.String(name),
		})

	if err != nil {
		var smithyError *smithy.OperationError
		if errors.As(err, &smithyError); smithyError != nil {
			var httpError *awshttp.ResponseError
			if errors.As(smithyError.Err, &httpError); httpError != nil {
				if httpError.HTTPStatusCode() == http.StatusNotFound {
					return nil
				}
			}
		}
		return err
	}
	r := bytes.NewReader(writerAtBuffer.Bytes())
	return c.Client.Read(r, table, res)
}
