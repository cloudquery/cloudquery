package spec

import (
	"testing"
	"time"

	"github.com/cloudquery/filetypes/v4"
	"github.com/google/go-cmp/cmp"
)

func TestReplacePathVariables(t *testing.T) {
	cases := []struct {
		inputPath    string
		uuid         string
		tableName    string
		expectedPath string
	}{
		{
			inputPath:    "test/test/{{TABLE}}/{{UUID}}",
			uuid:         "",
			tableName:    "",
			expectedPath: "test/test",
		},
		{
			inputPath:    "test/test/{{TABLE}}/{{UUID}}.json",
			tableName:    "test-table",
			uuid:         "",
			expectedPath: "test/test/test-table/.json",
		},
		{
			inputPath:    "test/test/{{TABLE}}/{{UUID}}.json",
			tableName:    "test-table",
			uuid:         "FAKE-UUID",
			expectedPath: "test/test/test-table/FAKE-UUID.json",
		},
		{
			inputPath:    "test/test/{{TABLE}}/{{UUID}}.json",
			tableName:    "",
			uuid:         "FAKE-UUID",
			expectedPath: "test/test/FAKE-UUID.json",
		},
		{
			inputPath:    "test/test/{{TABLE}}/{{UUID}}.{{FORMAT}}",
			tableName:    "",
			uuid:         "FAKE-UUID",
			expectedPath: "test/test/FAKE-UUID.json",
		},
		{
			inputPath:    "test/test/{{TABLE}}/year={{YEAR}}/month={{MONTH}}/day={{DAY}}/hour={{HOUR}}/minute={{MINUTE}}/{{UUID}}.json",
			tableName:    "test-table",
			uuid:         "FAKE-UUID",
			expectedPath: "test/test/test-table/year=2021/month=03/day=05/hour=04/minute=01/FAKE-UUID.json",
		},
	}
	tm := time.Date(2021, 3, 5, 4, 1, 2, 3, time.UTC)
	for _, tc := range cases {
		s := &Spec{
			Path:     tc.inputPath,
			FileSpec: filetypes.FileSpec{Format: filetypes.FormatTypeJSON},
		}
		if diff := cmp.Diff(tc.expectedPath, s.ReplacePathVariables(tc.tableName, tc.uuid, tm, "")); diff != "" {
			t.Errorf("unexpected Path Substitution (-want +got):\n%s", diff)
		}
	}
}
