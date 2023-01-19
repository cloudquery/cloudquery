package client

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSpec_SetDefaults(t *testing.T) {
	cases := []struct {
		Give Spec
		Want Spec
	}{
		{Give: Spec{Path: "test/path", Format: "json"}, Want: Spec{Path: "test/path/{{TABLE}}.json.{{UUID}}", Format: "json", Delimiter: ','}},
		{Give: Spec{Path: "test/path/{{TABLE}}.json"}, Want: Spec{Path: "test/path/{{TABLE}}.json", Delimiter: ','}},
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
		{Give: Spec{Path: "test/path", Format: "json"}, WantErr: true},
		{Give: Spec{Path: "test/path", Format: "json", Bucket: "mybucket"}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", Format: "json", Bucket: "mybucket", NoRotate: false}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", Format: "json", Bucket: "mybucket", NoRotate: true}, WantErr: true},
	}
	for _, tc := range cases {
		err := tc.Give.Validate()
		gotErr := err != nil
		if diff := cmp.Diff(tc.WantErr, gotErr); diff != "" {
			s := "nil"
			if tc.WantErr {
				s = "error"
			}
			t.Errorf("Validate(%v) got err = %v, want %v", tc.Give, err, s)
		}
	}
}
