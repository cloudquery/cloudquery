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

func postgresqlDestinationSpec(writeMode specs.WriteMode, migrateMode specs.MigrateMode) *specs.Destination {
	return &specs.Destination{
		Metadata:    specs.Metadata{Name: "postgresql", Path: "cloudquery/postgresql", Version: "v8.14.0", Registry: specs.RegistryCloudQuery},
		WriteMode:   writeMode,
		MigrateMode: migrateMode,
	}
}

func s3DestinationSpec() *specs.Destination {
	return &specs.Destination{Metadata: specs.Metadata{Name: "s3", Path: "cloudquery/s3", Version: "v7.0.0", Registry: specs.RegistryCloudQuery}}
}

func datadogTagsReport(destination *specs.Destination) upgradeReport {
	tagsTable := func(name string, tagsType arrow.DataType) *schema.Table {
		return &schema.Table{Name: name, Columns: schema.ColumnList{{Name: "tags", Type: tagsType}}}
	}
	tagsFinding := func(table, column string) *pluginPb.AssessTables_TableFinding {
		return &pluginPb.AssessTables_TableFinding{
			TableName:          table,
			Category:           pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED,
			SafeModeBehavior:   "rejects the changes",
			ForcedModeBehavior: "drops and recreates the table, deleting existing rows",
			Columns: []*pluginPb.AssessTables_ColumnFinding{
				{ColumnName: "id", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE, OldType: "text", NewType: "text"},
				{ColumnName: column, Category: pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED, OldType: "text[]", NewType: "jsonb"},
			},
		}
	}
	return upgradeReport{
		SourceName:  "datadog",
		FromVersion: "v5.19.10",
		ToVersion:   "v6.0.0",
		Destination: destination,
		Tables: map[string]upgradeTablePair{
			"datadog_monitors":   {From: tagsTable("datadog_monitors", arrow.ListOf(arrow.BinaryTypes.String)), To: tagsTable("datadog_monitors", types.ExtensionTypes.JSON)},
			"datadog_dashboards": {From: tagsTable("datadog_dashboards", arrow.ListOf(arrow.BinaryTypes.String)), To: tagsTable("datadog_dashboards", types.ExtensionTypes.JSON)},
		},
		Findings: []*pluginPb.AssessTables_TableFinding{
			tagsFinding("datadog_monitors", "tags"),
			tagsFinding("datadog_dashboards", "monitor_tags"),
			{TableName: "datadog_users", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE},
		},
	}
}

func oktaReport(destination *specs.Destination, policyRulesCategory pluginPb.AssessTables_Category, policyIDIsPrimaryKey bool) upgradeReport {
	id := schema.Column{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: policyIDIsPrimaryKey}
	newColumn := func(name string, category pluginPb.AssessTables_Category, newType string) *pluginPb.AssessTables_ColumnFinding {
		return &pluginPb.AssessTables_ColumnFinding{ColumnName: name, Category: category, NewType: newType}
	}
	return upgradeReport{
		SourceName:  "okta",
		FromVersion: "v6.8.2",
		ToVersion:   "v7.0.0",
		Destination: destination,
		Tables: map[string]upgradeTablePair{
			"okta_policy_rules": {
				From: &schema.Table{Name: "okta_policy_rules", Columns: schema.ColumnList{id}},
				To: &schema.Table{Name: "okta_policy_rules", Columns: schema.ColumnList{
					id,
					{Name: "policy_id", Type: arrow.BinaryTypes.String, PrimaryKey: policyIDIsPrimaryKey},
					{Name: "actions", Type: types.ExtensionTypes.JSON},
					{Name: "conditions", Type: types.ExtensionTypes.JSON},
				}},
			},
			"okta_policy_mappings": {To: &schema.Table{Name: "okta_policy_mappings", Columns: schema.ColumnList{id}}},
		},
		Findings: []*pluginPb.AssessTables_TableFinding{
			{
				TableName: "okta_policy_rules",
				Category:  policyRulesCategory,
				Columns: []*pluginPb.AssessTables_ColumnFinding{
					newColumn("policy_id", policyRulesCategory, "text"),
					newColumn("actions", pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, "jsonb"),
					newColumn("conditions", pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, "jsonb"),
				},
			},
			{TableName: "okta_policy_mappings", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE},
		},
	}
}

