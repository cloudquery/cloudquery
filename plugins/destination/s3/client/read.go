package client

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

const maxFileSize = 1024 * 1024 * 20

func (c *Client) Read(ctx context.Context, arrowSchema *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	tableName := schema.TableName(arrowSchema)
	if !c.pluginSpec.NoRotate {
		return fmt.Errorf("reading is not supported when no_rotate is false. Table: %q; Source: %q", tableName, sourceName)
	}
	if strings.Contains(c.pluginSpec.Path, PathVarUUID) {
		return fmt.Errorf("reading is not supported when path contains uuid variable. Table: %q; Source: %q", tableName, sourceName)
	}
	name := strings.ReplaceAll(c.pluginSpec.Path, PathVarTable, tableName)
	writerAtBuffer := manager.NewWriteAtBuffer(make([]byte, 0, maxFileSize))
	_, err := c.downloader.Download(ctx,
		writerAtBuffer,
		&s3.GetObjectInput{
			Bucket: aws.String(c.pluginSpec.Bucket),
			Key:    aws.String(name),
		})
	if err != nil {
		return err
	}
	r := bytes.NewReader(writerAtBuffer.Bytes())
	return c.Client.Read(r, arrowSchema, sourceName, res)
}
