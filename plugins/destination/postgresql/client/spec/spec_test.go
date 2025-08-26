package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestSpec_JSONSchemaExtend(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "missing connection_string",
			Spec: `{}`,
			Err:  true,
		},
		{
			Name: "empty connection_string",
			Spec: `{"connection_string": ""}`,
			Err:  true,
		},
		{
			Name: "null connection_string",
			Spec: `{"connection_string": null}`,
			Err:  true,
		},
		{
			Name: "integer connection_string",
			Spec: `{"connection_string": 123}`,
			Err:  true,
		},
		{
			Name: "non-empty connection_string",
			Spec: `{"connection_string": "abc"}`,
		},
		// pgx_log_level is tested separately, just test null here
		{
			Name: "null pgx_log_level",
			Spec: `{"connection_string": "abc", "pgx_log_level": null}`,
			Err:  true,
		},
		{
			Name: "zero batch_size",
			Spec: `{"connection_string": "abc", "batch_size": 0}`,
			Err:  true,
		},
		{
			Name: "negative batch_size",
			Spec: `{"connection_string": "abc", "batch_size": -1}`,
			Err:  true,
		},
		{
			Name: "float batch_size",
			Spec: `{"connection_string": "abc", "batch_size": 1.5}`,
			Err:  true,
		},
		{
			Name: "null batch_size",
			Spec: `{"connection_string": "abc", "batch_size": null}`,
			Err:  true,
		},
		{
			Name: "string batch_size",
			Spec: `{"connection_string": "abc", "batch_size": "123"}`,
			Err:  true,
		},
		{
			Name: "proper batch_size",
			Spec: `{"connection_string": "abc", "batch_size": 123}`,
		},
		{
			Name: "zero batch_size_bytes",
			Spec: `{"connection_string": "abc", "batch_size_bytes": 0}`,
			Err:  true,
		},
		{
			Name: "negative batch_size_bytes",
			Spec: `{"connection_string": "abc", "batch_size_bytes": -1}`,
			Err:  true,
		},
		{
			Name: "float batch_size_bytes",
			Spec: `{"connection_string": "abc", "batch_size_bytes": 1.5}`,
			Err:  true,
		},
		{
			Name: "null batch_size_bytes",
			Spec: `{"connection_string": "abc", "batch_size_bytes": null}`,
			Err:  true,
		},
		{
			Name: "string batch_size_bytes",
			Spec: `{"connection_string": "abc", "batch_size_bytes": "123"}`,
			Err:  true,
		},
		{
			Name: "proper batch_size_bytes",
			Spec: `{"connection_string": "abc", "batch_size_bytes": 123}`,
		},
		// batch_timeout is tested in configtype package, test only null & empty here
		{
			Name: "empty batch_timeout",
			Spec: `{"connection_string": "abc", "batch_timeout": ""}`,
			Err:  true,
		},
		{
			Name: "null batch_timeout",
			Spec: `{"connection_string": "abc", "batch_timeout": null}`,
			Err:  true,
		},
		// minimal valid with pgvector_config provided and complete
		{
			Name: "pgvector minimal",
			Spec: `{
				"connection_string":"abc",
				"pgvector_config":{
					"tables":[{"table_name":"box_file_contents","embed_columns":["content"],"metadata_columns":["id"]}],
					"embedding":{"dimensions":1536,"api_key":"k","model_name":"text-embedding-3-small"}
				}
			}`,
		},
	})
}