func s3JSONReport() upgradeReport {
	tagsFinding := func(table, column string) *pluginPb.AssessTables_TableFinding {
		return &pluginPb.AssessTables_TableFinding{
			TableName: table,
			Category:  pluginPb.AssessTables_CATEGORY_NO_CHANGE,
			Columns: []*pluginPb.AssessTables_ColumnFinding{
				{ColumnName: "id", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE},
				{
					ColumnName: column,
					Category:   pluginPb.AssessTables_CATEGORY_NO_CHANGE,
					OldType:    "list<item: utf8, nullable>",
					NewType:    "json",
					Evidence: []*pluginPb.AssessTables_Evidence{
						{SyntheticValue: `["env:prod"]`, Before: `{"` + column + `":["env:prod"]}`, After: `{"` + column + `":["env:prod"]}`},
						{SyntheticValue: `null`, Before: `{"` + column + `":null}`, After: `{"` + column + `":null}`},
					},
				},
			},
		}
	}
	return upgradeReport{
		SourceName:  "datadog",
		FromVersion: "v5.19.10",
		ToVersion:   "v6.0.0",
		Destination: s3DestinationSpec(),
		Findings: []*pluginPb.AssessTables_TableFinding{
			{TableName: "datadog_users", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE},
			tagsFinding("datadog_monitors", "tags"),
			tagsFinding("datadog_downtimes", "monitor_tags"),
			tagsFinding("datadog_slos", "tags"),
			tagsFinding("datadog_synthetics", "tags"),
		},
	}
}

