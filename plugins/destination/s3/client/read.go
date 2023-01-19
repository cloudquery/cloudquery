package client

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/plugin-sdk/schema"
)

const maxFileSize = 1024 * 1024 * 20

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	if !c.pluginSpec.NoRotate {
		return fmt.Errorf("reading is not supported when no_rotate is false. Table: %q; Source: %q", table.Name, sourceName)
	}
	name := strings.ReplaceAll(c.pluginSpec.Path, PathVarTable, table.Name)
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
	return c.Client.Read(r, table, sourceName, res)
}
