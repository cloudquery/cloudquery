package client

import (
	"bytes"
	"context"
	"io"
	"path"
	"reflect"
	"regexp"
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

var reInvalidJSONKey = regexp.MustCompile(`\W`)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, data [][]any) error {
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
		Key:    aws.String(replacePathVariables(c.pluginSpec.Path, table.Name, uuid.NewString())),
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
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			k := iter.Key()
			if k.Kind() == reflect.String {
				nk := reInvalidJSONKey.ReplaceAllString(k.String(), "_")
				v := iter.Value()
				sanitizeJSONKeys(v.Interface())
				value.SetMapIndex(k, reflect.Value{})
				value.SetMapIndex(reflect.ValueOf(nk), v)
			}
		}
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			sanitizeJSONKeys(value.Index(i).Interface())
		}
	}
}

func replacePathVariables(specPath, table, fileIdentifier string) string {
	name := strings.ReplaceAll(specPath, PathVarTable, table)
	name = strings.ReplaceAll(name, PathVarUUID, fileIdentifier)
	return path.Clean(name)
}
