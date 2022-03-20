CREATE TABLE IF NOT EXISTS "aws_access_analyzer_analyzer_archive_rules"
(
    "cq_id"          uuid                        NOT NULL,
    "cq_meta"        jsonb,
    "cq_fetch_date"  timestamp without time zone NOT NULL,
    "analyzer_cq_id" uuid,
    "created_at"     timestamp without time zone,
    "filter"         jsonb,
    "rule_name"      text,
    "updated_at"     timestamp without time zone,
    CONSTRAINT aws_access_analyzer_analyzer_archive_rules_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_access_analyzer_analyzer_archive_rules (cq_fetch_date, analyzer_cq_id);
SELECT setup_tsdb_child('aws_access_analyzer_analyzer_archive_rules', 'analyzer_cq_id', 'aws_access_analyzer_analyzers',
                        'cq_id');
