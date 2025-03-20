package spec

import (
	"fmt"
	"testing"
	"time"

	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
)

const (
	region = "us-east-1"
)

func TestSpec_SetDefaults(t *testing.T) {
	dur0, dur30 := configtype.NewDuration(0), configtype.NewDuration(30*time.Second)

	cases := []struct {
		Give Spec
		Want Spec
	}{

		{
			Give: Spec{Path: "test/path", FileSpec: filetypes.FileSpec{Format: "json"}},
			Want: Spec{Path: "test/path/{{TABLE}}.json.{{UUID}}", FileSpec: filetypes.FileSpec{Format: "json"}, TestWrite: boolPtr(true), BatchSize: int64Ptr(10000), BatchSizeBytes: int64Ptr(50 * 1024 * 1024), BatchTimeout: &dur30, MaxRetries: intPtr(3), MaxBackoff: intPtr(30), PartSize: int64Ptr(5242880)},
		},
		{
			Give: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: filetypes.FileSpec{Format: "json", FormatSpec: map[string]any{"delimiter": ","}}, TestWrite: boolPtr(false)},
			Want: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: filetypes.FileSpec{Format: "json", FormatSpec: map[string]any{"delimiter": ","}}, TestWrite: boolPtr(false), BatchSize: int64Ptr(10000), BatchSizeBytes: int64Ptr(50 * 1024 * 1024), BatchTimeout: &dur30, MaxRetries: intPtr(3), MaxBackoff: intPtr(30), PartSize: int64Ptr(5242880)},
		},
		{
			Give: Spec{Path: "test/path", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: true},
			Want: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: filetypes.FileSpec{Format: "json"}, TestWrite: boolPtr(true), NoRotate: true, BatchSize: int64Ptr(0), BatchSizeBytes: int64Ptr(0), BatchTimeout: &dur0, MaxRetries: intPtr(3), MaxBackoff: intPtr(30), PartSize: int64Ptr(5242880)},
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
		{Give: Spec{Path: "test/path", FileSpec: filetypes.FileSpec{Format: "json"}, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0, ACL: "bucket-owner-full-control"}, WantErr: false},
		{Give: Spec{Path: "test/path", FileSpec: filetypes.FileSpec{Format: "json"}, Region: "region", BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: true}, // no bucket
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: false, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: filetypes.FileSpec{Format: "parquet"}, NoRotate: false, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0, GenerateEmptyObjects: true}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: filetypes.FileSpec{Format: "parquet"}, NoRotate: false, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0, GenerateEmptyObjects: false}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: false, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0, GenerateEmptyObjects: true}, WantErr: true}, // when empty_objects is enabled, format must be parquet
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: true, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: true},                              // can't have no_rotate and {{UUID}}
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: false, Bucket: "mybucket", Region: region, BatchSize: &one, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: true},                                       // can't have nonzero batch size and no {{UUID}}
		{Give: Spec{Path: "/test/path/{{TABLE}}.{{UUID}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: true, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: true},                             // begins with a slash
		{Give: Spec{Path: "//test/path/{{TABLE}}.{{UUID}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: true, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: true},                            // duplicate slashes
		{Give: Spec{Path: "test//path", FileSpec: filetypes.FileSpec{Format: "json"}, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: true},                                                                // duplicate slashes
		{Give: Spec{Path: "test/path", FileSpec: filetypes.FileSpec{Format: "json"}, Bucket: "mybucket", Region: region, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0, ACL: "invalid"}, WantErr: true},
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

func intPtr(i int) *int {
	return &i
}

func TestGetContentType(t *testing.T) {
	cases := []struct {
		Give        Spec
		ContentType string
	}{
		{Give: Spec{Path: "test/path", FileSpec: filetypes.FileSpec{Format: "json"}}, ContentType: "application/json"},
		{Give: Spec{Path: "test/path", FileSpec: filetypes.FileSpec{Format: "csv"}}, ContentType: "text/csv"},
		{Give: Spec{Path: "test/path", FileSpec: filetypes.FileSpec{Format: "parquet"}}, ContentType: "application/vnd.apache.parquet"},
		{Give: Spec{Path: "test/path", FileSpec: filetypes.FileSpec{Format: "parquet", Compression: "gzip"}}, ContentType: "application/gzip"},
		{Give: Spec{Path: "test/path", ContentType: "application/custom", FileSpec: filetypes.FileSpec{Format: "parquet", Compression: "gzip"}}, ContentType: "application/custom"},
	}
	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			require.Equal(t, tc.ContentType, tc.Give.GetContentType())
		})
	}
}
