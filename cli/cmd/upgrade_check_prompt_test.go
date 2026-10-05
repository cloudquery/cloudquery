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

func TestUpgradeAIPromptNeverIncludesSpecValues(t *testing.T) {
	destination := postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe)
	destination.DockerRegistryAuthToken = "docker-registry-token"
	destination.Spec = map[string]any{
		"connection_string": "postgres://admin:s3cr3t-password@db.internal:5432/cloudquery?sslmode=require",
		"pgx_log_level":     "trace-level-value",
		"nested":            map[string]any{"token": "nested-secret-token"},
	}

	for _, report := range []upgradeReport{
		datadogTagsReport(destination),
		oktaReport(destination, pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED, true),
	} {
		var text bytes.Buffer
		require.NoError(t, renderUpgradeReports(&text, upgradeOutputText, true, []upgradeReport{report}))
		var jsonOut bytes.Buffer
		require.NoError(t, renderUpgradeReports(&jsonOut, upgradeOutputJSON, false, []upgradeReport{report}))
		prompt := upgradeAIPrompt(report)
		require.NotEmpty(t, prompt)

		for _, secret := range []string{"s3cr3t-password", "db.internal", "admin:", "trace-level-value", "nested-secret-token", "docker-registry-token", "connection_string"} {
			require.NotContains(t, prompt, secret)
			require.NotContains(t, text.String(), secret)
			require.NotContains(t, jsonOut.String(), secret)
		}
	}
}

func TestUpgradeAIPrompt(t *testing.T) {
	cases := []struct {
		name        string
		report      upgradeReport
		contains    []string
		notContains []string
	}{
		{
			name:   "type change",
			report: datadogTagsReport(postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe)),
			contains: []string{
				"You are helping me migrate tables in my PostgreSQL database",
				"- Source plugin: datadog, upgrade from v5.19.10 to v6.0.0\n",
				"- Destination plugin: postgresql (cloudquery/postgresql@v8.14.0)\n",
				"- write_mode: overwrite-delete-stale, pk_mode: default, migrate_mode: safe\n",
				"### datadog_dashboards\n",
				"  - `monitor_tags`: `text[]` → `jsonb`\n",
				"### datadog_monitors\n\n- Primary key: none (unchanged)\n",
				"- Why `migrate_mode: safe` rejects it: safe mode cannot change a column type\n",
				"`_cq_id`, `_cq_parent_id` (on child tables), `_cq_source_name` and `_cq_sync_time`",
				"deletes the rows of the same `_cq_source_name` that have an older `_cq_sync_time`",
				"Pause all syncs",
				"- Type conversions:",
				"USING to_jsonb(c)",
				"ACCESS EXCLUSIVE",
				"- Rollback:",
			},
			notContains: []string{"datadog_users", "- Primary-key changes:", "CREATE UNIQUE INDEX CONCURRENTLY", "migrate_mode: forced`, the first sync"},
		},
		{
			name:   "primary-key change in forced mode",
			report: oktaReport(postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeForced), pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED, true),
			contains: []string{
				"### okta_policy_rules\n",
				"- Primary key now: `id`\n",
				"- Primary key after the upgrade: `id`, `policy_id`\n",
				"  - `policy_id`: new `text` column, NOT NULL, part of the primary key\n",
				"  - `actions`: new `jsonb` column\n",
				"- Why `migrate_mode: safe` rejects it: safe mode cannot change a primary key\n",
				"With `migrate_mode: forced`, the first sync of the new version drops and recreates these tables",
				"- Primary-key changes:",
				"CREATE UNIQUE INDEX CONCURRENTLY",
			},
			notContains: []string{"okta_policy_mappings", "- Type conversions:", "USING to_jsonb(c)"},
		},
		{
			name:     "append write mode",
			report:   datadogTagsReport(postgresqlDestinationSpec(specs.WriteModeAppend, specs.MigrateModeSafe)),
			contains: []string{"With `write_mode: append`, each sync only inserts rows"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			prompt := upgradeAIPrompt(tc.report)
			for _, want := range tc.contains {
				require.Contains(t, prompt, want)
			}
			for _, unwanted := range tc.notContains {
				require.NotContains(t, prompt, unwanted)
			}
			require.Equal(t, prompt, upgradeAIPrompt(tc.report))
		})
	}
}

