CREATE TABLE IF NOT EXISTS "aws_access_analyzer_analyzer_archive_rules"
(
    "cq_id"          uuid NOT NULL,
    "cq_meta"        jsonb,
    "analyzer_cq_id" uuid,
    "created_at"     timestamp without time zone,
    "filter"         jsonb,
    "rule_name"      text,
    "updated_at"     timestamp without time zone,
    CONSTRAINT aws_access_analyzer_analyzer_archive_rules_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (analyzer_cq_id) REFERENCES aws_access_analyzer_analyzers (cq_id) ON DELETE CASCADE
);

-- Resource: ec2.images
ALTER TABLE IF EXISTS "aws_ec2_images" ADD COLUMN IF NOT EXISTS "last_launched_time" timestamp without time zone;
