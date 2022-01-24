ALTER TABLE IF EXISTS fetches
    ADD COLUMN IF NOT EXISTS provider_alias TEXT;
ALTER TABLE IF EXISTS fetches
    ADD COLUMN IF NOT EXISTS core_version TEXT;
ALTER TABLE IF EXISTS fetches
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP;
ALTER TABLE IF EXISTS fetches
    DROP CONSTRAINT fetches_pk;
ALTER TABLE IF EXISTS fetches
    ADD CONSTRAINT "fetches_pk" UNIQUE (fetch_id, provider_name, provider_alias);