-- Resource: wafv2.ipsets
CREATE TABLE IF NOT EXISTS "aws_wafv2_ipsets" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
	"scope" text,
	"arn" text,
	"addresses" cidr[],
	"ip_address_version" text,
	"id" text,
	"name" text,
	"description" text,
	"tags" jsonb,
	CONSTRAINT aws_wafv2_ipsets_pk PRIMARY KEY(arn),
	UNIQUE(cq_id)
);

-- Resource: wafv2.regex_pattern_sets
CREATE TABLE IF NOT EXISTS "aws_wafv2_regex_pattern_sets" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
	"scope" text,
	"arn" text,
	"description" text,
	"id" text,
	"name" text,
	"regular_expression_list" text[],
	"tags" jsonb,
	CONSTRAINT aws_wafv2_regex_pattern_sets_pk PRIMARY KEY(arn),
	UNIQUE(cq_id)
);

-- Resource: iam.virtual_mfa_devices
ALTER TABLE IF EXISTS aws_iam_virtual_mfa_devices DROP CONSTRAINT aws_iam_virtual_mfa_devices_pk;
ALTER TABLE IF EXISTS aws_iam_virtual_mfa_devices ADD CONSTRAINT aws_iam_virtual_mfa_devices_pk PRIMARY KEY (serial_number);

-- Resource: ec2.egress_only_internet_gateways
CREATE TABLE IF NOT EXISTS "aws_ec2_egress_only_internet_gateways" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "account_id" text,
    "region" text,
    "arn" text,
    "attachments" jsonb,
    "id" text,
    "tags" jsonb,
    CONSTRAINT aws_ec2_egress_only_internet_gateways_pk PRIMARY KEY(arn),
    UNIQUE(cq_id)
    );



-- Resource: qldb.ledgers
CREATE TABLE IF NOT EXISTS "aws_qldb_ledgers" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "account_id" text,
    "region" text,
    "tags" jsonb,
    "arn" text,
    "creation_date_time" timestamp without time zone,
    "deletion_protection" boolean,
    "encryption_status" text,
    "kms_key_arn" text,
    "inaccessible_kms_key_date_time" timestamp without time zone,
    "name" text,
    "permissions_mode" text,
    "state" text,
    CONSTRAINT aws_qldb_ledgers_pk PRIMARY KEY(arn),
    UNIQUE(cq_id)
    );
CREATE TABLE IF NOT EXISTS "aws_qldb_ledger_journal_kinesis_streams" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "ledger_cq_id" uuid,
    "stream_arn" text,
    "aggregation_enabled" boolean,
    "ledger_name" text,
    "role_arn" text,
    "status" text,
    "stream_id" text,
    "stream_name" text,
    "arn" text,
    "creation_time" timestamp without time zone,
    "error_cause" text,
    "exclusive_end_time" timestamp without time zone,
    "inclusive_start_time" timestamp without time zone,
    CONSTRAINT aws_qldb_ledger_journal_kinesis_streams_pk PRIMARY KEY(cq_id),
    UNIQUE(cq_id),
    FOREIGN KEY (ledger_cq_id) REFERENCES aws_qldb_ledgers(cq_id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS "aws_qldb_ledger_journal_s3_exports" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "ledger_cq_id" uuid,
    "exclusive_end_time" timestamp without time zone,
    "export_creation_time" timestamp without time zone,
    "export_id" text,
    "inclusive_start_time" timestamp without time zone,
    "ledger_name" text,
    "role_arn" text,
    "bucket" text,
    "object_encryption_type" text,
    "kms_key_arn" text,
    "prefix" text,
    "status" text,
    "output_format" text,
    CONSTRAINT aws_qldb_ledger_journal_s3_exports_pk PRIMARY KEY(cq_id),
    UNIQUE(cq_id),
    FOREIGN KEY (ledger_cq_id) REFERENCES aws_qldb_ledgers(cq_id) ON DELETE CASCADE
    );