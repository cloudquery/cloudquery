CREATE TABLE IF NOT EXISTS fetches
(
    id                   UUID NOT NULL,
    fetch_id             UUID NOT NULL,
    START                TIMESTAMP,
    finish               TIMESTAMP,
    total_resource_count BIGINT,
    total_errors_count   BIGINT,
    provider_name        TEXT,
    provider_version     TEXT,
    is_success           BOOLEAN,
    results              jsonb,
    CONSTRAINT "fetches_id" PRIMARY KEY (id),
    CONSTRAINT "fetches_pk" UNIQUE (fetch_id, provider_name),
    CONSTRAINT "non_nil_fetch_id" CHECK (fetch_id != '00000000-0000-0000-0000-000000000000')
);