-- Resource: sql.servers
ALTER TABLE IF EXISTS "azure_sql_databases"
    ADD COLUMN IF NOT EXISTS "backup_long_term_retention_policy" jsonb;


-- Resource: account.locations
CREATE TABLE IF NOT EXISTS "azure_account_locations"
(
    "cq_id"           uuid                        NOT NULL,
    "cq_meta"         jsonb,
    "cq_fetch_date"   timestamp without time zone NOT NULL,
    "subscription_id" text,
    "id"              text,
    "name"            text,
    "display_name"    text,
    "latitude"        text,
    "longitude"       text,
    CONSTRAINT azure_account_locations_pk PRIMARY KEY (cq_fetch_date, subscription_id, id),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('azure_account_locations');