func TestUpgradeAIPromptOnlyForManualMigrations(t *testing.T) {
	cases := []struct {
		name   string
		report upgradeReport
	}{
		{name: "automatically migratable", report: oktaReport(postgresqlDestinationSpec(specs.WriteModeAppend, specs.MigrateModeSafe), pluginPb.AssessTables_CATEGORY_AUTOMATICALLY_MIGRATABLE, false)},
		{name: "output comparison", report: s3JSONReport()},
		{name: "source unknown", report: upgradeReport{SourceName: "s3", FromVersion: "v1.0.0", ToVersion: "v2.0.0", SourceUnknown: true}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Empty(t, upgradeAIPrompt(tc.report))

			var text bytes.Buffer
			require.NoError(t, renderUpgradeReports(&text, upgradeOutputText, false, []upgradeReport{tc.report}))
			require.NotContains(t, text.String(), upgradeAIPromptHintText)

			text.Reset()
			require.NoError(t, renderUpgradeReports(&text, upgradeOutputText, true, []upgradeReport{tc.report}))
			require.True(t, strings.HasSuffix(text.String(), "\n"+upgradeNoManualText+"\n\n"))
			require.NotContains(t, text.String(), "AI migration prompt")

			var raw struct {
				Reports []map[string]json.RawMessage `json:"reports"`
			}
			var jsonOut bytes.Buffer
			require.NoError(t, renderUpgradeReportsJSON(&jsonOut, []upgradeReport{tc.report}))
			require.NoError(t, json.Unmarshal(jsonOut.Bytes(), &raw))
			require.NotContains(t, raw.Reports[0], "ai_prompt")
			require.NotContains(t, raw.Reports[0], "prompt_version")
		})
	}
}

func TestRenderUpgradeReportsAIPrompt(t *testing.T) {
	setColorOutput(t, false)
	report := datadogTagsReport(postgresqlDestinationSpec(specs.WriteModeOverwriteDeleteStale, specs.MigrateModeSafe))

	var withoutFlag bytes.Buffer
	require.NoError(t, renderUpgradeReports(&withoutFlag, upgradeOutputText, false, []upgradeReport{report}))
	require.Contains(t, withoutFlag.String(), "Action: migrate datadog_dashboards, datadog_monitors manually before upgrading, or switch to migrate_mode: forced and accept losing their rows.\n"+upgradeAIPromptHintText+"\n")
	require.NotContains(t, withoutFlag.String(), "AI migration prompt")

	var withFlag bytes.Buffer
	require.NoError(t, renderUpgradeReports(&withFlag, upgradeOutputText, true, []upgradeReport{report}))
	require.NotContains(t, withFlag.String(), upgradeAIPromptHintText)
	reportText, prompt, found := strings.Cut(withFlag.String(), "--- AI migration prompt for postgresql (prompt_version "+upgradeAIPromptVersion+") ---\n")
	require.True(t, found)
	require.Equal(t, strings.Replace(withoutFlag.String(), upgradeAIPromptHintText+"\n", "", 1), reportText)
	require.Equal(t, upgradeAIPrompt(report)+upgradeAIPromptEndText+"\n\n", prompt)

	got := renderUpgradeReportJSON(t, report)
	require.Equal(t, upgradeAIPrompt(report), got.AIPrompt)
	require.Equal(t, upgradeAIPromptVersion, got.PromptVersion)
}

func TestUpgradeDatabaseGuidanceFor(t *testing.T) {
	cases := []struct {
		registry specs.Registry
		path     string
		want     string
	}{
		{specs.RegistryCloudQuery, "cloudquery/postgresql", "PostgreSQL"},
		{specs.RegistryCloudQuery, "cloudquery/mysql", "MySQL"},
		{specs.RegistryCloudQuery, "cloudquery/mssql", "SQL Server"},
		{specs.RegistryCloudQuery, "cloudquery/sqlite", "SQLite"},
		{specs.RegistryCloudQuery, "cloudquery/duckdb", "DuckDB"},
		{specs.RegistryLocal, "/plugins/destination/postgresql/cq-destination-postgresql", "PostgreSQL"},
		{specs.RegistryDocker, "ghcr.io/cloudquery/cq-destination-mysql:v5.0.0", "MySQL"},
		{specs.RegistryCloudQuery, "cloudquery/snowflake", "snowflake"},
	}
	for _, tc := range cases {
		t.Run(tc.path, func(t *testing.T) {
			guidance := upgradeDatabaseGuidanceFor(&specs.Destination{Metadata: specs.Metadata{Registry: tc.registry, Path: tc.path}})
			require.Equal(t, tc.want, guidance.name)
			require.NotEmpty(t, guidance.general)
		})
	}
}
