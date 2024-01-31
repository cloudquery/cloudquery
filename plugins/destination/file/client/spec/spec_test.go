package spec

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
)

func TestSpec_SetDefaults(t *testing.T) {
	dur30 := configtype.NewDuration(30 * time.Second)

	cases := []struct {
		Give Spec
		Want Spec
	}{

		{
			Give: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: filetypes.FileSpec{Format: "json", FormatSpec: map[string]any{"delimiter": ", "}}},
			Want: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: filetypes.FileSpec{Format: "json", FormatSpec: map[string]any{"delimiter": ", "}},
				BatchSize: ptr(int64(10000)), BatchSizeBytes: ptr(int64(50 * 1024 * 1024)), BatchTimeout: &dur30},
		},
	}
	for _, tc := range cases {
		got := tc.Give
		got.SetDefaults()
		if diff := cmp.Diff(tc.Want, got, cmpopts.IgnoreUnexported(filetypes.FileSpec{}, configtype.Duration{})); diff != "" {
			t.Errorf("SetDefaults() mismatch (-want +got):\n%s", diff)
		}
		require.Equal(t, tc.Want.BatchTimeout, got.BatchTimeout)
	}
}

func TestSpec_Validate(t *testing.T) {
	zero, one, dur0 := int64(0), int64(1), configtype.NewDuration(0)
	cases := []struct {
		Give    Spec
		WantErr bool
	}{
		{Give: Spec{Path: "test/path", FileSpec: filetypes.FileSpec{Format: "json"}, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: false, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: true, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: true}, // can't have no_rotate and {{UUID}}
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: true, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: false},         // norotate with zero batchsize
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: true}, WantErr: false},                                                                       // norotate with default batchsize
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: true, BatchSize: &one}, WantErr: true},                                                       // norotate with non zero batchsize
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: false, BatchSize: &one, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: true},          // can't have nonzero batch size and no {{UUID}}
	}
	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			err := tc.Give.Validate()
			if tc.WantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSpecUnmarshalJSON(t *testing.T) {
	data := `{
	"format": "csv",
	"format_spec": {
		"skip_header": true,
		"delimiter": "#"
	},
	"path": "abc"
}`
	var s Spec
	require.NoError(t, json.Unmarshal([]byte(data), &s))
	require.Exactly(t, Spec{
		FileSpec: filetypes.FileSpec{
			Format: filetypes.FormatTypeCSV,
			FormatSpec: map[string]any{
				"skip_header": true,
				"delimiter":   "#",
			},
		},
		Path: "abc",
	}, s)
}

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
	"path": "abc"
}`,
		},
		{
			Name: "bad format value", // also a part of embedded FileSpec testing
			Spec: `{"format": "cs22v", "path": "abc"}`,
			Err:  true,
		},
		{
			Name: "missing path",
			Spec: `{"format": "csv"}`,
			Err:  true,
		},
		{
			Name: "empty path",
			Spec: `{"format": "csv", "path": ""}`,
			Err:  true,
		},
		{
			Name: "null path",
			Spec: `{"format": "csv", "path": null}`,
			Err:  true,
		},
		{
			Name: "integer path",
			Spec: `{"format": "csv", "path": 123}`,
			Err:  true,
		},
		{
			Name: "proper path",
			Spec: `{"format": "csv", "path": "abc"}`,
		},
		{
			Name: "null no_rotate",
			Spec: `{"format": "csv", "path": "abc", "no_rotate": null}`,
			Err:  true,
		},
		{
			Name: "bad no_rotate",
			Spec: `{"format": "csv", "path": "abc", "no_rotate": 123}`,
			Err:  true,
		},
		{
			Name: "no_rotate:true",
			Spec: `{"format": "csv", "path": "abc", "no_rotate": true}`,
		},
		{
			Name: "no_rotate:false",
			Spec: `{"format": "csv", "path": "abc", "no_rotate": false}`,
		},
		{
			Name: "zero batch_size",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "batch_size":0}`,
		},
		{
			Name: "float batch_size",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "batch_size":5.3}`,
		},
		{
			Name: "bad batch_size",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "batch_size":false}`,
		},
		{
			Name: "null batch_size",
			Spec: `{"format": "csv", "path": "abc", "batch_size":null}`,
		},
		{
			Name: "proper batch_size",
			Spec: `{"format": "csv", "path": "abc", "batch_size":123}`,
		},
		{
			Name: "zero batch_size_bytes",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "batch_size_bytes":0}`,
		},
		{
			Name: "float batch_size_bytes",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "batch_size_bytes":5.3}`,
		},
		{
			Name: "bad batch_size_bytes",
			Err:  true,
			Spec: `{"format": "csv", "path": "abc", "batch_size_bytes":false}`,
		},
		{
			Name: "null batch_size_bytes",
			Spec: `{"format": "csv", "path": "abc", "batch_size_bytes":null}`,
		},
		{
			Name: "proper batch_size_bytes",
			Spec: `{"format": "csv", "path": "abc", "batch_size_bytes":123}`,
		},
		// configtype.Duration is tested in plugin-sdk
		// test only null here
		{
			Name: "null batch_timeout",
			Spec: `{"format": "csv", "path": "abc", "batch_timeout":null}`,
		},

		// no_rotate + path
		{
			Name: "no_rotate:false & path:{{UUID}}",
			Spec: `{"format": "csv", "path": "{{UUID}}", "no_rotate":false}`,
		},
		{
			Name: "no_rotate:true & path:{{UUID}}",
			Spec: `{"format": "csv", "path": "{{UUID}}", "no_rotate":false}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & path:abc",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":false}`,
		},
		{
			Name: "no_rotate:true & path:{{UUID}}",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":false}`,
		},
		// no_rotate + batching
		{
			Name: "no_rotate:false & batch_size:100",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":false, "batch_size":100}`,
		},
		{
			Name: "no_rotate:true & batch_size:100",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":true, "batch_size":100}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & batch_size:null",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":false, "batch_size":null}`,
		},
		{
			Name: "no_rotate:true & batch_size:null",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":true, "batch_size":null}`,
		},
		{
			Name: "no_rotate:false & batch_size_bytes:100",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":false, "batch_size_bytes":100}`,
		},
		{
			Name: "no_rotate:true & batch_size_bytes:100",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":true, "batch_size_bytes":100}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & batch_size_bytes:null",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":false, "batch_size_bytes":null}`,
		},
		{
			Name: "no_rotate:true & batch_size_bytes:null",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":true, "batch_size_bytes":null}`,
		},
		{
			Name: "no_rotate:false & batch_timeout:100s",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":false, "batch_timeout":"100s"}`,
		},
		{
			Name: "no_rotate:true & batch_timeout:100s",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":true, "batch_timeout":"100s"}`,
			Err:  true,
		},
		{
			Name: "no_rotate:false & batch_timeout:null",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":false, "batch_timeout":null}`,
		},
		{
			Name: "no_rotate:true & batch_timeout:null",
			Spec: `{"format": "csv", "path": "abc", "no_rotate":true, "batch_timeout":null}`,
		},
	})
}
