package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/stretchr/testify/require"
)

func TestRenderUpgradeReportsJSONMatchesText(t *testing.T) {
	setColorOutput(t, false)
	destination := postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe)
	cases := []struct {
		name   string
		report upgradeReport
	}{
		{
			name: "source tables could not be listed",
			report: upgradeReport{
				SourceName: "s3", FromVersion: "v1.0.0", ToVersion: "v2.0.0", SourceUnknown: true,
				SourceGaps: []string{"s3 v1.0.0: tables could not be listed: source returned no tables with or without a connection"},
			},
		},
		{
			name: "destination could not assess tables",
			report: upgradeReport{
				SourceName: "datadog", FromVersion: "v5.19.10", ToVersion: "v6.0.0", Destination: destination,
				Findings: unknownTableFindings([]upgradeTableSchemas{{Name: "datadog_slos"}, {Name: "datadog_monitors"}}, destinationNoAssessmentReason),
			},
		},
		{
			name: "no changes with tables listed using a connection",
			report: upgradeReport{
				SourceName: "bigquery", FromVersion: "v1.0.0", ToVersion: "v2.0.0", Destination: destination,
				SourceGaps: []string{"bigquery v2.0.0: tables were listed with a connection (metadata only, no rows read)"},
				Findings:   []*pluginPb.AssessTables_TableFinding{{TableName: "bigquery_table", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE}},
			},
		},
		{
			name: "file output differs for equivalent values",
			report: upgradeReport{
				SourceName: "datadog", FromVersion: "v5.19.10", ToVersion: "v6.0.0", Destination: destination,
				Findings: []*pluginPb.AssessTables_TableFinding{{
					TableName: "datadog_monitors",
					Category:  pluginPb.AssessTables_CATEGORY_NO_CHANGE,
					Columns: []*pluginPb.AssessTables_ColumnFinding{{
						ColumnName: "tags", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE, OldType: "list<string>", NewType: "json",
						Evidence: []*pluginPb.AssessTables_Evidence{{SyntheticValue: `["env:prod"]`, Before: `{"tags":["env:prod"]}`, After: `{"tags":"[\"env:prod\"]"}`}},
					}},
				}},
			},
		},
		{
			name: "manual migration and removed tables",
			report: upgradeReport{
				SourceName: "datadog", FromVersion: "v5.19.10", ToVersion: "v6.0.0", Destination: destination,
				RemovedTables: []string{"datadog_removed"},
				Findings: []*pluginPb.AssessTables_TableFinding{{
					TableName: "datadog_monitors",
					Category:  pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED,
					Columns: []*pluginPb.AssessTables_ColumnFinding{
						{ColumnName: "tags", Category: pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED, OldType: "text[]", NewType: "jsonb"},
					},
				}},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var text bytes.Buffer
			require.NoError(t, renderUpgradeReport(&text, tc.report))
			got := renderUpgradeReportJSON(t, tc.report)

			lines := strings.Split(text.String(), "\n")
			verdictLine := lines[1]
			if tc.report.Destination != nil {
				verdictLine = lines[2]
			}
			wantVerdict, wantSummary, _ := strings.Cut(verdictLine, " — ")
			require.Equal(t, wantVerdict, got.Verdict)
			require.Equal(t, wantSummary, got.Summary)
			require.Equal(t, textCoverageGaps(text.String()), got.CoverageGaps)
			if wantAction, found := textAction(text.String()); found {
				require.Equal(t, wantAction, got.Action)
			}
		})
	}
}

func renderUpgradeReportJSON(t *testing.T, report upgradeReport) upgradeReportJSON {
	t.Helper()
	var out bytes.Buffer
	require.NoError(t, renderUpgradeReportsJSON(&out, []upgradeReport{report}))
	var raw struct {
		Reports []map[string]json.RawMessage `json:"reports"`
	}
	require.NoError(t, json.Unmarshal(out.Bytes(), &raw))
	require.Len(t, raw.Reports, 1)
	for _, key := range []string{"tables", "removed_tables", "output_comparisons", "coverage_gaps"} {
		require.NotEqual(t, "null", string(raw.Reports[0][key]), key)
	}
	var got upgradeReportsJSON
	require.NoError(t, json.Unmarshal(out.Bytes(), &got))
	return got.Reports[0]
}

func textCoverageGaps(report string) []string {
	gaps := []string{}
	_, section, found := strings.Cut(report, "\nCoverage gaps\n")
	if !found {
		return gaps
	}
	for line := range strings.Lines(section) {
		gap, isGap := strings.CutPrefix(strings.TrimSuffix(line, "\n"), "  ")
		if !isGap {
			break
		}
		gaps = append(gaps, gap)
	}
	return gaps
}

func textAction(report string) (string, bool) {
	_, rest, found := strings.Cut(report, "\nAction: ")
	action, _, _ := strings.Cut(rest, "\n")
	return action, found
}

