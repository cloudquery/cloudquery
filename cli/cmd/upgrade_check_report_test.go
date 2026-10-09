package cmd

import (
	"bytes"
	"fmt"
	"strings"
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
	return &specs.Destination{
		Metadata:  specs.Metadata{Name: "s3", Path: "cloudquery/s3", Version: "v7.0.0", Registry: specs.RegistryCloudQuery},
		WriteMode: specs.WriteModeAppend,
		Spec:      map[string]any{"format": "parquet"},
	}
}

func k8sParquetReport(nodeNameType string) upgradeReport {
	nodes := func(columns ...schema.Column) *schema.Table {
		return &schema.Table{Name: "k8s_core_nodes", Columns: append(schema.ColumnList{{Name: "uid", Type: arrow.BinaryTypes.String, NotNull: true}}, columns...)}
	}
	return upgradeReport{
		SourceName:  "k8s",
		FromVersion: "v7.9.8",
		ToVersion:   "v8.4.1",
		Destination: s3DestinationSpec(),
		Tables: map[string]upgradeTablePair{
			"k8s_core_nodes": {
				From: nodes(),
				To: nodes(
					schema.Column{Name: "cloud_cluster_id", Type: arrow.BinaryTypes.String},
					schema.Column{Name: "spec_taints", Type: arrow.ListOf(arrow.BinaryTypes.String)},
					schema.Column{Name: "node_name", Type: arrow.BinaryTypes.String},
				),
			},
			"k8s_openshift_routes": {To: &schema.Table{Name: "k8s_openshift_routes", Columns: schema.ColumnList{{Name: "uid", Type: arrow.BinaryTypes.String}}}},
		},
		Findings: []*pluginPb.AssessTables_TableFinding{
			{
				TableName:          "k8s_core_nodes",
				Category:           pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE,
				SafeModeBehavior:   "new files add the new columns, existing files are not changed",
				ForcedModeBehavior: "new files add the new columns, existing files are not changed",
				Columns: []*pluginPb.AssessTables_ColumnFinding{
					{ColumnName: "cloud_cluster_id", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, NewType: "optional byte_array (String)"},
					{ColumnName: "spec_taints", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, NewType: "optional group (List) {list: repeated group {element: optional byte_array (String)}}"},
					{ColumnName: "node_name", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, NewType: nodeNameType},
				},
			},
			{
				TableName:          "k8s_openshift_routes",
				Category:           pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE,
				SafeModeBehavior:   "new files add the new table",
				ForcedModeBehavior: "new files add the new table",
				Columns:            []*pluginPb.AssessTables_ColumnFinding{{ColumnName: "uid", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, NewType: "optional byte_array (String)"}},
			},
		},
	}
}

