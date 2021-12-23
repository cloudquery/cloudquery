CREATE TABLE IF NOT EXISTS cq_fetches
(
    cq_id                UUID NOT NULL,
    cq_fetch_id          UUID NOT NULL,
    START                TIMESTAMP,
    finish               TIMESTAMP,
    total_resource_count BIGINT,
    total_errors_count   BIGINT,
    provider_name        TEXT,
    provider_version     TEXT,
    is_success           BOOLEAN,
    provider_meta        jsonb,
    results              jsonb,
    CONSTRAINT "cq_fetches_id" PRIMARY KEY (cq_id),
    CONSTRAINT "cq_fetches_pk" UNIQUE (cq_fetch_id, provider_name),
    CONSTRAINT "non_nil_fetch_id" CHECK (cq_fetch_id != '00000000-0000-0000-0000-000000000000')
);