-- Resource: sql.servers
ALTER TABLE IF EXISTS "azure_sql_databases" DROP COLUMN IF EXISTS "backup_long_term_retention_policy";
