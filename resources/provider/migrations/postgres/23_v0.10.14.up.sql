-- Resource: ec2.nat_gateways
ALTER TABLE IF EXISTS aws_ec2_nat_gateway_addresses DROP CONSTRAINT aws_ec2_nat_gateway_addresses_pk;
ALTER TABLE IF EXISTS aws_ec2_nat_gateway_addresses ADD CONSTRAINT aws_ec2_nat_gateway_addresses_pk PRIMARY KEY (nat_gateway_cq_id,network_interface_id);
ALTER TABLE aws_ec2_nat_gateway_addresses ALTER COLUMN allocation_id DROP NOT NULL;

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