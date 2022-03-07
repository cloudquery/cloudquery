-- Resource: sql.servers
ALTER TABLE IF EXISTS "azure_sql_databases"
    ADD COLUMN IF NOT EXISTS "backup_long_term_retention_policy" jsonb;


-- Resource: account.locations
CREATE TABLE IF NOT EXISTS "azure_account_locations"
(
    "cq_id"           uuid NOT NULL,
    "cq_meta"         jsonb,
    "subscription_id" text,
    "id"              text,
    "name"            text,
    "display_name"    text,
    "latitude"        text,
    "longitude"       text,
    CONSTRAINT azure_account_locations_pk PRIMARY KEY (subscription_id, id),
    UNIQUE (cq_id)
);
