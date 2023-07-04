package client

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

func TestSpec_SetDefaults(t *testing.T) {
	dur0, dur30 := configtype.NewDuration(0), configtype.NewDuration(30*time.Second)

	cases := []struct {
		Give Spec
		Want Spec
	}{

		{
			Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}},
			Want: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}, BatchSize: int64Ptr(10000), BatchSizeBytes: int64Ptr(50 * 1024 * 1024), BatchTimeout: &dur30},
		},
		{
			Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true},
			Want: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true, BatchSize: int64Ptr(0), BatchSizeBytes: int64Ptr(0), BatchTimeout: &dur0},
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
		{Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}, StorageAccount: storage_account, Container: container, BatchSize: &zero, BatchSizeBytes: &zero}, WantErr: false},
		{Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}, Container: container, BatchSize: &zero, BatchSizeBytes: &zero}, WantErr: true}, // no StorageAccount
		{Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: false, StorageAccount: storage_account, Container: container, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: false},
		{Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true, StorageAccount: storage_account, Container: container, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: false},
		{Give: Spec{Path: "test/path", FileSpec: &filetypes.FileSpec{Format: "json"}, NoRotate: true, StorageAccount: storage_account, Container: container, BatchSize: &one, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: true},
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
