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
    "bucket": "abc"
}`,
		},
		{
			Name: "bad format value", // also a part of embedded FileSpec testing
			Spec: `{"format": "cs22v", "path": "abc"}`,
			Err:  true,
		},
		{
			Name: "minimal",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc"}`,
		},
		{
			Name: "missing path",
			Spec: `{"format": "csv", "bucket": "abc"}`,
			Err:  true,
		},
		{
			Name: "empty path",
			Spec: `{"format": "csv", "path": "", "bucket": "abc"}`,
			Err:  true,
		},
		{
			Name: "null path",
			Spec: `{"format": "csv", "path": null, "bucket": "abc"}`,
			Err:  true,
		},
		{
			Name: "integer path",
			Spec: `{"format": "csv", "path": 123, "bucket": "abc"}`,
			Err:  true,
		},
		{
			Name: "missing bucket",
			Spec: `{"format": "csv", "path": "abc"}`,
			Err:  true,
		},
		{
			Name: "empty bucket",
			Spec: `{"format": "csv", "path": "abc", "bucket": ""}`,
			Err:  true,
		},
		{
			Name: "null bucket",
			Spec: `{"format": "csv", "path": "abc", "bucket": null}`,
			Err:  true,
		},
		{
			Name: "integer bucket",
			Spec: `{"format": "csv", "path": "abc", "bucket": 123}`,
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
			Name: "null no_rotate",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate": null}`,
			Err:  true,
		},
		{
			Name: "bad no_rotate",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate": 123}`,
			Err:  true,
		},
		{
			Name: "no_rotate:true",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate": true}`,
		},
		{
			Name: "no_rotate:false",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate": false}`,
		},
		{
			Name: "zero batch_size",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "batch_size":0}`,
		},
		{
			Name: "float batch_size",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "batch_size":5.3}`,
		},
		{
			Name: "bad batch_size",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "batch_size":false}`,
		},
		{
			Name: "null batch_size",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "batch_size":null}`,
		},
		{
			Name: "proper batch_size",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "batch_size":123}`,
		},
		{
			Name: "zero batch_size_bytes",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "batch_size_bytes":0}`,
		},
		{
			Name: "float batch_size_bytes",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "batch_size_bytes":5.3}`,
		},
		{
			Name: "bad batch_size_bytes",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "batch_size_bytes":false}`,
		},
		{
			Name: "null batch_size_bytes",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "batch_size_bytes":null}`,
		},
		{
			Name: "proper batch_size_bytes",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "batch_size_bytes":123}`,
		},
		// configtype.Duration is tested in plugin-sdk
		// test only null here
		{
			Name: "null batch_timeout",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "batch_timeout":null}`,
		},

		// no_rotate + path({{UUID}})
		{
			Name: "no_rotate:false & path:{{UUID}}",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "no_rotate":false}`,
		},
		{
			Name: "no_rotate:true & path:{{UUID}}",
			Spec: `{"format": "csv", "path": "{{UUID}}", "bucket": "b", "no_rotate":true}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & path:abc",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "no_rotate":false}`,
		},
		{
			Name: "no_rotate:true & path:abc",
			Spec: `{"format": "csv", "path": "abc", "bucket": "b", "no_rotate":true}`,
		},
		{
			Name: "no_rotate:false & path:{{TABLE}}",
			Spec: `{"format": "csv", "path": "{{TABLE}}", "bucket": "b", "no_rotate":false}`,
			Err:  true,
		},

		// no_rotate + batching
		{
			Name: "no_rotate:false & batch_size:100",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate":false, "batch_size":100}`,
		},
		{
			Name: "no_rotate:true & batch_size:100",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate":true, "batch_size":100}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & batch_size:null",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate":false, "batch_size":null}`,
		},
		{
			Name: "no_rotate:true & batch_size:null",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate":true, "batch_size":null}`,
		},
		{
			Name: "no_rotate:false & batch_size_bytes:100",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate":false, "batch_size_bytes":100}`,
		},
		{
			Name: "no_rotate:true & batch_size_bytes:100",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate":true, "batch_size_bytes":100}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & batch_size_bytes:null",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate":false, "batch_size_bytes":null}`,
		},
		{
			Name: "no_rotate:true & batch_size_bytes:null",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate":true, "batch_size_bytes":null}`,
		},
		{
			Name: "no_rotate:false & batch_timeout:100s",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate":false, "batch_timeout":"100s"}`,
		},
		{
			Name: "no_rotate:true & batch_timeout:100s",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate":true, "batch_timeout":"100s"}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & batch_timeout:null",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate":false, "batch_timeout":null}`,
		},
		{
			Name: "no_rotate:true & batch_timeout:null",
			Spec: `{"format": "csv", "path": "abc", "bucket": "abc", "no_rotate":true, "batch_timeout":null}`,
		},
	})
}
