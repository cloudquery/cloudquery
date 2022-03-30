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

-- Resource: ec2.security_groups
ALTER TABLE IF EXISTS aws_ec2_security_group_ip_permission_user_id_group_pairs DROP CONSTRAINT aws_ec2_security_group_ip_permission_user_id_group_pairs_pk;
ALTER TABLE IF EXISTS aws_ec2_security_group_ip_permission_user_id_group_pairs ADD CONSTRAINT aws_ec2_security_group_ip_permission_user_id_group_pairs_pk PRIMARY KEY (cq_id);
