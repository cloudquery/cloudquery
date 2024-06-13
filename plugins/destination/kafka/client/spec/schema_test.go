package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestSpecJSONSchema(t *testing.T) {
	// cases about embedded filetypes.FileSpec are tested in the corresponding package
	// However, we add some tests to verify that it actually is properly working
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "csv file spec",
			Spec: `{
	"format": "csv",
	"format_spec": {
		"skip_header": true,
		"delimiter": "#"
	},
	"brokers": ["abc"]
}`,
		},
		{
			Name: "bad format value", // also a part of embedded FileSpec testing
			Spec: `{"format": "cs22v", "brokers": ["abc"]}`,
			Err:  true,
		},
		{
			Name: "missing brokers",
			Spec: `{"format": "csv"}`,
			Err:  true,
		},
		{
			Name: "empty brokers",
			Spec: `{"format": "csv", "brokers": []}`,
			Err:  true,
		},
		{
			Name: "null brokers",
			Spec: `{"format": "csv", "brokers": null}`,
			Err:  true,
		},
		{
			Name: "integer brokers",
			Spec: `{"format": "csv", "brokers": 123}`,
			Err:  true,
		},
		{
			Name: "empty brokers value",
			Spec: `{"format": "csv", "brokers": [""]}`,
			Err:  true,
		},
		{
			Name: "null brokers value",
			Spec: `{"format": "csv", "brokers": [null]}`,
			Err:  true,
		},
		{
			Name: "integer brokers value",
			Spec: `{"format": "csv", "brokers": [123]}`,
			Err:  true,
		},
		{
			Name: "proper brokers",
			Spec: `{"format": "csv", "brokers": ["abc"]}`,
		},
		{
			Name: "null verbose",
			Spec: `{"format": "csv", "brokers": ["abc"], "verbose": null}`,
			Err:  true,
		},
		{
			Name: "integer verbose",
			Spec: `{"format": "csv", "brokers": ["abc"], "verbose": 123}`,
			Err:  true,
		},
		{
			Name: "verbose:true",
			Spec: `{"format": "csv", "brokers": ["abc"], "verbose": true}`,
		},
		{
			Name: "verbose:false",
			Spec: `{"format": "csv", "brokers": ["abc"], "verbose": false}`,
		},
		// sasl_username & sasl_password have to go together
		{
			Name: "empty sasl_username with empty sasl_password",
			Spec: `{"format": "csv", "brokers": ["abc"], "sasl_username": "", "sasl_password": ""}`,
		},
		{
			Name: "non-empty sasl_username with non-empty sasl_password",
			Spec: `{"format": "csv", "brokers": ["abc"], "sasl_username": "user", "sasl_password": "password"}`,
		},
		{
			Name: "non-empty sasl_username without sasl_password",
			Spec: `{"format": "csv", "brokers": ["abc"], "sasl_username": "user"}`,
			Err:  true,
		},
		{
			Name: "non-empty sasl_username with empty sasl_password",
			Spec: `{"format": "csv", "brokers": ["abc"], "sasl_username": "user", "sasl_password": ""}`,
			Err:  true,
		},
		{
			Name: "non-empty sasl_username with integer sasl_password",
			Spec: `{"format": "csv", "brokers": ["abc"], "sasl_username": "user", "sasl_password": 123}`,
			Err:  true,
		},
		{
			Name: "non-empty sasl_username with null sasl_password",
			Spec: `{"format": "csv", "brokers": ["abc"], "sasl_username": "user", "sasl_password": null}`,
			Err:  true,
		},
		{
			Name: "non-empty sasl_password without sasl_username",
			Spec: `{"format": "csv", "brokers": ["abc"], "sasl_password": "password"}`,
			Err:  true,
		},
		{
			Name: "non-empty sasl_password with empty sasl_username",
			Spec: `{"format": "csv", "brokers": ["abc"], "sasl_username": "", "sasl_password": "password"}`,
			Err:  true,
		},
		{
			Name: "non-empty sasl_password with integer sasl_username",
			Spec: `{"format": "csv", "brokers": ["abc"], "sasl_username": 123, "sasl_password": "password"}`,
			Err:  true,
		},
		{
			Name: "non-empty sasl_password with null sasl_username",
			Spec: `{"format": "csv", "brokers": ["abc"], "sasl_username": null, "sasl_password": "password"}`,
			Err:  true,
		},
		{
			Name: "zero batch_size",
			Spec: `{"format": "csv", "brokers": ["abc"], "batch_size": 0}`,
			Err:  true,
		},
		{
			Name: "float batch_size",
			Spec: `{"format": "csv", "brokers": ["abc"], "batch_size": 1.5}`,
			Err:  true,
		},
		{
			Name: "negative batch_size",
			Spec: `{"format": "csv", "brokers": ["abc"], "batch_size": -1}`,
			Err:  true,
		},
		{
			Name: "null batch_size",
			Spec: `{"format": "csv", "brokers": ["abc"], "batch_size": null}`,
			Err:  true,
		},
		{
			Name: "string batch_size",
			Spec: `{"format": "csv", "brokers": ["abc"], "batch_size": "abc"}`,
			Err:  true,
		},
		{
			Name: "proper batch_size",
			Spec: `{"format": "csv", "brokers": ["abc"], "batch_size": 100}`,
		},

		{
			Name: "proper num_partitions",
			Spec: `{"format": "csv", "brokers": ["abc"], "topic_details": {"num_partitions": 10}}`,
		},

		{
			Name: "proper replication_factor",
			Spec: `{"format": "csv", "brokers": ["abc"], "topic_details": {"replication_factor": 10}}`,
		},

		{
			Name: "proper replication_factor and num_partitions",
			Spec: `{"format": "csv", "brokers": ["abc"], "topic_details": {"num_partitions": 10, "replication_factor": 10}}`,
		},
	})
}
