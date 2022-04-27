-- Resource: backup.global_settings
CREATE TABLE IF NOT EXISTS "aws_backup_global_settings" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"global_settings" jsonb,
	"last_update_time" timestamp without time zone,
	CONSTRAINT aws_backup_global_settings_pk PRIMARY KEY(account_id),
	UNIQUE(cq_id)
);

-- Resource: backup.region_settings
CREATE TABLE IF NOT EXISTS "aws_backup_region_settings" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
	"resource_type_management_preference" jsonb,
	"resource_type_opt_in_preference" jsonb,
	CONSTRAINT aws_backup_region_settings_pk PRIMARY KEY(account_id,region),
	UNIQUE(cq_id)
);
