ALTER TABLE IF EXISTS "aws_sns_topics" ADD COLUMN IF NOT EXISTS "tags" jsonb;

CREATE TABLE IF NOT EXISTS "aws_elasticbeanstalk_application_versions" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
	"application_name" text,
	"arn" text,
	"build_arn" text,
	"date_created" timestamp without time zone,
	"date_updated" timestamp without time zone,
	"description" text,
	"source_location" text,
	"source_repository" text,
	"source_type" text,
	"source_bundle_s3_bucket" text,
	"source_bundle_s3_key" text,
	"status" text,
	"version_label" text,
	CONSTRAINT aws_elasticbeanstalk_application_versions_pk PRIMARY KEY(arn),
	UNIQUE(cq_id)
);