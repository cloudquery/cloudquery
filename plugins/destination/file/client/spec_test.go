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
			Give: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: &filetypes.FileSpec{Format: "json", FormatSpec: map[string]any{"delimiter": ","}}},
			Want: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: &filetypes.FileSpec{Format: "json", FormatSpec: map[string]any{"delimiter": ","}}, BatchSize: int64Ptr(10000), BatchSizeBytes: int64Ptr(50 * 1024 * 1024)},
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
		{Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}, BatchSize: &zero, BatchSizeBytes: &zero}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: false, BatchSize: &zero, BatchSizeBytes: &zero}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true, BatchSize: &zero, BatchSizeBytes: &zero}, WantErr: true}, // can't have no_rotate and {{UUID}}
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true, BatchSize: &zero, BatchSizeBytes: &zero}, WantErr: false},         // norotate with zero batchsize
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true}, WantErr: false},                                                  // norotate with default batchsize
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true, BatchSize: &one}, WantErr: true},                                  // norotate with non zero batchsize
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: false, BatchSize: &one, BatchSizeBytes: &zero}, WantErr: true},          // can't have nonzero batch size and no {{UUID}}
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

func int64Ptr(i int64) *int64 {
	return &i
}
