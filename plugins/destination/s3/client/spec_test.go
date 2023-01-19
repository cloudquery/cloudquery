package client

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestSpec_SetDefaults(t *testing.T) {
	headersDefault := true
	cases := []struct {
		Give Spec
		Want Spec
	}{
		{Give: Spec{Path: "test/path", FileSpec: FileSpec{Format: "json"}}, Want: Spec{Path: "test/path/{{TABLE}}.json.{{UUID}}", FileSpec: FileSpec{Format: "json", Delimiter: ',', IncludeHeaders: &headersDefault}}},
		{Give: Spec{Path: "test/path/{{TABLE}}.json"}, Want: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: FileSpec{Delimiter: ',', IncludeHeaders: &headersDefault}}},
	}
	for _, tc := range cases {
		got := tc.Give
		got.SetDefaults()
		if diff := cmp.Diff(tc.Want, got); diff != "" {
			t.Errorf("SetDefaults() mismatch (-want +got):\n%s", diff)
		}
	}
}

func TestSpec_Validate(t *testing.T) {
	cases := []struct {
		Give    Spec
		WantErr bool
	}{
		{Give: Spec{Path: "test/path", FileSpec: FileSpec{Format: "json"}}, WantErr: true},
		{Give: Spec{Path: "test/path", FileSpec: FileSpec{Format: "json"}, Bucket: "mybucket"}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: FileSpec{Format: "json", NoRotate: false}, Bucket: "mybucket"}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: FileSpec{Format: "json", NoRotate: true}, Bucket: "mybucket"}, WantErr: true},
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