func TestRenderUpgradeReport(t *testing.T) {
	cases := []struct {
		name   string
		report upgradeReport
		want   string
	}{
		{
			name:   "type change fails in safe mode",
			report: datadogTagsReport(postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe)),
			want: `datadog v5.19.10 → v6.0.0 | postgresql (cloudquery/postgresql@v8.14.0)
write_mode: overwrite-delete-stale | pk_mode: default | migrate_mode: safe
REVIEW REQUIRED — 2 tables need a manual migration

Changes
  datadog_dashboards
    ~ monitor_tags   text[] → jsonb   type changed
  datadog_monitors
    ~ tags           text[] → jsonb   type changed

Next sync
  migrate_mode: safe (your config)
    ✗ datadog_dashboards   fails: safe mode cannot change a column type
    ✗ datadog_monitors     fails: safe mode cannot change a column type
  migrate_mode: forced
    ! datadog_dashboards   dropped and recreated, existing rows deleted
    ! datadog_monitors     dropped and recreated, existing rows deleted

Action: migrate datadog_dashboards, datadog_monitors manually before upgrading, or switch to migrate_mode: forced and accept losing their rows.
Run again with --ai-prompt to get a prompt for an AI agent that guides a gradual migration without data loss.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name:   "type change deletes rows in forced mode",
			report: datadogTagsReport(postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeForced)),
			want: `datadog v5.19.10 → v6.0.0 | postgresql (cloudquery/postgresql@v8.14.0)
write_mode: overwrite-delete-stale | pk_mode: default | migrate_mode: forced
REVIEW REQUIRED — 2 tables will be recreated, deleting existing rows

Changes
  datadog_dashboards
    ~ monitor_tags   text[] → jsonb   type changed
  datadog_monitors
    ~ tags           text[] → jsonb   type changed

Next sync
  migrate_mode: forced (your config)
    ! datadog_dashboards   dropped and recreated, existing rows deleted
    ! datadog_monitors     dropped and recreated, existing rows deleted
  migrate_mode: safe
    ✗ datadog_dashboards   fails: safe mode cannot change a column type
    ✗ datadog_monitors     fails: safe mode cannot change a column type

Action: the next sync drops and recreates datadog_dashboards, datadog_monitors and deletes existing rows; back up any data you need before upgrading.
Run again with --ai-prompt to get a prompt for an AI agent that guides a gradual migration without data loss.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name:   "primary key addition with a new table",
			report: oktaReport(postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe), pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED, true),
			want: `okta v6.8.2 → v7.0.0 | postgresql (cloudquery/postgresql@v8.14.0)
write_mode: overwrite-delete-stale | pk_mode: default | migrate_mode: safe
REVIEW REQUIRED — 1 table needs a manual migration, 1 new table

Changes
  okta_policy_mappings   new table
  okta_policy_rules
    + policy_id    text    new column, part of the primary key
    + actions      jsonb   new column
    + conditions   jsonb   new column

Next sync
  migrate_mode: safe (your config)
    ✓ okta_policy_mappings   created
    ✗ okta_policy_rules      fails: safe mode cannot change a primary key
  migrate_mode: forced
    ✓ okta_policy_mappings   created
    ! okta_policy_rules      dropped and recreated, existing rows deleted

Action: migrate okta_policy_rules manually before upgrading, or switch to migrate_mode: forced and accept losing its rows.
Run again with --ai-prompt to get a prompt for an AI agent that guides a gradual migration without data loss.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name:   "append mode applies every change in both modes",
			report: oktaReport(postgresqlDestinationSpec(specs.WriteModeAppend, specs.MigrateModeSafe), pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, false),
			want: `okta v6.8.2 → v7.0.0 | postgresql (cloudquery/postgresql@v8.14.0)
write_mode: append | pk_mode: default | migrate_mode: safe
AUTOMATICALLY MIGRATABLE — 1 changed table, 1 new table

Changes
  okta_policy_mappings   new table
  okta_policy_rules
    + policy_id    text    new column
    + actions      jsonb   new column
    + conditions   jsonb   new column

Next sync
  migrate_mode: safe (your config)
    ✓ okta_policy_mappings   created
    ✓ okta_policy_rules      migrated in place
  migrate_mode: forced — same as safe

Action: use safe migration.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name: "removed tables are sorted, with unknown tables and a source listed with a connection",
			report: upgradeReport{
				SourceName:    "gcp",
				FromVersion:   "v22.1.2",
				ToVersion:     "v23.0.0",
				SourceGaps:    []string{"gcp v23.0.0: tables were listed with a connection (metadata only, no rows read)"},
				RemovedTables: []string{"gcp_aiplatform_specialistpool_locations", "gcp_aiplatform_specialist_pools"},
				Destination:   postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe),
				Findings: []*pluginPb.AssessTables_TableFinding{
					{TableName: "gcp_storage_buckets", Category: pluginPb.AssessTables_CATEGORY_UNKNOWN, CoverageIncomplete: true},
					{TableName: "gcp_compute_instances", Category: pluginPb.AssessTables_CATEGORY_UNKNOWN, CoverageIncomplete: true, CoverageIncompleteReason: destinationNoAssessmentReason},
				},
			},
			want: `gcp v22.1.2 → v23.0.0 | postgresql (cloudquery/postgresql@v8.14.0)
write_mode: overwrite-delete-stale | pk_mode: default
SELECTED TABLES REMOVED — 2 removed tables, 2 tables not assessed

Changes
  gcp_aiplatform_specialist_pools           removed table, the new source version no longer provides it
  gcp_aiplatform_specialistpool_locations   removed table, the new source version no longer provides it

Coverage gaps
  gcp v23.0.0: tables were listed with a connection (metadata only, no rows read)
  gcp_compute_instances: destination version does not support assessment
  gcp_storage_buckets: coverage incomplete

Action: remove explicit selections and update dependent consumers.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name: "file schema changed",
			report: upgradeReport{
				SourceName:  "cloudflare",
				FromVersion: "v11.4.0",
				ToVersion:   "v12.0.0",
				Destination: s3DestinationSpec(),
				Findings: []*pluginPb.AssessTables_TableFinding{{
					TableName: "cloudflare_certificate_packs",
					Category:  pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED,
					Columns: []*pluginPb.AssessTables_ColumnFinding{
						{ColumnName: "created_on", Category: pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED, OldType: "TIMESTAMP(MICROS)"},
						{ColumnName: "validation_records", Category: pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED, NewType: "BYTE_ARRAY (JSON)"},
					},
				}},
			},
			want: `cloudflare v11.4.0 → v12.0.0 | s3 (cloudquery/s3@v7.0.0)
