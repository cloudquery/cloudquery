-- Resource: sql.servers
ALTER TABLE IF EXISTS "azure_sql_databases" ADD COLUMN IF NOT EXISTS "backup_long_term_retention_policy" jsonb;
