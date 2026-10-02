package cmd

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/stretchr/testify/require"
)

func TestRenderUpgradeReportsJSONCoverageGaps(t *testing.T) {
	setColorOutput(t, false)
	destination := postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale)
	cases := []struct {
		name     string
		report   upgradeReport
		wantGaps []string
	}{
		{
			name: "source tables could not be listed",
			report: upgradeReport{
				SourceName: "s3", FromVersion: "v1.0.0", ToVersion: "v2.0.0", SourceUnknown: true,
				SourceGaps: []string{"s3 v1.0.0: tables could not be listed: source returned no tables with or without a connection"},
			},
			wantGaps: []string{"s3 v1.0.0: tables could not be listed: source returned no tables with or without a connection"},
		},
		{
			name: "destination could not assess tables",
			report: upgradeReport{
				SourceName: "datadog", FromVersion: "v5.19.10", ToVersion: "v6.0.0", Destination: destination,
				Findings: unknownTableFindings([]upgradeTableSchemas{{Name: "datadog_monitors"}, {Name: "datadog_slos"}}, destinationNoAssessmentReason),
			},
			wantGaps: []string{
				"datadog_monitors: " + destinationNoAssessmentReason,
				"datadog_slos: " + destinationNoAssessmentReason,
			},
		},
		{
			name: "no changes with tables listed using a connection",
			report: upgradeReport{
				SourceName: "bigquery", FromVersion: "v1.0.0", ToVersion: "v2.0.0", Destination: destination,
				SourceGaps: []string{"bigquery v2.0.0: tables were listed with a connection (metadata only, no rows read)"},
				Findings:   []*pluginPb.AssessTables_TableFinding{{TableName: "bigquery_table", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE}},
			},
			wantGaps: []string{"bigquery v2.0.0: tables were listed with a connection (metadata only, no rows read)"},
		},
		{
			name: "no coverage gaps",
			report: upgradeReport{
				SourceName: "datadog", FromVersion: "v5.19.10", ToVersion: "v6.0.0", Destination: destination,
				Findings: []*pluginPb.AssessTables_TableFinding{{TableName: "datadog_users", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE}},
			},
			wantGaps: []string{},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var text bytes.Buffer
			require.NoError(t, renderUpgradeReport(&text, tc.report))
			require.Equal(t, tc.wantGaps, textCoverageGaps(text.String()))

			var out bytes.Buffer
			require.NoError(t, renderUpgradeReportsJSON(&out, []upgradeReport{tc.report}))
			var got struct {
				Reports []map[string]json.RawMessage `json:"reports"`
			}
			require.NoError(t, json.Unmarshal(out.Bytes(), &got))
			require.Len(t, got.Reports, 1)
			require.Contains(t, got.Reports[0], "coverage_gaps")
			var gaps []string
			require.NoError(t, json.Unmarshal(got.Reports[0]["coverage_gaps"], &gaps))
			require.Equal(t, tc.wantGaps, gaps)
		})
	}
}

func textCoverageGaps(report string) []string {
	gaps := []string{}
	_, section, found := strings.Cut(report, "\nCoverage gaps:\n")
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

func TestRenderUpgradeReportsJSON(t *testing.T) {
	report := upgradeReport{
		SourceName: "datadog", FromVersion: "v5.19.10", ToVersion: "v6.0.0",
		Destination:   postgresqlDestinationSpec(specs.WriteModeAppend),
		RemovedTables: []string{"datadog_removed"},
		Findings: []*pluginPb.AssessTables_TableFinding{{
			TableName:          "datadog_monitors",
			Category:           pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED,
			SafeModeBehavior:   "writes files with the new schema",
			ForcedModeBehavior: "writes files with the new schema",
			Columns: []*pluginPb.AssessTables_ColumnFinding{
				{ColumnName: "id", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE},
				{
					ColumnName: "tags", Category: pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED, OldType: "list<string>", NewType: "json",
					Evidence: []*pluginPb.AssessTables_Evidence{{SyntheticValue: `["env:prod"]`, Before: `{"tags":["env:prod"]}`, After: `{"tags":["env:prod"]}`}},
				},
			},
		}},
	}

	var out bytes.Buffer
	require.NoError(t, renderUpgradeReportsJSON(&out, []upgradeReport{report}))

	require.JSONEq(t, `{"reports": [{
		"source": {"name": "datadog", "from_version": "v5.19.10", "to_version": "v6.0.0"},
		"destination": {"name": "postgresql", "registry": "cloudquery", "path": "cloudquery/postgresql", "version": "v8.14.0", "write_mode": "append", "pk_mode": "default", "migrate_mode": "safe"},
		"verdicts": [
			{"category": "table_removed", "tables": ["datadog_removed"], "changed_columns": 0, "action": "remove explicit selections and update dependent consumers."},
			{"category": "file_schema_changed", "tables": ["datadog_monitors"], "changed_columns": 1, "action": "review readers that combine old and new files."}
		],
		"tables": [{
			"name": "datadog_monitors",
			"category": "file_schema_changed",
			"safe_mode_behavior": "writes files with the new schema",
			"forced_mode_behavior": "writes files with the new schema",
			"columns": [{
				"name": "tags",
				"category": "file_schema_changed",
				"source_type": {},
				"destination_type": {"old": "list<string>", "new": "json"},
				"evidence": [{"synthetic_value": "[\"env:prod\"]", "before": "{\"tags\":[\"env:prod\"]}", "after": "{\"tags\":[\"env:prod\"]}"}]
			}],
			"coverage_incomplete": false
		}],
		"coverage_gaps": []
	}]}`, out.String())
	require.NotContains(t, out.String(), `\u003c`)
}

func TestRenderUpgradeReportsJSONWithoutDestination(t *testing.T) {
	var out bytes.Buffer
	require.NoError(t, renderUpgradeReportsJSON(&out, []upgradeReport{{SourceName: "s3", FromVersion: "v1.0.0", ToVersion: "v2.0.0", SourceUnknown: true}}))

	require.JSONEq(t, `{"reports": [{
		"source": {"name": "s3", "from_version": "v1.0.0", "to_version": "v2.0.0"},
		"destination": null,
		"verdicts": [{"category": "unknown", "tables": [], "changed_columns": 0, "action": "review the source changelog for what this check could not assess."}],
		"tables": [],
		"coverage_gaps": []
	}]}`, out.String())
}

func TestUpgradeCheckInvalidOutput(t *testing.T) {
	cmd := NewCmdRoot()
	cmd.SetArgs(append([]string{"upgrade", "check", "testdata/transformation.yml", "--source", "test", "--to", "v4.7.0", "--output", "yaml"}, testCommandArgs(t)...))

	require.ErrorContains(t, cmd.Execute(), `invalid output format "yaml". One of: text, json`)
}
