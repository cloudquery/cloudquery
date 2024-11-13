package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestSpec_JSONSchemaExtend(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true, // missing connection_string
			Spec: `{}`,
		},
		{
			Name: "extra keyword",
			Err:  true,
			Spec: `{"connection_string":"abc", "extra": true}`,
		},
		{
			Name: "empty connection_string",
			Err:  true,
			Spec: `{"connection_string":""}`,
		},
		{
			Name: "null connection_string",
			Err:  true,
			Spec: `{"connection_string":null}`,
		},
		{
			Name: "bad connection_string",
			Err:  true,
			Spec: `{"connection_string":123}`,
		},
		{
			Name: "proper connection_string",
			Spec: `{"connection_string":"abc"}`,
		},
		{
			Name: "empty cluster",
			Spec: `{"connection_string":"abc","cluster":""}`,
		},
		{
			Name: "null cluster",
			Err:  true,
			Spec: `{"connection_string":"abc","cluster":null}`,
		},
		{
			Name: "bad cluster",
			Err:  true,
			Spec: `{"connection_string":"abc","cluster":123}`,
		},
		{
			Name: "proper cluster",
			Spec: `{"connection_string":"abc","cluster":"abc"}`,
		},
		// Engine is tested separately, we will test only null here
		{
			Name: "null engine",
			Spec: `{"connection_string":"abc","engine":null}`,
		},
		{
			Name: "empty ca_cert",
			Spec: `{"connection_string":"abc","ca_cert":""}`,
		},
		{
			Name: "null ca_cert",
			Err:  true,
			Spec: `{"connection_string":"abc","ca_cert":null}`,
		},
		{
			Name: "bad ca_cert",
			Err:  true,
			Spec: `{"connection_string":"abc","ca_cert":123}`,
		},
		{
			Name: "proper ca_cert",
			Spec: `{"connection_string":"abc","ca_cert":"abc"}`,
		},
		{
			Name: "zero batch_size",
			Err:  true,
			Spec: `{"connection_string":"abc","batch_size":0}`,
		},
		{
			Name: "float batch_size",
			Err:  true,
			Spec: `{"connection_string":"abc","batch_size":5.3}`,
		},
		{
			Name: "bad batch_size",
			Err:  true,
			Spec: `{"connection_string":"abc","batch_size":false}`,
		},
		{
			Name: "null batch_size",
			Err:  true,
			Spec: `{"connection_string":"abc","batch_size":null}`,
		},
		{
			Name: "proper batch_size",
			Spec: `{"connection_string":"abc","batch_size":123}`,
		},
		{
			Name: "zero batch_size_bytes",
			Err:  true,
			Spec: `{"connection_string":"abc","batch_size_bytes":0}`,
		},
		{
			Name: "float batch_size_bytes",
			Err:  true,
			Spec: `{"connection_string":"abc","batch_size_bytes":5.3}`,
		},
		{
			Name: "bad batch_size_bytes",
			Err:  true,
			Spec: `{"connection_string":"abc","batch_size_bytes":false}`,
		},
		{
			Name: "null batch_size_bytes",
			Err:  true,
			Spec: `{"connection_string":"abc","batch_size_bytes":null}`,
		},
		{
			Name: "proper batch_size_bytes",
			Spec: `{"connection_string":"abc","batch_size_bytes":123}`,
		},
		// configtype.Duration is tested in plugin-sdk
		// test only null here
		{
			Name: "null batch_timeout",
			Spec: `{"connection_string":"abc","batch_timeout":null}`,
		},
	})
}

func TestSpec_ValidateEmptyPartitionBy(t *testing.T) {
	spec := Spec{Partition: []PartitionStrategy{{}}}
	spec.SetDefaults()
	err := spec.Validate()
	require.Error(t, err)
	require.ErrorContains(t, err, "partition_by is required")
}

func TestSpec_ValidateEmptyTables(t *testing.T) {
	spec := Spec{Partition: []PartitionStrategy{{PartitionBy: "test_field"}}}
	spec.SetDefaults()
	require.NoError(t, spec.Validate())
	require.Equal(t, []string{"*"}, spec.Partition[0].Tables)
}
