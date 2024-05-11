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
	"path": "{{UUID}}",
	"bucket": "b",
	"region": "r"
}`,
		},
		{
			Name: "bad format value", // also a part of embedded FileSpec testing
			Spec: `{"format": "cs22v", "path": "{{UUID}}", "bucket": "b", "region": "r"}`,
			Err:  true,
		},
		{
			Name: "minimal",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r"}`,
		},
		{
			Name: "missing path",
			Spec: `{"format": "csv", "bucket": "b", "region": "r"}`,
			Err:  true,
		},
		{
			Name: "empty path",
			Spec: `{"format": "csv", "path": ", ", "bucket": "b", "region": "r"}`,
			Err:  true,
		},
		{
			Name: "null path",
			Spec: `{"format": "csv", "path": null, "bucket": "b", "region": "r"}`,
			Err:  true,
		},
		{
			Name: "integer path",
			Spec: `{"format": "csv", "path": 123}`,
			Err:  true,
		},
		{
			Name: "path starts with /",
			Spec: `{"format": "csv", "path": "/{{UUID}}", "bucket": "b", "region": "r"}`,
			Err:  true,
		},
		{
			Name: "path contains //",
			Spec: `{"format": "csv", "path": "{{UUID}}//", "bucket": "b", "region": "r"}`,
			Err:  true,
		},
		{
			Name: "path contains ./",
			Spec: `{"format": "csv", "path": "{{UUID}}/./", "bucket": "b", "region": "r"}`,
			Err:  true,
		},
		{
			Name: "path contains ../",
			Spec: `{"format": "csv", "path": "{{UUID}}/../", "bucket": "b", "region": "r"}`,
			Err:  true,
		},
		{
			Name: "missing bucket",
			Spec: `{"format": "csv", "path": "{{UUID}}", "region": "r"}`,
			Err:  true,
		},
		{
			Name: "empty bucket",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "", "region": "r"}`,
			Err:  true,
		},
		{
			Name: "null bucket",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": null, "region": "r"}`,
			Err:  true,
		},
		{
			Name: "integer bucket",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": 123, "region": "r"}`,
			Err:  true,
		},
		{
			Name: "missing region",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b"}`,
			Err:  true,
		},
		{
			Name: "empty region",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": ""}`,
			Err:  true,
		},
		{
			Name: "null region",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": null}`,
			Err:  true,
		},
		{
			Name: "integer region",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": 123}`,
			Err:  true,
		},
		{
			Name: "null no_rotate",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "no_rotate": null}`,
			Err:  true,
		},
		{
			Name: "bad no_rotate",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "no_rotate": 123}`,
			Err:  true,
		},
		{
			Name: "no_rotate:true",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "region": "r", "no_rotate": true}`,
		},
		{
			Name: "no_rotate:false",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "no_rotate": false}`,
		},
		{
			Name: "zero batch_size",
			Err:  true,
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "batch_size":0}`,
		},
		{
			Name: "float batch_size",
			Err:  true,
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "batch_size":5.3}`,
		},
		{
			Name: "bad batch_size",
			Err:  true,
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "batch_size":false}`,
		},
		{
			Name: "null batch_size",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "batch_size":null}`,
		},
		{
			Name: "proper batch_size",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "batch_size":123}`,
		},
		{
			Name: "zero batch_size_bytes",
			Err:  true,
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "batch_size_bytes":0}`,
		},
		{
			Name: "float batch_size_bytes",
			Err:  true,
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "batch_size_bytes":5.3}`,
		},
		{
			Name: "bad batch_size_bytes",
			Err:  true,
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "batch_size_bytes":false}`,
		},
		{
			Name: "null batch_size_bytes",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "batch_size_bytes":null}`,
		},
		{
			Name: "proper batch_size_bytes",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "batch_size_bytes":123}`,
		},
		// configtype.Duration is tested in plugin-sdk
		// test only null here
		{
			Name: "null batch_timeout",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "batch_timeout":null}`,
		},

		// no_rotate + path({{UUID}})
		{
			Name: "no_rotate:false & path:{{UUID}}",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "no_rotate":false}`,
		},
		{
			Name: "no_rotate:true & path:{{UUID}}",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "no_rotate":true}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & path:abc",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "region": "r", "no_rotate":false}`,
			Err:  true,
		},
		{
			Name: "no_rotate:true & path:abc",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "region": "r", "no_rotate":true}`,
		},

		// no_rotate + batching
		{
			Name: "no_rotate:false & batch_size:100",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "no_rotate":false, "batch_size":100}`,
		},
		{
			Name: "no_rotate:true & batch_size:100",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "region": "r", "no_rotate":true, "batch_size":100}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & batch_size:null",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "no_rotate":false, "batch_size":null}`,
		},
		{
			Name: "no_rotate:true & batch_size:null",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "region": "r", "no_rotate":true, "batch_size":null}`,
		},
		{
			Name: "no_rotate:false & batch_size_bytes:100",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "no_rotate":false, "batch_size_bytes":100}`,
		},
		{
			Name: "no_rotate:true & batch_size_bytes:100",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "region": "r", "no_rotate":true, "batch_size_bytes":100}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & batch_size_bytes:null",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "no_rotate":false, "batch_size_bytes":null}`,
		},
		{
			Name: "no_rotate:true & batch_size_bytes:null",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "region": "r", "no_rotate":true, "batch_size_bytes":null}`,
		},
		{
			Name: "no_rotate:false & batch_timeout:100s",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "no_rotate":false, "batch_timeout":"100s"}`,
		},
		{
			Name: "no_rotate:true & batch_timeout:100s",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "region": "r", "no_rotate":true, "batch_timeout":"100s"}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & batch_timeout:null",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "no_rotate":false, "batch_timeout":null}`,
		},
		{
			Name: "no_rotate:true & batch_timeout:null",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "region": "r", "no_rotate":true, "batch_timeout":null}`,
		},

		// batching + path({{UUID}})
		{
			Name: "batching (no_rotate:false) & path:{{UUID}}",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "no_rotate":false}`,
		},
		{
			Name: "batching (no_rotate:false) & path:abc",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "region": "r", "no_rotate":false}`,
			Err:  true,
		},
		{
			Name: "batching (missing no_rotate) & path:{{UUID}}",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r"}`,
		},
		{
			Name: "batching (missing no_rotate) & path:abc",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "region": "r"}`,
			Err:  true,
		},
		{
			Name: "no batching (no_rotate:true) & path:{{UUID}}",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "no_rotate":true}`,
			Err:  true,
		},
		{
			Name: "no batching (no_rotate:true) & path:abc",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "region": "r", "no_rotate":true}`,
		},
		{
			Name: "server side encryption (empty)",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "server_side_encryption_configuration": {}}`,
			Err:  true,
		},
		{
			Name: "server side encryption (missing type)",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "server_side_encryption_configuration": {"sse_kms_key_id":"1234-5678"}}`,
			Err:  true,
		},
		{
			Name: "server side encryption (missing key)",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "server_side_encryption_configuration": {"server_side_encryption":"AES256"}}`,
			Err:  true,
		},
		{
			Name: "server side encryption (success)",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r", "server_side_encryption_configuration": {"server_side_encryption":"AES256", "sse_kms_key_id":"1234-5678"}}`,
		},
		{
			Name: "Empty objects false format parquet",
			Spec: `{"format": "parquet", "path": "{{UUID}}", "bucket": "b", "region": "r","write_empty_objects_for_empty_tables":false}`,
		},
		{
			Name: "Empty objects false format csv",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r","write_empty_objects_for_empty_tables":false}`,
		},
		{
			Name: "Empty objects true",
			Spec: `{"format": "parquet", "path": "{{UUID}}", "bucket": "b", "region": "r","write_empty_objects_for_empty_tables":true}`,
		},
		{
			Name: "Empty objects true format must be parquet",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "region": "r","write_empty_objects_for_empty_tables":true}`,
			Err:  true,
		},
	})
}
