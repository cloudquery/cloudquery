package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/filetypes/csv"
	"github.com/cloudquery/filetypes/json"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/uuid"
)

const (
	PathVarTable = "{{TABLE}}"
	PathVarUUID  = "{{UUID}}"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data [][]any) error {
	path := c.pluginSpec.Path
	if !strings.Contains(path, PathVarTable) {
		// for backwards-compatibility, default to given path plus /{{TABLE}}.[format].{{UUID}} if
		// no {{TABLE}} value is found in the path string
		path += fmt.Sprintf("/%s.%s", PathVarTable, c.pluginSpec.Format)
		if !c.pluginSpec.NoRotate {
			path += "." + PathVarUUID
		}
	}
	name := strings.ReplaceAll(path, PathVarTable, table.Name)
	if !c.pluginSpec.NoRotate {
		name = strings.ReplaceAll(name, PathVarUUID, uuid.NewString())
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
