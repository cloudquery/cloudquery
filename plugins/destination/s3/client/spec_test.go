package client

import (
	"fmt"
	"testing"

	"github.com/cloudquery/filetypes/v4"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
)

func TestSpec_SetDefaults(t *testing.T) {
	cases := []struct {
		Give Spec
		Want Spec
	}{

		{
			Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}},
			Want: Spec{Path: "test/path/{{TABLE}}.json.{{UUID}}", FileSpec: &filetypes.FileSpec{Format: "json"}, TestWrite: boolPtr(true), BatchSize: int64Ptr(10000), BatchSizeBytes: int64Ptr(50 * 1024 * 1024), BatchTimeoutMs: int64Ptr(30000)},
		},
		{
			Give: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: &filetypes.FileSpec{Format: "json", FormatSpec: map[string]any{"delimiter": ","}}, TestWrite: boolPtr(false)},
			Want: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: &filetypes.FileSpec{Format: "json", FormatSpec: map[string]any{"delimiter": ","}}, TestWrite: boolPtr(false), BatchSize: int64Ptr(10000), BatchSizeBytes: int64Ptr(50 * 1024 * 1024), BatchTimeoutMs: int64Ptr(30000)},
		},
		{
			Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true},
			Want: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: &filetypes.FileSpec{Format: "json"}, TestWrite: boolPtr(true), NoRotate: true, BatchSize: int64Ptr(0), BatchSizeBytes: int64Ptr(0), BatchTimeoutMs: int64Ptr(0)},
		},
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
	zero, one := int64(0), int64(1)
	cases := []struct {
		Give    Spec
		WantErr bool
	}{
		{Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeoutMs: &zero}, WantErr: false},
		{Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}, Region: "region", BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeoutMs: &zero}, WantErr: true}, // no bucket
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: false, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeoutMs: &zero}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeoutMs: &zero}, WantErr: true},   // can't have no_rotate and {{UUID}}
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: false, Bucket: "mybucket", Region: region, BatchSize: &one, BatchSizeBytes: &zero, BatchTimeoutMs: &zero}, WantErr: true},            // can't have nonzero batch size and no {{UUID}}
		{Give: Spec{Path: "/test/path/{{TABLE}}.{{UUID}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeoutMs: &zero}, WantErr: true},  // begins with a slash
		{Give: Spec{Path: "//test/path/{{TABLE}}.{{UUID}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeoutMs: &zero}, WantErr: true}, // duplicate slashes
		{Give: Spec{Path: "test//path", FileSpec: &filetypes.FileSpec{Format: "json"}, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeoutMs: &zero}, WantErr: true},                                     // duplicate slashes
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

func boolPtr(b bool) *bool {
	return &b
}

func int64Ptr(i int64) *int64 {
	return &i
}
