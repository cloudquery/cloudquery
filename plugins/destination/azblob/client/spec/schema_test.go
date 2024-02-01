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
	"path": "abc",
    "storage_account": "sa",
    "container": "c"
}`,
		},
		{
			Name: "bad format value", // also a part of embedded FileSpec testing
			Spec: `{"format": "cs22v", "path": "abc"}`,
			Err:  true,
		},
		{
			Name: "minimal",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c"}`,
		},
		{
			Name: "missing path",
			Spec: `{"format": "csv", "storage_account": "sa", "container": "c"}`,
			Err:  true,
		},
		{
			Name: "empty path",
			Spec: `{"format": "csv", "path": "", "storage_account": "sa", "container": "c"}`,
			Err:  true,
		},
		{
			Name: "null path",
			Spec: `{"format": "csv", "path": null, "storage_account": "sa", "container": "c"}`,
			Err:  true,
		},
		{
			Name: "integer path",
			Spec: `{"format": "csv", "path": 123, "storage_account": "sa", "container": "c"}`,
			Err:  true,
		},
		{
			Name: "missing storage_account",
			Spec: `{"format": "csv", "path": "abc", "container": "c"}`,
			Err:  true,
		},
		{
			Name: "empty storage_account",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "", "container": "c"}`,
			Err:  true,
		},
		{
			Name: "null storage_account",
			Spec: `{"format": "csv", "path": "abc", "storage_account": null, "container": "c"}`,
			Err:  true,
		},
		{
			Name: "integer storage_account",
			Spec: `{"format": "csv", "path": "abc", "storage_account": 123, "container": "c"}`,
			Err:  true,
		},
		{
			Name: "missing container",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa"}`,
			Err:  true,
		},
		{
			Name: "empty container",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": ""}`,
			Err:  true,
		},
		{
			Name: "null container",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": null}`,
			Err:  true,
		},
		{
			Name: "integer container",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": 123}`,
			Err:  true,
		},

		{
			Name: "null no_rotate",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate": null}`,
			Err:  true,
		},
		{
			Name: "bad no_rotate",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate": 123}`,
			Err:  true,
		},
		{
			Name: "no_rotate:true",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate": true}`,
		},
		{
			Name: "no_rotate:false",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate": false}`,
		},
		{
			Name: "zero batch_size",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "batch_size":0}`,
		},
		{
			Name: "float batch_size",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "batch_size":5.3}`,
		},
		{
			Name: "bad batch_size",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "batch_size":false}`,
		},
		{
			Name: "null batch_size",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "batch_size":null}`,
		},
		{
			Name: "proper batch_size",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "batch_size":123}`,
		},
		{
			Name: "zero batch_size_bytes",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "batch_size_bytes":0}`,
		},
		{
			Name: "float batch_size_bytes",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "batch_size_bytes":5.3}`,
		},
		{
			Name: "bad batch_size_bytes",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "batch_size_bytes":false}`,
		},
		{
			Name: "null batch_size_bytes",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "batch_size_bytes":null}`,
		},
		{
			Name: "proper batch_size_bytes",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "batch_size_bytes":123}`,
		},
		// configtype.Duration is tested in plugin-sdk
		// test only null here
		{
			Name: "null batch_timeout",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "batch_timeout":null}`,
		},

		// no_rotate + batching
		{
			Name: "no_rotate:false & batch_size:100",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate":false, "batch_size":100}`,
		},
		{
			Name: "no_rotate:true & batch_size:100",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate":true, "batch_size":100}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & batch_size:null",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate":false, "batch_size":null}`,
		},
		{
			Name: "no_rotate:true & batch_size:null",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate":true, "batch_size":null}`,
		},
		{
			Name: "no_rotate:false & batch_size_bytes:100",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate":false, "batch_size_bytes":100}`,
		},
		{
			Name: "no_rotate:true & batch_size_bytes:100",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate":true, "batch_size_bytes":100}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & batch_size_bytes:null",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate":false, "batch_size_bytes":null}`,
		},
		{
			Name: "no_rotate:true & batch_size_bytes:null",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate":true, "batch_size_bytes":null}`,
		},
		{
			Name: "no_rotate:false & batch_timeout:100s",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate":false, "batch_timeout":"100s"}`,
		},
		{
			Name: "no_rotate:true & batch_timeout:100s",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate":true, "batch_timeout":"100s"}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & batch_timeout:null",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate":false, "batch_timeout":null}`,
		},
		{
			Name: "no_rotate:true & batch_timeout:null",
			Spec: `{"format": "csv", "path": "abc", "storage_account": "sa", "container": "c", "no_rotate":true, "batch_timeout":null}`,
		},
	})
}
