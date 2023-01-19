package client

import (
	"bytes"
	"context"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/uuid"
)

const (
	PathVarTable = "{{TABLE}}"
	PathVarUUID  = "{{UUID}}"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data [][]any) error {
	name := strings.ReplaceAll(c.pluginSpec.Path, PathVarTable, table.Name)
	name = strings.ReplaceAll(name, PathVarUUID, uuid.NewString())

	var b bytes.Buffer
	w := io.Writer(&b)

	if err := c.filetype.WriteTableBatch(w, table, data); err != nil {
		return err
	}
	// we don't upload in parallel here because AWS sdk moves the burden to the developer, and
	// we don't want to deal with that yet. in the future maybe we can run some benchmarks and see if adding parallelization helps.
	r := io.Reader(&b)
	if _, err := c.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.pluginSpec.Bucket),
		Key:    aws.String(name),
		Body:   r,
	}); err != nil {
		return err
	}

	return nil
}