write_mode: overwrite-delete-stale | pk_mode: default
FILE SCHEMA CHANGED — 1 changed table

Changes
  cloudflare_certificate_packs
    - created_on           TIMESTAMP(MICROS)   column removed
    + validation_records   BYTE_ARRAY (JSON)   new column

Action: review readers that combine old and new files.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name:   "file output unchanged for equivalent test values",
			report: s3JSONReport(),
			want: `datadog v5.19.10 → v6.0.0 | s3 (cloudquery/s3@v7.0.0)
write_mode: overwrite-delete-stale | pk_mode: default
NO OUTPUT DIFFERENCE DETECTED for equivalent test values

Output comparison
  list<item: utf8, nullable> → json   4 columns: datadog_downtimes.monitor_tags, datadog_monitors.tags, datadog_slos.tags, …
    Synthetic value: ["env:prod"]
    Before:          {"monitor_tags":["env:prod"]}
    After:           {"monitor_tags":["env:prod"]}
    2 equivalent test values: identical output
  NO OUTPUT DIFFERENCE DETECTED for equivalent test values
  Actual source values and behavior were not assessed.

Action: no action needed; the output is the same for equivalent values.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name: "output that differs is listed in full",
			report: upgradeReport{
				SourceName:  "datadog",
				FromVersion: "v5.19.10",
				ToVersion:   "v6.0.0",
				Destination: s3DestinationSpec(),
				Findings: []*pluginPb.AssessTables_TableFinding{{
					TableName: "datadog_monitors",
					Category:  pluginPb.AssessTables_CATEGORY_NO_CHANGE,
					Columns: []*pluginPb.AssessTables_ColumnFinding{{
						ColumnName: "tags",
						Category:   pluginPb.AssessTables_CATEGORY_NO_CHANGE,
						OldType:    "list<item: utf8, nullable>",
						NewType:    "json",
						Evidence: []*pluginPb.AssessTables_Evidence{
							{SyntheticValue: `["env:prod"]`, Before: `{"tags":["env:prod"]}`, After: `{"tags":["env:prod"]}`},
							{SyntheticValue: `[]`, Before: `{"tags":[]}`, After: `{"tags":null}`},
						},
					}},
				}},
			},
			want: `datadog v5.19.10 → v6.0.0 | s3 (cloudquery/s3@v7.0.0)
write_mode: overwrite-delete-stale | pk_mode: default
NO OUTPUT DIFFERENCE DETECTED for equivalent test values

Output comparison
  list<item: utf8, nullable> → json   datadog_monitors.tags
    Synthetic value: ["env:prod"]
    Before:          {"tags":["env:prod"]}
    After:           {"tags":["env:prod"]}
    Synthetic value: []
    Before:          {"tags":[]}
    After:           {"tags":null}
  NO OUTPUT DIFFERENCE DETECTED for equivalent test values
  Actual source values and behavior were not assessed.

