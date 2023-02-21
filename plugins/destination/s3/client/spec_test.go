package client

import (
	"testing"

	"github.com/cloudquery/filetypes"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
)

func TestSpec_SetDefaults(t *testing.T) {
	cases := []struct {
		Give Spec
		Want Spec
	}{

		{Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}}, Want: Spec{Path: "test/path/{{TABLE}}.json.{{UUID}}", FileSpec: &filetypes.FileSpec{Format: "json"}}},
		{Give: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: &filetypes.FileSpec{Format: "json", FormatSpec: map[string]any{"delimiter": ","}}}, Want: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: &filetypes.FileSpec{Format: "json", FormatSpec: map[string]any{"delimiter": ","}}}},
	}
	for _, tc := range cases {
		got := tc.Give
		got.SetDefaults()
		if diff := cmp.Diff(tc.Want, got, cmpopts.IgnoreUnexported(filetypes.FileSpec{})); diff != "" {
			t.Errorf("SetDefaults() mismatch (-want +got):\n%s", diff)
		}
	}
}

func TestSpec_Validate(t *testing.T) {
	cases := []struct {
		Give    Spec
		WantErr bool
	}{
		{Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}, Bucket: "mybucket"}, WantErr: false},
		{Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}}, WantErr: true}, // no bucket
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: false, Bucket: "mybucket"}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true, Bucket: "mybucket"}, WantErr: true},   // can't have no_rotate and {{UUID}}
		{Give: Spec{Path: "/test/path/{{TABLE}}.{{UUID}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true, Bucket: "mybucket"}, WantErr: true},  // begins with a a slash
		{Give: Spec{Path: "//test/path/{{TABLE}}.{{UUID}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true, Bucket: "mybucket"}, WantErr: true}, // duplicate slashes
		{Give: Spec{Path: "test//path", FileSpec: &filetypes.FileSpec{Format: "json"}, Bucket: "mybucket"}, WantErr: true},                                     // duplicate slashes
	}
	for _, tc := range cases {
		err := tc.Give.Validate()
		if tc.WantErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
