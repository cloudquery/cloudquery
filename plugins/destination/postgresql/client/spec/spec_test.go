package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
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
					"tables":[{"source_table_name":"box_file_contents","target_table_name":"box_file_contents_embeddings","embed_columns":["content"],"metadata_columns":["id"]}],
					"openai_embedding":{"dimensions":1536,"api_key":"k","model_name":"text-embedding-3-small"}
				}
			}`,
		},
		{
			Name: "lakebase minimal",
			Spec: `{"connection_string":"abc","lakebase":{"endpoint":"projects/p/branches/b/endpoints/e"}}`,
		},
		{
			Name: "lakebase full",
			Spec: `{"connection_string":"abc","lakebase":{"endpoint":"projects/p/branches/b/endpoints/e","host":"https://example.cloud.databricks.com","client_id":"id","client_secret":"secret"}}`,
		},
		{
			Name: "lakebase missing endpoint",
			Spec: `{"connection_string":"abc","lakebase":{"host":"https://example.cloud.databricks.com"}}`,
			Err:  true,
		},
		{
			Name: "lakebase empty endpoint",
			Spec: `{"connection_string":"abc","lakebase":{"endpoint":""}}`,
			Err:  true,
		},
		{
			Name: "lakebase unknown field",
			Spec: `{"connection_string":"abc","lakebase":{"endpoint":"projects/p/branches/b/endpoints/e","unknown":"x"}}`,
			Err:  true,
		},
		{
			Name: "aws_iam_auth minimal",
			Spec: `{"connection_string":"abc","aws_iam_auth":{}}`,
		},
		{
			Name: "aws_iam_auth full",
			Spec: `{"connection_string":"abc","aws_iam_auth":{"service":"rds","region":"us-east-1","endpoint":"mydb.123456789012.us-east-1.rds.amazonaws.com:5432","local_profile":"my_profile","role_arn":"arn:aws:iam::123456789012:role/my-role","role_session_name":"cloudquery","external_id":"external-id"}}`,
		},
		{
			Name: "aws_iam_auth unsupported service",
			Spec: `{"connection_string":"abc","aws_iam_auth":{"service":"dsql"}}`,
			Err:  true,
		},
		{
			Name: "aws_iam_auth invalid role_arn",
			Spec: `{"connection_string":"abc","aws_iam_auth":{"role_arn":"not-an-arn"}}`,
			Err:  true,
		},
		{
			Name: "aws_iam_auth unknown field",
			Spec: `{"connection_string":"abc","aws_iam_auth":{"unknown":"x"}}`,
			Err:  true,
		},
		{
			Name: "proper write_concurrency",
			Spec: `{"connection_string": "abc", "write_concurrency": 8}`,
		},
		{
			Name: "zero write_concurrency",
			Spec: `{"connection_string": "abc", "write_concurrency": 0}`,
			Err:  true,
		},
		{
			Name: "negative write_concurrency",
			Spec: `{"connection_string": "abc", "write_concurrency": -1}`,
			Err:  true,
		},
		{
			Name: "float write_concurrency",
			Spec: `{"connection_string": "abc", "write_concurrency": 1.5}`,
			Err:  true,
		},
		{
			Name: "null write_concurrency",
			Spec: `{"connection_string": "abc", "write_concurrency": null}`,
			Err:  true,
		},
		{
			Name: "string write_concurrency",
			Spec: `{"connection_string": "abc", "write_concurrency": "8"}`,
			Err:  true,
		},
		{
			Name: "proper use_copy_from",
			Spec: `{"connection_string": "abc", "use_copy_from": true}`,
		},
		{
			Name: "string use_copy_from",
			Spec: `{"connection_string": "abc", "use_copy_from": "true"}`,
			Err:  true,
		},
	})
}

func TestSpec_WriteConcurrencyDefaults(t *testing.T) {
	for _, tt := range []struct {
		name string
		spec Spec
		want int64
	}{
		{name: "unset", spec: Spec{}, want: defaultWriteConcurrency},
		{name: "zero", spec: Spec{WriteConcurrency: 0}, want: defaultWriteConcurrency},
		{name: "negative", spec: Spec{WriteConcurrency: -4}, want: defaultWriteConcurrency},
		{name: "set", spec: Spec{WriteConcurrency: 16}, want: 16},
	} {
		t.Run(tt.name, func(t *testing.T) {
			tt.spec.SetDefaults()
			require.Equal(t, tt.want, tt.spec.WriteConcurrency)
		})
	}
}