func semgrepParquetReport() upgradeReport {
	findings := func(columns ...schema.Column) *schema.Table {
		return &schema.Table{Name: "semgrep_findings", Columns: columns}
	}
	return upgradeReport{
		SourceName:  "semgrep",
		FromVersion: "v1.0.5",
		ToVersion:   "v3.1.0",
		Destination: s3DestinationSpec(),
		Tables: map[string]upgradeTablePair{
			"semgrep_findings": {
				From: findings(
					schema.Column{Name: "id", Type: arrow.PrimitiveTypes.Float64},
					schema.Column{Name: "created_at", Type: arrow.FixedWidthTypes.Timestamp_us},
					schema.Column{Name: "categories", Type: arrow.ListOf(arrow.BinaryTypes.String)},
				),
				To: findings(
					schema.Column{Name: "id", Type: arrow.PrimitiveTypes.Int64},
					schema.Column{Name: "created_at", Type: arrow.FixedWidthTypes.Timestamp_us},
					schema.Column{Name: "categories", Type: arrow.BinaryTypes.String},
				),
			},
		},
		Findings: []*pluginPb.AssessTables_TableFinding{{
			TableName: "semgrep_findings",
			Category:  pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED,
			Columns: []*pluginPb.AssessTables_ColumnFinding{
				{ColumnName: "id", Category: pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED, OldType: "optional double", NewType: "optional int64 (Int(bitWidth=64, isSigned=true))"},
				{
					ColumnName: "categories",
					Category:   pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED,
					OldType:    "optional group (List) {list: repeated group {element: optional byte_array (String)}}",
					NewType:    "optional byte_array (String)",
				},
				{
					ColumnName: "first_seen_at",
					Category:   pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED,
					NewType:    "required int64 (Timestamp(isAdjustedToUTC=true, timeUnit=microseconds, is_from_converted_type=false, force_set_converted_type=true))",
				},
			},
		}},
	}
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
	tables := map[string]upgradeTablePair{}
	for table, column := range map[string]string{"datadog_monitors": "tags", "datadog_downtimes": "monitor_tags", "datadog_slos": "tags", "datadog_synthetics": "tags"} {
		tables[table] = upgradeTablePair{
			From: &schema.Table{Name: table, Columns: schema.ColumnList{{Name: column, Type: arrow.ListOf(arrow.BinaryTypes.String)}}},
			To:   &schema.Table{Name: table, Columns: schema.ColumnList{{Name: column, Type: types.ExtensionTypes.JSON}}},
		}
	}
	return upgradeReport{
		SourceName:  "datadog",
		FromVersion: "v5.19.10",
		ToVersion:   "v6.0.0",
		Destination: s3DestinationSpec(),
		Tables:      tables,
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
    ~ tags           text[] → jsonb   type changed (list<string> → json)

Next sync
  migrate_mode: safe (your config)
    ✗ datadog_dashboards   fails: safe mode cannot change a column type
    ✗ datadog_monitors     fails: safe mode cannot change a column type
  migrate_mode: forced
    ! datadog_dashboards   dropped and recreated, existing rows deleted
    ! datadog_monitors     dropped and recreated, existing rows deleted

Action: migrate datadog_dashboards, datadog_monitors manually before upgrading, or switch to migrate_mode: forced and accept losing their rows.
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
    ~ tags           text[] → jsonb   type changed (list<string> → json)

Next sync
  migrate_mode: forced (your config)
    ! datadog_dashboards   dropped and recreated, existing rows deleted
    ! datadog_monitors     dropped and recreated, existing rows deleted
  migrate_mode: safe
    ✗ datadog_dashboards   fails: safe mode cannot change a column type
    ✗ datadog_monitors     fails: safe mode cannot change a column type

Action: the next sync drops and recreates datadog_dashboards, datadog_monitors and deletes existing rows; back up any data you need before upgrading.
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
  okta_policy_rules
    + policy_id    text    new column, part of the primary key (string)
    + actions      jsonb   new column (json)
    + conditions   jsonb   new column (json)
  okta_policy_mappings   new table

Next sync
  migrate_mode: safe (your config)
    ✗ okta_policy_rules      fails: safe mode cannot change a primary key
    ✓ okta_policy_mappings   created
  migrate_mode: forced
    ! okta_policy_rules      dropped and recreated, existing rows deleted
    ✓ okta_policy_mappings   created

Action: migrate okta_policy_rules manually before upgrading, or switch to migrate_mode: forced and accept losing its rows.
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
  okta_policy_rules
    + actions      jsonb   new column (json)
    + conditions   jsonb   new column (json)
    + policy_id    text    new column (string)
  okta_policy_mappings   new table

Next sync
  migrate_mode: safe (your config)
    ✓ okta_policy_rules      migrated in place
    ✓ okta_policy_mappings   created
  migrate_mode: forced — same as safe

Action: use safe migration.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name: "removed tables are sorted, with unknown tables and a source listed with a connection",
			report: upgradeReport{
				SourceName:  "gcp",
				FromVersion: "v22.1.2",
				ToVersion:   "v23.0.0",
				SourceGaps:  []string{"gcp v23.0.0: tables were listed with a connection (metadata only, no rows read)"},
				RemovedTables: []upgradeRemovedTable{
					{Name: "gcp_aiplatform_specialistpool_locations", SelectedBy: upgradeSelectedByName},
					{Name: "gcp_aiplatform_specialist_pools", SelectedBy: upgradeSelectedByName},
				},
				Destination: postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe),
				Findings: []*pluginPb.AssessTables_TableFinding{
					{TableName: "gcp_storage_buckets", Category: pluginPb.AssessTables_CATEGORY_UNKNOWN, IncompleteCoverageReason: "nested types not assessed"},
					{TableName: "gcp_compute_instances", Category: pluginPb.AssessTables_CATEGORY_UNKNOWN, IncompleteCoverageReason: destinationNoAssessmentReason},
				},
			},
			want: `gcp v22.1.2 → v23.0.0 | postgresql (cloudquery/postgresql@v8.14.0)
write_mode: overwrite-delete-stale | pk_mode: default
SELECTED TABLES REMOVED — 2 removed tables, 2 tables not fully assessed

Changes
  gcp_aiplatform_specialist_pools           removed table, the new source version no longer provides it
  gcp_aiplatform_specialistpool_locations   removed table, the new source version no longer provides it

Coverage gaps
  gcp v23.0.0: tables were listed with a connection (metadata only, no rows read)
  gcp_compute_instances: destination version does not support assessment
  gcp_storage_buckets: nested types not assessed

Action: remove explicit selections (gcp_aiplatform_specialist_pools, gcp_aiplatform_specialistpool_locations) and update dependent consumers.
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
write_mode: append | pk_mode: default
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
			name:   "new nullable file columns and new tables extend the file schema",
			report: k8sParquetReport("optional byte_array (String)"),
			want: `k8s v7.9.8 → v8.4.1 | s3 (cloudquery/s3@v7.0.0)
write_mode: append | pk_mode: default
FILE SCHEMA EXTENDED — 1 changed table, 1 new table

Changes
  k8s_core_nodes
    + cloud_cluster_id   string         new column (string)
    + node_name          string         new column (string)
    + spec_taints        list<string>   new column (list<string>)
  k8s_openshift_routes   new table

Action: no action needed; new files add these columns and tables. Readers that match columns by position must be updated.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name:   "new file tables only extend the file schema",
			report: func() upgradeReport { r := k8sParquetReport(""); r.Findings = r.Findings[1:]; return r }(),
			want: `k8s v7.9.8 → v8.4.1 | s3 (cloudquery/s3@v7.0.0)
write_mode: append | pk_mode: default
FILE SCHEMA EXTENDED — 1 new table

Changes
  k8s_openshift_routes   new table

Action: no action needed; new files add these tables.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name:   "file type changes show short parquet types",
			report: semgrepParquetReport(),
			want: `semgrep v1.0.5 → v3.1.0 | s3 (cloudquery/s3@v7.0.0)
write_mode: append | pk_mode: default
FILE SCHEMA CHANGED — 1 changed table

Changes
  semgrep_findings
    ~ categories      list<string> → string          type changed (list<string> → string)
    ~ id              double → int64                 type changed (float64 → int64)
    + first_seen_at   required timestamp (us, UTC)   new column

Action: review readers that combine old and new files.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name: "file type changes that shorten to the same name show full parquet types",
			report: upgradeReport{
				SourceName: "aws", FromVersion: "v1.0.0", ToVersion: "v2.0.0", Destination: s3DestinationSpec(),
				Findings: []*pluginPb.AssessTables_TableFinding{{
					TableName: "aws_ec2_instances",
					Category:  pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED,
					Columns: []*pluginPb.AssessTables_ColumnFinding{{
						ColumnName: "tags",
						Category:   pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED,
						OldType:    "optional group (List) {list: repeated group {element: optional byte_array (String)}}",
						NewType:    "optional group (List) {list: repeated group {element: required byte_array (String)}}",
					}},
				}},
			},
			want: `aws v1.0.0 → v2.0.0 | s3 (cloudquery/s3@v7.0.0)
write_mode: append | pk_mode: default
FILE SCHEMA CHANGED — 1 changed table

Changes
  aws_ec2_instances
    ~ tags   optional group (List) {list: repeated group {element: optional byte_array (String)}} → optional group (List) {list: repeated group {element: required byte_array (String)}}   type changed

Action: review readers that combine old and new files.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name:   "file output unchanged for equivalent test values",
			report: s3JSONReport(),
			want: `datadog v5.19.10 → v6.0.0 | s3 (cloudquery/s3@v7.0.0)
write_mode: append | pk_mode: default
NO OUTPUT DIFFERENCE DETECTED for equivalent test values

Output comparison
  list<string> → json   4 columns: datadog_downtimes.monitor_tags, datadog_monitors.tags, datadog_slos.tags, …
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
			name: "output that differs needs review",
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
write_mode: append | pk_mode: default
FILE OUTPUT CHANGED

Output comparison
  list<item: utf8, nullable> → json   datadog_monitors.tags
    Synthetic value: ["env:prod"]
    Before:          {"tags":["env:prod"]}
    After:           {"tags":["env:prod"]}
    Synthetic value: []
    Before:          {"tags":[]}
    After:           {"tags":null}
  OUTPUT DIFFERENCE DETECTED for some test values
  Actual source values and behavior were not assessed.

Action: review readers that combine old and new files.
This check only previews the changes. It does not migrate, write, delete or upload anything.

`,
		},
		{
			name: "incomplete coverage is not approved",
			report: upgradeReport{
				SourceName:  "okta",
				FromVersion: "v6.8.2",
				ToVersion:   "v7.0.0",
				Destination: postgresqlDestinationSpec(specs.WriteModeAppend, specs.MigrateModeSafe),
				Findings: []*pluginPb.AssessTables_TableFinding{{
					TableName:                "okta_users",
					Category:                 pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE,
					IncompleteCoverageReason: "nested types not assessed",
					Columns:                  []*pluginPb.AssessTables_ColumnFinding{{ColumnName: "profile", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, NewType: "jsonb"}},
				}},
			},
			want: `okta v6.8.2 → v7.0.0 | postgresql (cloudquery/postgresql@v8.14.0)
write_mode: append | pk_mode: default | migrate_mode: safe
UNKNOWN — 1 changed table, 1 table not fully assessed

Changes
  okta_users
    + profile   jsonb   new column

Next sync
  migrate_mode: safe (your config)
    ✓ okta_users   migrated in place
  migrate_mode: forced — same as safe

Coverage gaps
  okta_users: nested types not assessed

Action: review the source changelog for what this check could not assess.
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
		{
			name: "source from a local binary compared with the hub",
			report: upgradeReport{
				SourceName: "semgrep", FromVersion: "v3.0.0", ToVersion: "v3.1.0", SourceUnknown: true,
				FromOrigin: upgradeSourceOrigin{Registry: specs.RegistryLocal, Path: "/opt/plugins/semgrep"},
				ToOrigin:   upgradeSourceOrigin{Registry: specs.RegistryCloudQuery, Path: "cloudquery/semgrep"},
				SourceGaps: upgradeSourceGaps("semgrep",
					upgradeSourceTables{Origin: upgradeSourceOrigin{Registry: specs.RegistryLocal, Path: "/opt/plugins/semgrep"}, Version: "v3.0.0", UnknownReason: "source returned no tables with or without a connection"},
					upgradeSourceTables{Origin: upgradeSourceOrigin{Registry: specs.RegistryCloudQuery, Path: "cloudquery/semgrep"}, Version: "v3.1.0", Connected: true},
				),
			},
			want: `semgrep (local: /opt/plugins/semgrep, v3.0.0) → cloudquery/semgrep@v3.1.0
UNKNOWN

Coverage gaps
  semgrep (local: /opt/plugins/semgrep, v3.0.0): tables could not be listed: source returned no tables with or without a connection
  semgrep cloudquery/semgrep@v3.1.0: tables were listed with a connection (metadata only, no rows read)

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
				"    \x1b[31m+ policy_id    text    new column, part of the primary key (string)\x1b[0m\n",
				"    \x1b[32m+ actions      jsonb   new column (json)\x1b[0m\n",
				"  \x1b[1mmigrate_mode: safe (your config)\x1b[22m\n",
				"    \x1b[32m✓ okta_policy_mappings   created\x1b[0m\n",
				"    \x1b[31m✗ okta_policy_rules      fails: safe mode cannot change a primary key\x1b[0m\n",
				"    \x1b[33m! okta_policy_rules      dropped and recreated, existing rows deleted\x1b[0m\n",
				"\x1b[1;31mAction: migrate okta_policy_rules manually before upgrading, or switch to migrate_mode: forced and accept losing its rows.\x1b[22;0m\n",
				"\x1b[2mThis check only previews the changes. It does not migrate, write, delete or upload anything.\x1b[22m\n",
			},
		},
		{
			name:   "append mode applies every change",
			report: oktaReport(postgresqlDestinationSpec(specs.WriteModeAppend, specs.MigrateModeSafe), pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, false),
			contains: []string{
				"    \x1b[32m+ policy_id    text    new column (string)\x1b[0m\n",
				"  \x1b[1mokta_policy_mappings\x1b[22m   \x1b[32mnew table\x1b[0m\n",
			},
		},
		{
			name: "removed column on a database destination",
			report: upgradeReport{
				SourceName: "okta", FromVersion: "v6.8.2", ToVersion: "v7.0.0", Destination: postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe),
				Findings: []*pluginPb.AssessTables_TableFinding{{
					TableName: "okta_users",
					Category:  pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE,
					Columns:   []*pluginPb.AssessTables_ColumnFinding{{ColumnName: "legacy", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, OldType: "text"}},
				}},
			},
			contains: []string{"    \x1b[33m- legacy   text   column removed\x1b[0m\n"},
		},
		{
			name:     "type change",
			report:   datadogTagsReport(postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe)),
			contains: []string{"    \x1b[31m~ tags           text[] → jsonb   type changed (list<string> → json)\x1b[0m\n"},
		},
		{
			name: "removed tables, coverage gaps and file changes",
			report: upgradeReport{
				SourceName: "gcp", FromVersion: "v22.1.2", ToVersion: "v23.0.0", Destination: s3DestinationSpec(),
				RemovedTables: []upgradeRemovedTable{{Name: "gcp_removed", SelectedBy: upgradeSelectedByPattern}},
				SourceGaps:    []string{"gcp v23.0.0: tables were listed with a connection (metadata only, no rows read)"},
				Findings: []*pluginPb.AssessTables_TableFinding{{
					TableName: "gcp_buckets",
					Category:  pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED,
					Columns:   []*pluginPb.AssessTables_ColumnFinding{{ColumnName: "created_on", Category: pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED, OldType: "INT64"}},
				}},
			},
			contains: []string{
				"\x1b[33mSELECTED TABLES REMOVED — 1 removed table, 1 changed table\x1b[0m\n",
				"    \x1b[33m- created_on   INT64   column removed\x1b[0m\n",
				"  \x1b[1mgcp_removed\x1b[22m   \x1b[31mremoved table, the new source version no longer provides it\x1b[0m\n",
				"  \x1b[33mgcp v23.0.0: tables were listed with a connection (metadata only, no rows read)\x1b[0m\n",
				"\x1b[1;33mAction: update dependent consumers; the removed tables stop syncing, and their existing data in the destination is kept but no longer updated.\x1b[22;0m\n",
			},
		},
		{
			name:   "output comparison",
			report: s3JSONReport(),
			contains: []string{
				"\x1b[32mNO OUTPUT DIFFERENCE DETECTED for equivalent test values\x1b[0m\n",
				"  list<string> → json   \x1b[1m4 columns: datadog_downtimes.monitor_tags, datadog_monitors.tags, datadog_slos.tags, …\x1b[22m\n",
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
			require.NoError(t, renderUpgradeReport(&out, tc.report))
			for _, want := range tc.contains {
				require.Contains(t, out.String(), want)
			}
		})
	}
}

func TestUpgradeRemovedTablesAction(t *testing.T) {
	const stopSyncing = "update dependent consumers; the removed tables stop syncing, and their existing data in the destination is kept but no longer updated."
	byName := func(name string) upgradeRemovedTable {
		return upgradeRemovedTable{Name: name, SelectedBy: upgradeSelectedByName}
	}
	cases := []struct {
		name   string
		report upgradeReport
		want   string
	}{
		{
			name:   "exact name",
			report: upgradeReport{RemovedTables: []upgradeRemovedTable{byName("gcp_firebaseappcheck_safety_net_configs")}},
			want:   "remove explicit selections (gcp_firebaseappcheck_safety_net_configs) and update dependent consumers.",
		},
		{
			name:   "wildcard",
			report: upgradeReport{RemovedTables: []upgradeRemovedTable{{Name: "gcp_firebaseappcheck_safety_net_configs", SelectedBy: upgradeSelectedByPattern}}},
			want:   stopSyncing,
		},
		{
			name: "glob prefix",
			report: upgradeReport{RemovedTables: []upgradeRemovedTable{
				{Name: "aws_ec2_removed", SelectedBy: upgradeSelectedByPattern},
				{Name: "aws_ec2_old", SelectedBy: upgradeSelectedByPattern},
			}},
			want: stopSyncing,
		},
		{
			name: "glob prefix with no match in the new version",
			report: upgradeReport{
				RemovedTables:          []upgradeRemovedTable{{Name: "aws_iam_old", SelectedBy: upgradeSelectedByPattern}},
				UnmatchedTablePatterns: []string{"aws_iam_*"},
			},
			want: "remove explicit selections (aws_iam_*) and update dependent consumers.",
		},
		{
			name:   "dependent table",
			report: upgradeReport{RemovedTables: []upgradeRemovedTable{{Name: "aws_s3_bucket_grants", SelectedBy: upgradeSelectedAsDependent}}},
			want:   stopSyncing,
		},
		{
			name: "mix",
			report: upgradeReport{RemovedTables: []upgradeRemovedTable{
				byName("aws_ec2_removed"),
				{Name: "aws_ec2_old", SelectedBy: upgradeSelectedByPattern},
				{Name: "aws_s3_bucket_grants", SelectedBy: upgradeSelectedAsDependent},
			}},
			want: "remove explicit selections (aws_ec2_removed) and update dependent consumers.",
		},
		{
			name: "many exact names",
			report: upgradeReport{
				RemovedTables:          []upgradeRemovedTable{byName("t1"), byName("t2"), byName("t3"), byName("t4")},
				UnmatchedTablePatterns: []string{"t_*"},
			},
			want: "remove explicit selections (t_*, t1, t2, t3 and 1 more) and update dependent consumers.",
		},
		{
			name: "unmatched patterns are never shortened",
			report: upgradeReport{
				RemovedTables:          []upgradeRemovedTable{byName("t1"), byName("t2"), byName("t3")},
				UnmatchedTablePatterns: []string{"a_*", "b_*", "c_*", "d_*"},
			},
			want: "remove explicit selections (a_*, b_*, c_*, d_*, t1, t2, t3) and update dependent consumers.",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.report.Destination = postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe)
			action, _ := upgradeAction(tc.report, upgradeTableImpacts(tc.report), nil, specs.MigrateModeSafe)
			require.Equal(t, tc.want, action)
			require.Equal(t, tc.want, upgradeReportToJSON(tc.report).Action)
		})
	}
}

func TestRenderUpgradeReportCapsTypeColumnWidth(t *testing.T) {
	setColorOutput(t, false)
	report := upgradeReport{
		SourceName: "aws", FromVersion: "v1.0.0", ToVersion: "v2.0.0", Destination: s3DestinationSpec(),
		Findings: []*pluginPb.AssessTables_TableFinding{{
			TableName: "aws_ec2_instances",
			Category:  pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED,
			Columns: []*pluginPb.AssessTables_ColumnFinding{
				{ColumnName: "cpu", Category: pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED, OldType: "Nullable(Float64)", NewType: "Array(Nullable(Map(String, Nullable(String))))"},
				{ColumnName: "name", Category: pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED, NewType: "LowCardinality(Nullable(String))"},
			},
		}},
	}

	var out bytes.Buffer
	require.NoError(t, renderUpgradeReport(&out, report))

	require.Contains(t, out.String(), "    ~ cpu    Nullable(Float64) → Array(Nullable(Map(String, Nullable(String))))   type changed\n")
	require.Contains(t, out.String(), "    + name   LowCardinality(Nullable(String))"+strings.Repeat(" ", upgradeMaxTypeWidth-len("LowCardinality(Nullable(String))"))+"   new column\n")
}

func upgradeOrderingReport(extraChangedTables int) upgradeReport {
	added := func(column string) *pluginPb.AssessTables_ColumnFinding {
		return &pluginPb.AssessTables_ColumnFinding{ColumnName: column, Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, NewType: "text"}
	}
	findings := make([]*pluginPb.AssessTables_TableFinding, 0, 3+extraChangedTables)
	findings = append(findings,
		&pluginPb.AssessTables_TableFinding{TableName: "a_new", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, Columns: []*pluginPb.AssessTables_ColumnFinding{added("id")}},
		&pluginPb.AssessTables_TableFinding{TableName: "b_changed", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, Columns: []*pluginPb.AssessTables_ColumnFinding{added("b")}},
		&pluginPb.AssessTables_TableFinding{
			TableName: "c_manual",
			Category:  pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED,
			Columns: []*pluginPb.AssessTables_ColumnFinding{
				added("a_added"),
				{ColumnName: "z_type", Category: pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED, OldType: "text", NewType: "jsonb"},
				{ColumnName: "m_removed", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, OldType: "text"},
			},
		},
	)
	for i := range extraChangedTables {
		findings = append(findings, &pluginPb.AssessTables_TableFinding{TableName: fmt.Sprintf("d_changed_%02d", i), Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, Columns: []*pluginPb.AssessTables_ColumnFinding{added("d")}})
	}
	return upgradeReport{
		SourceName: "gcp", FromVersion: "v23.4.1", ToVersion: "v24.1.1",
		Destination:   postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe),
		RemovedTables: []upgradeRemovedTable{{Name: "z_removed", SelectedBy: upgradeSelectedByPattern}},
		Tables:        map[string]upgradeTablePair{"a_new": {To: testTable("a_new")}},
		Findings:      findings,
	}
}

func TestRenderUpgradeReportPutsBreakingChangesFirst(t *testing.T) {
	setColorOutput(t, false)
	var out bytes.Buffer
	require.NoError(t, renderUpgradeReport(&out, upgradeOrderingReport(0)))

	require.Contains(t, out.String(), `
Changes
  z_removed   removed table, the new source version no longer provides it
  c_manual
    - m_removed   text           column removed
    ~ z_type      text → jsonb   type changed
    + a_added     text           new column
  b_changed
    + b           text           new column
  a_new       new table

Next sync
  migrate_mode: safe (your config)
    ✗ c_manual    fails: safe mode cannot change a column type
    ✓ b_changed   migrated in place
    ✓ a_new       created
  migrate_mode: forced
    ! c_manual    dropped and recreated, existing rows deleted
    ✓ b_changed   migrated in place
    ✓ a_new       created
`)
	require.NotContains(t, out.String(), "Breaking:")

	report := upgradeReportToJSON(upgradeOrderingReport(0))
	require.Equal(t, []string{"c_manual", "b_changed", "a_new"}, upgradeImpactNames(report.Tables))
	require.Equal(t, "m_removed", report.Tables[0].Changes[0].Column)
}

func TestRenderUpgradeReportBreakingSummary(t *testing.T) {
	setColorOutput(t, false)
	var out bytes.Buffer
	require.NoError(t, renderUpgradeReport(&out, upgradeOrderingReport(upgradeLongChangeList)))

	require.Contains(t, out.String(), "REVIEW REQUIRED — 1 removed table, 1 table needs a manual migration, 11 changed tables, 1 new table\nBreaking: 1 removed table (z_removed); 1 table needs a manual migration (c_manual)\n\nChanges\n")

	report := upgradeOrderingReport(upgradeLongChangeList)
	report.Findings = report.Findings[3:]
	out.Reset()
	require.NoError(t, renderUpgradeReport(&out, report))
	require.Contains(t, out.String(), "Breaking: 1 removed table (z_removed)\n")

	report.RemovedTables = nil
	out.Reset()
	require.NoError(t, renderUpgradeReport(&out, report))
	require.NotContains(t, out.String(), "Breaking:")
}

func TestUpgradeOutputComparisonsPutDifferencesFirst(t *testing.T) {
	setColorOutput(t, false)
	evidence := func(after string) []*pluginPb.AssessTables_Evidence {
		return []*pluginPb.AssessTables_Evidence{{SyntheticValue: "1", Before: "1", After: after}}
	}
	report := upgradeReport{
		SourceName: "aws", FromVersion: "v1.0.0", ToVersion: "v2.0.0", Destination: s3DestinationSpec(),
		Findings: []*pluginPb.AssessTables_TableFinding{
			{TableName: "a_same", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE, Columns: []*pluginPb.AssessTables_ColumnFinding{{ColumnName: "c", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE, OldType: "int64", NewType: "string", Evidence: evidence("1")}}},
			{TableName: "b_differs", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE, Columns: []*pluginPb.AssessTables_ColumnFinding{{ColumnName: "c", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE, OldType: "int64", NewType: "json", Evidence: evidence(`"1"`)}}},
		},
	}

	var out bytes.Buffer
	require.NoError(t, renderUpgradeReport(&out, report))
	require.Less(t, strings.Index(out.String(), "b_differs.c"), strings.Index(out.String(), "a_same.c"))

	comparisons := upgradeReportToJSON(report).OutputComparisons
	require.Equal(t, []string{"b_differs", "a_same"}, []string{comparisons[0].Table, comparisons[1].Table})
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

func TestUpgradeArrowTypeName(t *testing.T) {
	cases := map[arrow.DataType]string{
		arrow.BinaryTypes.String:                                          "string",
		arrow.BinaryTypes.LargeString:                                     "string",
		arrow.ListOf(arrow.BinaryTypes.String):                            "list<string>",
		arrow.LargeListOf(arrow.ListOf(arrow.BinaryTypes.String)):         "list<list<string>>",
		types.ExtensionTypes.JSON:                                         "json",
		arrow.PrimitiveTypes.Int64:                                        "int64",
		arrow.MapOf(arrow.BinaryTypes.String, arrow.PrimitiveTypes.Int64): arrow.MapOf(arrow.BinaryTypes.String, arrow.PrimitiveTypes.Int64).String(),
	}
	for dataType, want := range cases {
		require.Equal(t, want, upgradeArrowTypeName(dataType))
	}
}

func TestUpgradeExitCode(t *testing.T) {
	postgresql := postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe)
	cases := []struct {
		name   string
		report upgradeReport
		want   int
	}{
		{name: "no change", report: upgradeReport{Destination: postgresql, Findings: []*pluginPb.AssessTables_TableFinding{{TableName: "t", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE}}}, want: 0},
		{name: "automatically migratable", report: oktaReport(postgresqlDestinationSpec(specs.WriteModeAppend, specs.MigrateModeSafe), pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, false), want: 0},
		{name: "no output difference", report: s3JSONReport(), want: 0},
		{name: "source listed with a connection", report: upgradeReport{Destination: postgresql, SourceGaps: []string{"listed with a connection"}}, want: 0},
		{name: "manual migration", report: datadogTagsReport(postgresql), want: 3},
		{name: "manual migration in forced mode", report: datadogTagsReport(postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeForced)), want: 3},
		{name: "selected tables removed", report: upgradeReport{Destination: postgresql, RemovedTables: []upgradeRemovedTable{{Name: "t", SelectedBy: upgradeSelectedByName}}}, want: 3},
		{name: "tables removed from a wildcard selection", report: upgradeReport{Destination: postgresql, RemovedTables: []upgradeRemovedTable{{Name: "t", SelectedBy: upgradeSelectedByPattern}}}, want: 3},
		{name: "dependent tables removed", report: upgradeReport{Destination: postgresql, RemovedTables: []upgradeRemovedTable{{Name: "t", SelectedBy: upgradeSelectedAsDependent}}}, want: 3},
		{name: "file schema extended", report: k8sParquetReport("optional byte_array (String)"), want: 0},
		{name: "file schema changed with short types", report: semgrepParquetReport(), want: 3},
		{name: "file schema changed", report: upgradeReport{Destination: s3DestinationSpec(), Findings: []*pluginPb.AssessTables_TableFinding{{TableName: "t", Category: pluginPb.AssessTables_CATEGORY_FILE_SCHEMA_CHANGED}}}, want: 3},
		{name: "output differs", report: upgradeReport{Destination: s3DestinationSpec(), Findings: []*pluginPb.AssessTables_TableFinding{{
			TableName: "t",
			Category:  pluginPb.AssessTables_CATEGORY_NO_CHANGE,
			Columns:   []*pluginPb.AssessTables_ColumnFinding{{ColumnName: "c", Category: pluginPb.AssessTables_CATEGORY_NO_CHANGE, Evidence: []*pluginPb.AssessTables_Evidence{{SyntheticValue: "[]", Before: "[]", After: "null"}}}},
		}}}, want: 3},
		{name: "unknown destination", report: upgradeReport{Destination: postgresql, Findings: []*pluginPb.AssessTables_TableFinding{{TableName: "t", Category: pluginPb.AssessTables_CATEGORY_UNKNOWN}}}, want: 4},
		{name: "incomplete coverage", report: upgradeReport{Destination: postgresql, Findings: []*pluginPb.AssessTables_TableFinding{{TableName: "t", Category: pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, IncompleteCoverageReason: "nested types not assessed"}}}, want: 4},
		{name: "unknown wins over action needed", report: upgradeReport{Destination: postgresql, RemovedTables: []upgradeRemovedTable{{Name: "t", SelectedBy: upgradeSelectedByPattern}}, Findings: []*pluginPb.AssessTables_TableFinding{{TableName: "u", Category: pluginPb.AssessTables_CATEGORY_UNKNOWN}}}, want: 4},
		{name: "source tables could not be listed", report: upgradeReport{SourceUnknown: true}, want: 4},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, upgradeExitCode(tc.report))
		})
	}
}
