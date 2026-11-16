package spec

import (
	"testing"
	"time"

	"github.com/cloudquery/filetypes/v4"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
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
			tableName:    "test-table",
			expectedPath: "test/test/test-table",
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
			inputPath:    "test/test/{{TABLE}}/{{UUID}}.{{FORMAT}}",
			tableName:    "test-table",
			uuid:         "FAKE-UUID",
			expectedPath: "test/test/test-table/FAKE-UUID.json",
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
		got, err := s.ReplacePathVariables(tc.tableName, tc.uuid, tm, "")
		require.NoError(t, err)
		if diff := cmp.Diff(tc.expectedPath, got); diff != "" {
			t.Errorf("unexpected Path Substitution (-want +got):\n%s", diff)
		}
	}
}

func TestSpec_ReplacePathVariablesKeepsTableInsidePath(t *testing.T) {
	cases := []struct {
		Path   string
		Table  string
		SyncID string
		Want   string
		Err    bool
	}{
		{Path: "cq/data/{{TABLE}}/{{UUID}}.json", Table: "test_table", Want: "cq/data/test_table/FAKE-UUID.json"},
		{Path: "cq/data/{{TABLE}}/{{UUID}}.json", Table: "test.table", Want: "cq/data/test.table/FAKE-UUID.json"},
		{Path: "{{TABLE}}/{{UUID}}.json", Table: "test_table", Want: "test_table/FAKE-UUID.json"},
		{Path: "cq/data/{{TABLE}}/{{UUID}}.json", Table: "../x", Err: true},
		{Path: "cq/data/{{TABLE}}/{{UUID}}.json", Table: "../../other-team/reports", Err: true},
		{Path: "cq/data/{{TABLE}}/{{UUID}}.json", Table: "..", Err: true},
		{Path: "cq/data/{{TABLE}}/{{UUID}}.json", Table: ".", Err: true},
		{Path: "cq/data/{{TABLE}}/{{UUID}}.json", Table: "", Err: true},
		{Path: "cq/data/{{TABLE}}/{{UUID}}.json", Table: "a/b", Err: true},
		{Path: "cq/data/{{TABLE}}/{{UUID}}.json", Table: `a\b`, Err: true},
		{Path: "cq/data/{{TABLE}}/{{UUID}}.json", Table: "a\x00b", Err: true},
		{Path: "cq/data/{{SYNC_ID}}/{{TABLE}}/{{UUID}}.json", Table: "test_table", SyncID: "sync", Want: "cq/data/sync/test_table/FAKE-UUID.json"},
		{Path: "cq/data/{{SYNC_ID}}/{{TABLE}}/{{UUID}}.json", Table: "test_table", SyncID: "../../other-team", Err: true},
		{Path: "{{SYNC_ID}}/{{TABLE}}/{{UUID}}.json", Table: "test_table", SyncID: "..", Err: true},
		{Path: "cq/data", Table: "test_table", Want: "cq/data/test_table.json.FAKE-UUID"},
		{Path: "cq/data", Table: "../x", Err: true},
	}
	tm := time.Date(2021, 3, 5, 4, 1, 2, 3, time.UTC)
	for _, tc := range cases {
		s := &Spec{Path: tc.Path, FileSpec: filetypes.FileSpec{Format: filetypes.FormatTypeJSON}}
		got, err := s.ReplacePathVariables(tc.Table, "FAKE-UUID", tm, tc.SyncID)
		if tc.Err {
			require.Error(t, err, "path %q table %q", tc.Path, tc.Table)
			require.Empty(t, got)
			continue
		}
		require.NoError(t, err, "path %q table %q", tc.Path, tc.Table)
		require.Equal(t, tc.Want, got)
	}
}
