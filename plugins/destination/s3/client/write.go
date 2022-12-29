package client

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/filetypes/csv"
	"github.com/cloudquery/filetypes/json"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/uuid"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data [][]any) error {
	name := fmt.Sprintf("%s/%s.%s.%s", c.pluginSpec.Path, table.Name, c.pluginSpec.Format, uuid.NewString())
	if c.pluginSpec.NoRotate {
		name = fmt.Sprintf("%s/%s.%s", c.pluginSpec.Path, table.Name, c.pluginSpec.Format)
	}
	var b bytes.Buffer
	w := io.Writer(&b)
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		if err := csv.WriteTableBatch(w, table, data); err != nil {
			return err
		}
	case FormatTypeJSON:
		if err := json.WriteTableBatch(w, table, data); err != nil {
			return err
		}
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
	// we don't do that in parallel here because aws sdk moves the burden to the devleoper and
	// we don't want to deal with that yet. in the future maybe we can run some benchmarks and see if adding parralization helps.
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