func TestRenderUpgradeReportsJSON(t *testing.T) {
	report := upgradeReport{
		SourceName: "okta", FromVersion: "v6.8.2", ToVersion: "v7.0.0",
		Destination:   postgresqlDestinationSpec(specs.WriteModeAppend, specs.MigrateModeSafe),
		RemovedTables: []string{"okta_removed"},
		Findings: []*pluginPb.AssessTables_TableFinding{
			{
				TableName: "okta_policy_rules",
				Category:  pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE,
				Columns: []*pluginPb.AssessTables_ColumnFinding{
					{ColumnName: "id", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE},
					{ColumnName: "policy_id", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, NewType: "text"},
				},
			},
			{
				TableName: "okta_users",
				Category:  pluginPb.AssessTables_CATEGORY_NO_CHANGE,
				Columns: []*pluginPb.AssessTables_ColumnFinding{{
					ColumnName: "tags", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE, OldType: "list<string>", NewType: "json",
					Evidence: []*pluginPb.AssessTables_Evidence{{SyntheticValue: `["env:prod"]`, Before: `{"tags":["env:prod"]}`, After: `{"tags":["env:prod"]}`}},
				}},
			},
		},
	}

	var out bytes.Buffer
	require.NoError(t, renderUpgradeReportsJSON(&out, []upgradeReport{report}))

	require.JSONEq(t, `{"reports": [{
		"source": {"name": "okta", "from_version": "v6.8.2", "to_version": "v7.0.0"},
		"destination": {"name": "postgresql", "registry": "cloudquery", "path": "cloudquery/postgresql", "version": "v8.14.0", "write_mode": "append", "pk_mode": "default", "migrate_mode": "safe"},
		"verdict": "SELECTED TABLES REMOVED",
		"summary": "1 changed table, 1 removed table",
		"tables": [{
			"name": "okta_policy_rules",
			"changes": [{"kind": "added", "column": "policy_id", "new_type": "text"}],
			"outcomes": {
				"safe": {"result": "applied", "text": "migrated in place"},
				"forced": {"result": "applied", "text": "migrated in place"}
			}
		}],
		"removed_tables": ["okta_removed"],
		"output_comparisons": [{
			"table": "okta_users",
			"column": "tags",
			"old_type": "list<string>",
			"new_type": "json",
			"evidence": [{"synthetic_value": "[\"env:prod\"]", "before": "{\"tags\":[\"env:prod\"]}", "after": "{\"tags\":[\"env:prod\"]}"}]
		}],
		"coverage_gaps": [],
		"action": "remove explicit selections and update dependent consumers."
	}]}`, out.String())
	require.NotContains(t, out.String(), `\u003c`)
}

func TestRenderUpgradeReportsJSONWithoutDestination(t *testing.T) {
	var out bytes.Buffer
	require.NoError(t, renderUpgradeReportsJSON(&out, []upgradeReport{{SourceName: "s3", FromVersion: "v1.0.0", ToVersion: "v2.0.0", SourceUnknown: true}}))

	require.JSONEq(t, `{"reports": [{
		"source": {"name": "s3", "from_version": "v1.0.0", "to_version": "v2.0.0"},
		"destination": null,
		"verdict": "UNKNOWN",
		"tables": [],
		"removed_tables": [],
		"output_comparisons": [],
		"coverage_gaps": [],
		"action": "review the source changelog for what this check could not assess."
	}]}`, out.String())
}

func TestUpgradeCheckInvalidOutput(t *testing.T) {
	cmd := NewCmdRoot()
	cmd.SetArgs(append([]string{"upgrade", "check", "testdata/transformation.yml", "--source", "test", "--to", "v4.7.0", "--output", "yaml"}, testCommandArgs(t)...))

	require.ErrorContains(t, cmd.Execute(), `invalid output format "yaml". One of: text, json`)
}

func TestUpgradeCheckPrintsCompletedReportsBeforeADestinationFails(t *testing.T) {
	cmd := NewCmdRoot()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetArgs(append([]string{"upgrade", "check", "testdata/upgrade-check-second-destination-fails.yml", "--source", "test", "--to", "v4.7.0"}, testCommandArgs(t)...))

	require.ErrorContains(t, cmd.Execute(), "failed to start destination broken")
	require.Contains(t, out.String(), "test v4.5.1 → v4.7.0 | test (cloudquery/test@v2.5.1)\n")
}

func TestUpgradeCheckJSONKeepsDownloadMessagesOffStdout(t *testing.T) {
	stdout := redirectStdFile(t, &os.Stdout)
	stderr := redirectStdFile(t, &os.Stderr)
	cmd := NewCmdRoot()
	cmd.SetArgs(append([]string{"upgrade", "check", "testdata/transformation.yml", "--source", "test", "--to", "v4.7.0", "--output", "json"}, testCommandArgs(t)...))

	requireUpgradeCheckExitCode(t, cmd.Execute(), 4)
	var report upgradeReportsJSON
	require.NoError(t, json.Unmarshal([]byte(readStdFile(t, stdout)), &report))
	require.Len(t, report.Reports, 1)
	require.Contains(t, readStdFile(t, stderr), "Downloading ")
}

func redirectStdFile(t *testing.T, target **os.File) *os.File {
	t.Helper()
	file, err := os.CreateTemp(t.TempDir(), "std")
	require.NoError(t, err)
	original := *target
	*target = file
	t.Cleanup(func() {
		*target = original
		_ = file.Close()
	})
	return file
}

func readStdFile(t *testing.T, file *os.File) string {
	t.Helper()
	b, err := os.ReadFile(file.Name())
	require.NoError(t, err)
	return string(b)
}
