package cmd

import (
	"fmt"
	"io"
	"path/filepath"
	"slices"
	"strings"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const (
	upgradeAIPromptVersion  = "1"
	upgradeAIPromptHintText = "Run again with --ai-prompt to get a prompt for an AI agent that guides a gradual migration without data loss."
	upgradeNoManualText     = "No manual migration needed."
	upgradeAIPromptEndText  = "--- end ---"
)

type upgradeDatabaseGuidance struct {
	plugin     string
	name       string
	general    []string
	typeChange []string
	primaryKey []string
}

var upgradeDatabaseGuidances = []upgradeDatabaseGuidance{
	{
		plugin: "postgresql",
		name:   "PostgreSQL",
		general: []string{
			"`ALTER TABLE` takes an ACCESS EXCLUSIVE lock that blocks all reads and writes on the table. Set `lock_timeout` (for example `SET lock_timeout = '5s'`) so that a waiting `ALTER TABLE` does not block the queries behind it, and retry off-peak.",
			"A table rewrite needs free disk space for a full copy of the table and its indexes, and writes WAL that can increase replication lag. Check sizes with `pg_total_relation_size` and lag with `pg_stat_replication`.",
			"Views that depend on a column block `ALTER COLUMN ... TYPE`. Find them in `pg_depend` and recreate them in the same transaction.",
			"CloudQuery names the primary-key constraint `<table>_cqpk`. Keep this name.",
		},
		typeChange: []string{
			"`ALTER TABLE t ALTER COLUMN c TYPE jsonb USING to_jsonb(c)` converts `text[]` to `jsonb` in one statement, but rewrites the table while it holds the ACCESS EXCLUSIVE lock. Use a single statement like this only for small tables.",
			"For large tables, add a new column with the target type, backfill it in batches (for example 1,000 to 10,000 rows per transaction, by primary-key ranges) with the conversion expression, then drop the old column and rename the new one in one short transaction. Run `VACUUM (ANALYZE)` after the backfill.",
		},
		primaryKey: []string{
			"Add new key columns as nullable, backfill them in batches, then add `CHECK (col IS NOT NULL) NOT VALID`, run `VALIDATE CONSTRAINT`, and then `SET NOT NULL` (PostgreSQL 12 and later skip the full table scan when a validated check constraint exists).",
			"Build the new key with `CREATE UNIQUE INDEX CONCURRENTLY` (it cannot run in a transaction), then swap in one short transaction: `ALTER TABLE t DROP CONSTRAINT t_cqpk, ADD CONSTRAINT t_cqpk PRIMARY KEY USING INDEX <new_index>`.",
		},
	},
	{
		plugin: "mysql",
		name:   "MySQL",
		general: []string{
			"Most column type and primary-key changes rebuild an InnoDB table. Add `ALGORITHM=INPLACE, LOCK=NONE` so that MySQL fails instead of blocking writes when it cannot make the change online.",
			"DDL waits for a metadata lock behind open transactions and blocks new queries while it waits. Set a low `lock_wait_timeout` for the session.",
			"For large tables, consider an online schema change tool such as gh-ost or pt-online-schema-change.",
			"A replica applies a large `ALTER TABLE` as one statement, so replication lags for about as long as the change takes. A rebuild also needs free disk space for a full copy of the table.",
		},
		typeChange: []string{
			"Convert with an explicit expression, for example `CAST(col AS JSON)` for text that already holds valid JSON. Check converted values on a sample with `JSON_VALID` before you change the column.",
		},
		primaryKey: []string{
			"Change the primary key in one statement: `ALTER TABLE t DROP PRIMARY KEY, ADD PRIMARY KEY (...)`. The primary key is the clustered index, so this rebuilds the table. Make new key columns NOT NULL first.",
		},
	},
	{
		plugin: "mssql",
		name:   "SQL Server",
		general: []string{
			"`ALTER TABLE ... ALTER COLUMN` takes a schema modification (Sch-M) lock that blocks all access to the table, and a type change can update every row. Use `SET LOCK_TIMEOUT` for the session.",
			"Large updates fill the transaction log and slow down Always On secondaries. Backfill in batches with `UPDATE TOP (n)` and check log space with `sys.dm_db_log_space_usage`.",
			"CloudQuery names the primary-key constraint `<table>_cqpk`. Keep this name.",
		},
		typeChange: []string{
			"For large tables, add a new column with the target type, backfill it in batches, then swap the names with `sp_rename` in one short transaction. Check converted values, for example with `ISJSON` for JSON text.",
		},
		primaryKey: []string{
			"Build the new key as a unique index with `ONLINE = ON` where your edition supports it, then drop and recreate the primary-key constraint in one transaction. A clustered primary key rebuild rewrites the table.",
		},
	},
	{
		plugin: "sqlite",
		name:   "SQLite",
		general: []string{
			"SQLite has one writer at a time and locks the whole database during a write transaction. Stop all other writers first.",
			"Back up the database file before the first change, for example with `VACUUM INTO 'backup.db'`.",
			"`ALTER TABLE` cannot change a column type or a primary key. Use the table rebuild procedure from the SQLite documentation: in one transaction with `PRAGMA foreign_keys=OFF`, create a new table with the target schema, copy the rows with `INSERT INTO new_table SELECT ...` and conversion expressions, drop the old table, rename the new one, and recreate its indexes. This needs free disk space for a second copy of the table.",
		},
		typeChange: []string{
			"SQLite stores JSON as text. Check converted values with `json_valid`.",
		},
	},
	{
		plugin: "duckdb",
		name:   "DuckDB",
		general: []string{
			"DuckDB has one writer process at a time. Stop CloudQuery and all other writers first.",
			"Back up the database file before the first change, for example with `EXPORT DATABASE` or a copy of the file while no process has it open.",
			"DuckDB cannot alter a column that an index or a primary key depends on. When `ALTER TABLE` fails for this reason, create a new table with the target schema, copy the rows with `INSERT INTO new_table SELECT ...` and conversion expressions, then swap the names.",
		},
		typeChange: []string{
			"`ALTER TABLE t ALTER COLUMN c TYPE new_type USING expression` converts a column in place when no index depends on it. Test the expression with a `SELECT` on a sample first.",
		},
		primaryKey: []string{
			"To change a primary key, create a new table with the target primary key and copy the rows.",
		},
	},
}

var upgradeGenericDatabaseGuidance = upgradeDatabaseGuidance{
	general: []string{
		"Find out which schema changes this database runs online and which ones rewrite or lock the table. Use online schema change features where they exist.",
		"Back up the affected tables before the first write.",
	},
	typeChange: []string{
		"Prefer a new column with the target type, a batched backfill with an explicit conversion expression, and a short swap of the column names.",
	},
	primaryKey: []string{
		"Prefer a new unique index or a new table with the target key, built without blocking writes, and a short swap.",
	},
}

func upgradeDatabaseGuidanceFor(destination *specs.Destination) upgradeDatabaseGuidance {
	plugin := strings.ToLower(filepath.Base(destination.Path))
	for _, guidance := range upgradeDatabaseGuidances {
		if strings.Contains(plugin, guidance.plugin) {
			return guidance
		}
	}
	guidance := upgradeGenericDatabaseGuidance
	guidance.name = plugin
	return guidance
}

func upgradeManualImpacts(impacts []upgradeTableImpact) []upgradeTableImpact {
	return slices.DeleteFunc(slices.Clone(impacts), func(impact upgradeTableImpact) bool {
		return impact.Category != pluginPb.AssessTables_CATEGORY_MANUAL_MIGRATION_REQUIRED
	})
}

func upgradeAIPrompt(r upgradeReport) string {
	r = sortUpgradeReport(r)
	manual := upgradeManualImpacts(upgradeTableImpacts(r))
	if r.Destination == nil || len(manual) == 0 {
		return ""
	}
	d := r.Destination
	database := upgradeDatabaseGuidanceFor(d)
	hasTypeChange, hasPKChange := false, false
	for _, impact := range manual {
		table := r.Tables[impact.Name]
		hasPKChange = hasPKChange || !slices.Equal(upgradePrimaryKeys(table.From), upgradePrimaryKeys(table.To))
		hasTypeChange = hasTypeChange || slices.ContainsFunc(impact.Changes, func(change upgradeColumnChange) bool {
			return change.Kind == upgradeColumnChanged && change.OldType != change.NewType
		})
	}

	var b strings.Builder
	fmt.Fprintf(&b, "You are helping me migrate tables in my %s database before I upgrade a CloudQuery source plugin.\n", database.name)

	b.WriteString("\n## Goal\n\n")
	fmt.Fprintf(&b, "Change the tables listed below so that version %s of the %s source plugin can sync into them with `migrate_mode: safe`, and keep their existing rows. ", r.ToVersion, r.SourceName)
	b.WriteString("Work in small steps that I can reverse. Before you run any statement that writes data or changes the schema, show it to me and wait for my approval.\n")

	b.WriteString("\n## Environment\n\n")
	fmt.Fprintf(&b, "- Source plugin: %s, upgrade from %s to %s\n", r.SourceName, r.FromVersion, r.ToVersion)
	fmt.Fprintf(&b, "- Destination plugin: %s\n", d.VersionString())
	fmt.Fprintf(&b, "- Database: %s\n", database.name)
	fmt.Fprintf(&b, "- write_mode: %s, pk_mode: %s, migrate_mode: %s\n", d.WriteMode, d.PKMode, d.MigrateMode)
	if d.MigrateMode == specs.MigrateModeForced {
		b.WriteString("- With `migrate_mode: forced`, the first sync of the new version drops and recreates these tables, and deletes their rows, unless they already match the new schema.\n")
	}

	b.WriteString("\n## Tables to migrate\n")
	for _, impact := range manual {
		writeUpgradePromptTable(&b, impact, r.Tables[impact.Name])
	}

	b.WriteString("\n## CloudQuery rules you must respect\n\n")
	b.WriteString("- Change only the tables listed above. CloudQuery migrates any other tables of this upgrade on the next sync.\n")
	b.WriteString("- CloudQuery adds the internal columns `_cq_id`, `_cq_parent_id` (on child tables), `_cq_source_name` and `_cq_sync_time`. Keep these columns and their values.\n")
	b.WriteString("- " + upgradeWriteModeFact(d.WriteMode) + "\n")
	fmt.Fprintf(&b, "- The final schema of each table must match exactly what version %s creates: column names, types, nullability and primary key. If it does not match, the next sync with `migrate_mode: safe` fails again. If you are not sure about a target definition, ask me to run `cloudquery migrate` with version %s against an empty scratch database, and compare with the tables it creates.\n", r.ToVersion, r.ToVersion)
	b.WriteString("- Pause all syncs that write to these tables until the migration is complete.\n")
	fmt.Fprintf(&b, "- Change the source plugin version to %s in the CloudQuery configuration only after the migration is complete and verified.\n", r.ToVersion)

	b.WriteString("\n## Step 1: Inspect the database (read-only)\n\n")
	b.WriteString("Before you propose changes, run only read-only queries (no statement that writes data or changes the schema) and tell me, for each table:\n\n")
	for _, item := range []string{
		"The row count, and the size of the table and its indexes.",
		"The columns, indexes and constraints, compared with the changes above.",
		"Views, functions or triggers that depend on the table.",
		"Foreign keys to or from the table.",
		"Replication, change data capture or other consumers that read the table.",
		"Current locks, long-running transactions, and when the table has the most and the least traffic.",
		"The free disk space.",
	} {
		b.WriteString("- " + item + "\n")
	}

	b.WriteString("\n## Step 2: Propose a plan\n\n")
	b.WriteString("Propose a plan before you run anything. Prefer gradual steps that keep the data. The plan must include:\n\n")
	if hasTypeChange {
		b.WriteString("- Type conversions: an explicit conversion expression for each changed column, what it does with NULL and empty values, and a check that the result matches what the new plugin version writes.\n")
	}
	if hasPKChange {
		b.WriteString("- Primary-key changes: how to fill new key columns for existing rows, a unique index built without blocking writes, and a short swap of the constraint. Existing rows can have no value for a new key column. Do not invent values. Explain the options and how my write_mode affects them, for example a placeholder value that the next sync replaces.\n")
	}
	b.WriteString("- Performance: table rewrites, lock levels and how long each lock is held, batch sizes, when to run each step (off-peak), index builds that do not block writes, free disk space for rewrites, and replication lag.\n")
	b.WriteString("- Verification: queries to run after each step, for example row counts before and after, a sample of converted values compared with the original values, and a schema diff of each table against the target schema.\n")
	b.WriteString("- Rollback: how to undo each step.\n")
	b.WriteString("- Trade-offs: if a gradual migration is not possible or not worth it, say why. For example, `migrate_mode: forced` and a full sync can be simpler for small tables, or when the next sync replaces all rows anyway, but the tables are empty until that sync completes.\n")

	b.WriteString("\n## Step 3: Run the plan\n\n")
	b.WriteString("- Run one step at a time. Before each statement that writes data or changes the schema, show me the exact statement and wait for my approval.\n")
	b.WriteString("- Never run `DROP`, `TRUNCATE`, `DELETE` or any other statement that removes data or schema objects without my explicit approval of that exact statement.\n")
	b.WriteString("- Back up the affected tables before the first write, and make sure that I can restore the backup.\n")
	b.WriteString("- After each step, run the verification queries. Stop and tell me if a check fails.\n")
	fmt.Fprintf(&b, "- When all tables match the target schema, tell me to change the source plugin version to %s, resume the syncs, and check that the first sync completes without migration errors.\n", r.ToVersion)
	b.WriteString("- Never print, log or ask me to paste passwords, tokens or connection strings.\n")

	fmt.Fprintf(&b, "\n## %s guidance\n\n", database.name)
	guidance := slices.Clone(database.general)
	if hasTypeChange {
		guidance = append(guidance, database.typeChange...)
	}
	if hasPKChange {
		guidance = append(guidance, database.primaryKey...)
	}
	for _, item := range guidance {
		b.WriteString("- " + item + "\n")
	}
	return b.String()
}

func writeUpgradePromptTable(b *strings.Builder, impact upgradeTableImpact, table upgradeTablePair) {
	fmt.Fprintf(b, "\n### %s\n\n", impact.Name)
	oldPrimaryKey, newPrimaryKey := upgradePrimaryKeys(table.From), upgradePrimaryKeys(table.To)
	if slices.Equal(oldPrimaryKey, newPrimaryKey) {
		fmt.Fprintf(b, "- Primary key: %s (unchanged)\n", upgradePromptColumnList(oldPrimaryKey))
	} else {
		fmt.Fprintf(b, "- Primary key now: %s\n", upgradePromptColumnList(oldPrimaryKey))
		fmt.Fprintf(b, "- Primary key after the upgrade: %s\n", upgradePromptColumnList(newPrimaryKey))
	}
	b.WriteString("- Changes:\n")
	for _, change := range impact.Changes {
		fmt.Fprintf(b, "  - %s\n", upgradePromptChange(change, upgradeTableColumn(table.To, change.Column)))
	}
	fmt.Fprintf(b, "- Why `migrate_mode: safe` rejects it: %s\n", strings.TrimPrefix(impact.Outcomes[specs.MigrateModeSafe.String()].Text, upgradeFailsPrefix))
}

func upgradePromptChange(change upgradeColumnChange, newColumn *schema.Column) string {
	var description string
	switch change.Kind {
	case upgradeColumnAdded:
		description = fmt.Sprintf("`%s`: new `%s` column", change.Column, change.NewType)
	case upgradeColumnRemoved:
		description = fmt.Sprintf("`%s`: `%s` column removed", change.Column, change.OldType)
	default:
		description = fmt.Sprintf("`%s`: `%s` → `%s`", change.Column, change.OldType, change.NewType)
	}
	if newColumn != nil && (newColumn.NotNull || newColumn.PrimaryKey) {
		description += ", NOT NULL"
	}
	return upgradeJoinNonEmpty(description, change.Reason)
}

func upgradePrimaryKeys(table *schema.Table) []string {
	if table == nil {
		return nil
	}
	return table.PrimaryKeys()
}

func upgradePromptColumnList(columns []string) string {
	if len(columns) == 0 {
		return "none"
	}
	quoted := make([]string, len(columns))
	for i, column := range columns {
		quoted[i] = "`" + column + "`"
	}
	return strings.Join(quoted, ", ")
}

func upgradeWriteModeFact(writeMode specs.WriteMode) string {
	switch writeMode {
	case specs.WriteModeOverwrite:
		return "With `write_mode: overwrite`, each sync upserts rows by primary key and does not delete rows that the source no longer returns."
	case specs.WriteModeAppend:
		return "With `write_mode: append`, each sync only inserts rows, so the existing rows are history that exists only in this database."
	}
	return "With `write_mode: overwrite-delete-stale`, each sync upserts rows by primary key, then deletes the rows of the same `_cq_source_name` that have an older `_cq_sync_time`. The next complete sync replaces or deletes the rows that you keep now."
}

func renderUpgradeAIPrompt(w io.Writer, r upgradeReport) error {
	prompt := upgradeAIPrompt(r)
	if prompt == "" {
		_, err := fmt.Fprintf(w, "%s\n\n", upgradeNoManualText)
		return err
	}
	_, err := fmt.Fprintf(w, "--- AI migration prompt for %s (prompt_version %s) ---\n%s%s\n\n", r.Destination.Name, upgradeAIPromptVersion, prompt, upgradeAIPromptEndText)
	return err
}
