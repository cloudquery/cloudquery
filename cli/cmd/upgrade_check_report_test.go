package cmd

import (
	"bytes"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/fatih/color"
	"github.com/stretchr/testify/require"
)

func setColorOutput(t *testing.T, enabled bool) {
	noColor := color.NoColor
	color.NoColor = !enabled
	t.Cleanup(func() { color.NoColor = noColor })
}

func postgresqlDestinationSpec(writeMode specs.WriteMode) *specs.Destination {
	return &specs.Destination{
		Metadata:  specs.Metadata{Name: "postgresql", Path: "cloudquery/postgresql", Version: "v8.14.0", Registry: specs.RegistryCloudQuery},
		WriteMode: writeMode,
	}
}

func TestRenderUpgradeReport(t *testing.T) {
	tagsTable := func(name string, tagsType arrow.DataType) *schema.Table {
		return &schema.Table{Name: name, Columns: schema.ColumnList{{Name: "tags", Type: tagsType}}}
	}
	tagsFinding := func(table string) *pluginPb.AssessTables_TableFinding {
		return &pluginPb.AssessTables_TableFinding{
			TableName:          table,
			Category:           pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED,
			SafeModeBehavior:   "rejects this change",
			ForcedModeBehavior: "drops and recreates the affected table",
			Columns: []*pluginPb.AssessTables_ColumnFinding{
				{ColumnName: "id", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE, OldType: "text", NewType: "text"},
				{
					ColumnName:         "tags",
					Category:           pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED,
					OldType:            "text[]",
					NewType:            "jsonb",
					SafeModeBehavior:   "rejects this change",
					ForcedModeBehavior: "drops and recreates the affected table",
				},
			},
		}
	}

	cases := []struct {
		name   string
		report upgradeReport
		want   string
	}{
		{
			name: "type change requiring migration",
			report: upgradeReport{
				SourceName:  "datadog",
				FromVersion: "v5.19.10",
				ToVersion:   "v6.0.0",
				Destination: postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale),
				Tables: map[string]upgradeTablePair{
					"datadog_monitors": {
						From: tagsTable("datadog_monitors", arrow.ListOf(arrow.BinaryTypes.String)),
						To:   tagsTable("datadog_monitors", types.ExtensionTypes.JSON),
					},
				},
				Findings: []*pluginPb.AssessTables_TableFinding{
					tagsFinding("datadog_monitors"),
					tagsFinding("datadog_dashboards"),
					{TableName: "datadog_users", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE},
				},
			},
			want: `datadog v5.19.10 → v6.0.0 | postgresql (cloudquery/postgresql@v8.14.0)
write_mode: overwrite-delete-stale | pk_mode: default
REVIEW REQUIRED — 2 tables / 2 columns

datadog_dashboards.tags
  postgresql:  text[] → jsonb
  Safe mode:   rejects this change
  Forced mode: drops and recreates the affected table

datadog_monitors.tags
  Source:      list<item: utf8, nullable> → json
  postgresql:  text[] → jsonb
  Safe mode:   rejects this change
  Forced mode: drops and recreates the affected table

Action: plan a manual migration or rebuild.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name: "automatically migratable with table level behavior",
			report: upgradeReport{
				SourceName:  "okta",
				FromVersion: "v6.8.2",
				ToVersion:   "v7.0.0",
				Destination: postgresqlDestinationSpec(specs.WriteModeAppend),
				Findings: []*pluginPb.AssessTables_TableFinding{{
					TableName:        "okta_policy_rules",
					Category:         pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE,
					SafeModeBehavior: "adds the columns; no table recreation required",
					Columns: []*pluginPb.AssessTables_ColumnFinding{
						{ColumnName: "policy_id", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, NewType: "text"},
						{ColumnName: "actions", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, NewType: "jsonb"},
					},
				}},
			},
			want: `okta v6.8.2 → v7.0.0 | postgresql (cloudquery/postgresql@v8.14.0)
write_mode: append | pk_mode: default
AUTOMATICALLY MIGRATABLE — okta_policy_rules

okta_policy_rules.policy_id
  postgresql:  none → text

okta_policy_rules.actions
  postgresql:  none → jsonb

okta_policy_rules
  Safe mode:   adds the columns; no table recreation required

Action: use safe migration.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name: "no changes",
			report: upgradeReport{
				SourceName:  "datadog",
				FromVersion: "v5.19.10",
				ToVersion:   "v6.0.0",
				Destination: postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale),
				Findings:    []*pluginPb.AssessTables_TableFinding{{TableName: "datadog_users", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE}},
			},
			want: `datadog v5.19.10 → v6.0.0 | postgresql (cloudquery/postgresql@v8.14.0)
write_mode: overwrite-delete-stale | pk_mode: default
No schema changes affect your selected tables.

`,
		},
		{
			name: "removed tables, unknown tables and a source listed with a connection",
			report: upgradeReport{
				SourceName:    "gcp",
				FromVersion:   "v22.1.2",
				ToVersion:     "v23.0.0",
				SourceGaps:    []string{"gcp v23.0.0: tables were listed with a connection (metadata only, no rows read)"},
				RemovedTables: []string{"gcp_aiplatform_specialist_pools", "gcp_aiplatform_specialistpool_locations"},
				Destination:   postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale),
				Findings: []*pluginPb.AssessTables_TableFinding{
					{TableName: "gcp_compute_instances", Category: pluginPb.AssessTables_CATEGORY_UNKNOWN, CoverageIncomplete: true, CoverageIncompleteReason: destinationNoAssessmentReason},
					{TableName: "gcp_storage_buckets", Category: pluginPb.AssessTables_CATEGORY_UNKNOWN, CoverageIncomplete: true},
				},
			},
			want: `gcp v22.1.2 → v23.0.0 | postgresql (cloudquery/postgresql@v8.14.0)
write_mode: overwrite-delete-stale | pk_mode: default
SELECTED TABLES REMOVED — 2 tables

gcp_aiplatform_specialist_pools
gcp_aiplatform_specialistpool_locations
The new source version no longer provides these tables.

UNKNOWN — 2 tables

Coverage gaps:
  gcp v23.0.0: tables were listed with a connection (metadata only, no rows read)
  gcp_compute_instances: destination version does not support assessment
  gcp_storage_buckets: coverage incomplete

Action: remove explicit selections and update dependent consumers.
Action: review the source changelog for what this check could not assess.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name: "source tables could not be listed",
			report: upgradeReport{
				SourceName:    "s3",
				FromVersion:   "v1.0.0",
				ToVersion:     "v2.0.0",
				SourceUnknown: true,
				SourceGaps:    []string{"s3 v1.0.0: tables could not be listed: source returned no tables with or without a connection"},
			},
			want: `s3 v1.0.0 → v2.0.0
