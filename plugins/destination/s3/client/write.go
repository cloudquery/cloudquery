package client

import (
	"bytes"
	"context"
	"io"
	"regexp"
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

var reInvalidJSONKey = regexp.MustCompile(`\W`)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data [][]any) error {
	name := strings.ReplaceAll(c.pluginSpec.Path, PathVarTable, table.Name)
	name = strings.ReplaceAll(name, PathVarUUID, uuid.NewString())

	if c.pluginSpec.Athena {
		for _, resource := range data {
			for u := range resource {
				if table.Columns[u].Type != schema.TypeJSON {
					continue
				}
				sanitizeJSONKeys(resource[u])
			}
		}
	}

	var b bytes.Buffer
	w := io.Writer(&b)

	if err := c.Client.WriteTableBatchFile(w, table, data); err != nil {
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

// sanitizeJSONKeys replaces all invalid characters in JSON keys with underscores.
// It does the replacement in-place, modifying the original object. This is required
// for compatibility with Athena.
func sanitizeJSONKeys(obj any) {
	switch m := obj.(type) {
	case map[string]any:
		for k, v := range m {
			nk := reInvalidJSONKey.ReplaceAllString(k, "_")
			// if a duplicate key is created by the replacement, it will be overwritten,
			// but we consider this highly unlikely
			delete(m, k)
			m[nk] = v
			sanitizeJSONKeys(v)
		}
	case []any:
		for _, v := range m {
			sanitizeJSONKeys(v)
		}
	}
}
