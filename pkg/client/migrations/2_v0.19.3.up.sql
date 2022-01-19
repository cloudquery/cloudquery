ALTER TABLE IF EXISTS fetches
    ADD COLUMN provider_alias TEXT;
ALTER TABLE IF EXISTS fetches
    DROP CONSTRAINT fetches_pk;
ALTER TABLE IF EXISTS fetches
    ADD CONSTRAINT "fetches_pk" UNIQUE (fetch_id, provider_name, provider_alias);