Action: no action needed; the output is the same for equivalent values.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name: "no changes",
			report: upgradeReport{
				SourceName:  "datadog",
				FromVersion: "v5.19.10",
				ToVersion:   "v6.0.0",
				Destination: postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe),
				Findings:    []*pluginPb.AssessTables_TableFinding{{TableName: "datadog_users", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE}},
			},
			want: `datadog v5.19.10 → v6.0.0 | postgresql (cloudquery/postgresql@v8.14.0)
write_mode: overwrite-delete-stale | pk_mode: default
No schema changes affect your selected tables.

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

Coverage gaps
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
			require.NoError(t, renderUpgradeReport(&out, tc.report, false))
			require.Equal(t, tc.want, out.String())
		})
	}
}

func TestRenderUpgradeReportWithColor(t *testing.T) {
	setColorOutput(t, true)
	cases := []struct {
		name     string
		report   upgradeReport
		contains []string
	}{
		{
			name:   "manual migration",
			report: oktaReport(postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe), pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED, true),
			contains: []string{
				"\x1b[1mokta v6.8.2 → v7.0.0 | \x1b[22m\x1b[1;36mpostgresql (cloudquery/postgresql@v8.14.0)\x1b[22;0m\n",
				"\x1b[2mwrite_mode: overwrite-delete-stale | pk_mode: default | migrate_mode: safe\x1b[22m\n",
				"\x1b[31mREVIEW REQUIRED — 1 table needs a manual migration, 1 new table\x1b[0m\n",
				"  \x1b[1mokta_policy_mappings\x1b[22m   \x1b[32mnew table\x1b[0m\n",
				"  \x1b[1mokta_policy_rules\x1b[22m\n",
				"    \x1b[32m+ policy_id    text    new column, part of the primary key\x1b[0m\n",
				"  \x1b[1mmigrate_mode: safe (your config)\x1b[22m\n",
				"    \x1b[32m✓ okta_policy_mappings   created\x1b[0m\n",
				"    \x1b[31m✗ okta_policy_rules      fails: safe mode cannot change a primary key\x1b[0m\n",
				"    \x1b[33m! okta_policy_rules      dropped and recreated, existing rows deleted\x1b[0m\n",
				"\x1b[1;31mAction: migrate okta_policy_rules manually before upgrading, or switch to migrate_mode: forced and accept losing its rows.\x1b[22;0m\n",
				"\x1b[2mThis check only previews the changes. It does not migrate, write, delete or upload anything.\x1b[22m\n",
			},
		},
		{
			name:     "type change",
			report:   datadogTagsReport(postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe)),
			contains: []string{"    \x1b[33m~ tags           text[] → jsonb   type changed\x1b[0m\n"},
		},
		{
			name: "removed tables, coverage gaps and file changes",
			report: upgradeReport{
				SourceName: "gcp", FromVersion: "v22.1.2", ToVersion: "v23.0.0", Destination: s3DestinationSpec(),
				RemovedTables: []string{"gcp_removed"},
				SourceGaps:    []string{"gcp v23.0.0: tables were listed with a connection (metadata only, no rows read)"},
				Findings: []*pluginPb.AssessTables_TableFinding{{
					TableName: "gcp_buckets",
					Category:  pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED,
					Columns:   []*pluginPb.AssessTables_ColumnFinding{{ColumnName: "created_on", Category: pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED, OldType: "INT64"}},
				}},
			},
			contains: []string{
				"\x1b[33mSELECTED TABLES REMOVED — 1 changed table, 1 removed table\x1b[0m\n",
				"    \x1b[31m- created_on   INT64   column removed\x1b[0m\n",
				"  \x1b[1mgcp_removed\x1b[22m   \x1b[31mremoved table, the new source version no longer provides it\x1b[0m\n",
				"  \x1b[33mgcp v23.0.0: tables were listed with a connection (metadata only, no rows read)\x1b[0m\n",
				"\x1b[1;33mAction: remove explicit selections and update dependent consumers.\x1b[22;0m\n",
			},
		},
		{
			name:   "output comparison",
			report: s3JSONReport(),
			contains: []string{
				"\x1b[32mNO OUTPUT DIFFERENCE DETECTED for equivalent test values\x1b[0m\n",
				"  list<item: utf8, nullable> → json   \x1b[1m4 columns: datadog_downtimes.monitor_tags, datadog_monitors.tags, datadog_slos.tags, …\x1b[22m\n",
				"    \x1b[2mSynthetic value:\x1b[22m [\"env:prod\"]\n",
				"\x1b[1;32mAction: no action needed; the output is the same for equivalent values.\x1b[22;0m\n",
			},
		},
		{
			name: "no changes",
			report: upgradeReport{
				SourceName: "datadog", FromVersion: "v5.19.10", ToVersion: "v6.0.0", Destination: s3DestinationSpec(),
				Findings: []*pluginPb.AssessTables_TableFinding{{TableName: "datadog_users", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE}},
			},
			contains: []string{"\x1b[32mNo schema changes affect your selected tables.\x1b[0m\n"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			require.NoError(t, renderUpgradeReport(&out, tc.report, false))
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