UNKNOWN

Coverage gaps:
  s3 v1.0.0: tables could not be listed: source returned no tables with or without a connection

Action: review the source changelog for what this check could not assess.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
	}
	setColorOutput(t, false)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			require.NoError(t, renderUpgradeReport(&out, tc.report))
			require.Equal(t, tc.want, out.String())
		})
	}
}

func TestRenderUpgradeReportWithColor(t *testing.T) {
	setColorOutput(t, true)
	destination := postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale)
	cases := []struct {
		name     string
		report   upgradeReport
		contains []string
	}{
		{
			name: "manual migration",
			report: upgradeReport{
				SourceName: "datadog", FromVersion: "v5.19.10", ToVersion: "v6.0.0", Destination: destination,
				RemovedTables: []string{"datadog_removed"},
				SourceGaps:    []string{"datadog v6.0.0: tables were listed with a connection (metadata only, no rows read)"},
				Findings: []*pluginPb.AssessTables_TableFinding{{
					TableName: "datadog_monitors",
					Category:  pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED,
					Columns: []*pluginPb.AssessTables_ColumnFinding{{
						ColumnName:         "tags",
						Category:           pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED,
						OldType:            "text[]",
						NewType:            "jsonb",
						SafeModeBehavior:   "rejects the changes",
						ForcedModeBehavior: "drops and recreates the table",
					}},
				}},
			},
			contains: []string{
				"\x1b[1mdatadog v5.19.10 → v6.0.0 | \x1b[22m\x1b[1;36mpostgresql (cloudquery/postgresql@v8.14.0)\x1b[22;0m\n",
				"\x1b[2mwrite_mode: overwrite-delete-stale | pk_mode: default\x1b[22m\n",
				"\x1b[31mREVIEW REQUIRED — datadog_monitors\x1b[0m\n",
				"\x1b[33mSELECTED TABLES REMOVED — datadog_removed\x1b[0m\n",
				"\n\x1b[1mdatadog_removed\x1b[22m\n",
				"\n\x1b[1mdatadog_monitors.tags\x1b[22m\n",
				"  \x1b[2mpostgresql: \x1b[22m \x1b[31mtext[]\x1b[0m → \x1b[32mjsonb\x1b[0m\n",
				"  \x1b[2mSafe mode:  \x1b[22m \x1b[31mrejects the changes\x1b[0m\n",
				"  \x1b[2mForced mode:\x1b[22m \x1b[31mdrops and recreates the table\x1b[0m\n",
				"  \x1b[33mdatadog v6.0.0: tables were listed with a connection (metadata only, no rows read)\x1b[0m\n",
				"\x1b[1;31mAction: plan a manual migration or rebuild.\x1b[22;0m\n",
				"\x1b[1;33mAction: remove explicit selections and update dependent consumers.\x1b[22;0m\n",
				"\x1b[2mThis check only previews the changes. It does not migrate, write, delete or upload anything.\x1b[22m\n",
			},
		},
		{
			name: "automatically migratable and unknown",
			report: upgradeReport{
				SourceName: "okta", FromVersion: "v6.8.2", ToVersion: "v7.0.0", Destination: destination,
				Findings: []*pluginPb.AssessTables_TableFinding{
					{
						TableName:          "okta_policy_rules",
						Category:           pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE,
						SafeModeBehavior:   "adds the columns",
						ForcedModeBehavior: "adds the columns",
					},
					{TableName: "okta_users", Category: pluginPb.AssessTables_CATEGORY_UNKNOWN},
				},
			},
			contains: []string{
				"\x1b[33mUNKNOWN — okta_users\x1b[0m\n",
				"\x1b[33mAUTOMATICALLY MIGRATABLE — okta_policy_rules\x1b[0m\n",
				"\n\x1b[1mokta_policy_rules\x1b[22m\n",
				"  \x1b[2mSafe mode:  \x1b[22m \x1b[32madds the columns\x1b[0m\n",
				"  \x1b[2mForced mode:\x1b[22m \x1b[33madds the columns\x1b[0m\n",
				"\x1b[1;33mAction: use safe migration.\x1b[22;0m\n",
			},
		},
		{
			name: "no changes",
			report: upgradeReport{
				SourceName: "datadog", FromVersion: "v5.19.10", ToVersion: "v6.0.0", Destination: destination,
				Findings: []*pluginPb.AssessTables_TableFinding{{TableName: "datadog_users", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE}},
			},
			contains: []string{"\x1b[32mNo schema changes affect your selected tables.\x1b[0m\n"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			require.NoError(t, renderUpgradeReport(&out, tc.report))
			for _, want := range tc.contains {
				require.Contains(t, out.String(), want)
			}
		})
	}
}

func TestUpgradeTablesByName(t *testing.T) {
	pairs := []upgradeTablePair{
		{Name: "test_kept", From: testTable("test_kept"), To: testTable("test_kept")},
		{Name: "test_added", To: testTable("test_added")},
	}
	tables, err := transformUpgradeTables(t.Context(), specs.Source{Metadata: specs.Metadata{Name: "test"}}, specs.Destination{Metadata: specs.Metadata{Name: "postgresql"}}, []pluginPb.PluginClient{renamingTransformerClient{}}, false, pairs)
	require.NoError(t, err)

	byName, err := upgradeTablesByName(tables)
	require.NoError(t, err)

	require.Len(t, byName, 2)
	require.Equal(t, "renamed_test_kept", byName["renamed_test_kept"].From.Name)
	require.Equal(t, "renamed_test_kept", byName["renamed_test_kept"].To.Name)
	require.Nil(t, byName["renamed_test_added"].From)
	require.Equal(t, "renamed_test_added", byName["renamed_test_added"].To.Name)
}
