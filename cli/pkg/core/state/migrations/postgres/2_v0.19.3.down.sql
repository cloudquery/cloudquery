ALTER TABLE IF EXISTS fetches
    DROP CONSTRAINT fetches_pk;
ALTER TABLE IF EXISTS fetches ADD CONSTRAINT  "fetches_pk" UNIQUE (fetch_id, provider_name);
ALTER TABLE IF EXISTS fetches
    DROP COLUMN provider_alias;