-- Resource: redshift.event_subscriptions
CREATE TABLE IF NOT EXISTS "aws_redshift_event_subscriptions" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"id" text,
	"customer_aws_id" text,
	"enabled" boolean,
	"event_categories_list" text[],
	"severity" text,
	"sns_topic_arn" text,
	"source_ids_list" text[],
	"source_type" text,
	"status" text,
	"subscription_creation_time" timestamp without time zone,
	"tags" jsonb,
	CONSTRAINT aws_redshift_event_subscriptions_pk PRIMARY KEY(cq_fetch_date,account_id,id),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_redshift_event_subscriptions');

-- Resource: redshift.clusters
CREATE TABLE IF NOT EXISTS "aws_redshift_snapshots" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"actual_incremental_backup_size" float,
	"availability_zone" text,
	"backup_progress" float,
	"cluster_create_time" timestamp without time zone,
	"cluster_identifier" text,
	"cluster_version" text,
	"current_backup_rate" float,
	"db_name" text,
	"elapsed_time" bigint,
	"encrypted" boolean,
	"encrypted_with_hsm" boolean,
	"engine_full_version" text,
	"enhanced_vpc_routing" boolean,
	"estimated_seconds_to_completion" bigint,
	"kms_key_id" text,
	"maintenance_track_name" text,
	"manual_snapshot_remaining_days" integer,
	"manual_snapshot_retention_period" integer,
	"master_username" text,
	"node_type" text,
	"number_of_nodes" integer,
	"owner_account" text,
	"port" integer,
	"restorable_node_types" text[],
	"snapshot_create_time" timestamp without time zone,
	"snapshot_identifier" text,
	"snapshot_retention_start_time" timestamp without time zone,
	"snapshot_type" text,
	"source_region" text,
	"status" text,
	"total_backup_size_in_mega_bytes" float,
	"vpc_id" text,
	"tags" jsonb,
	CONSTRAINT aws_redshift_snapshots_pk PRIMARY KEY(cq_fetch_date,cluster_identifier,cluster_create_time),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_redshift_snapshots');
CREATE TABLE IF NOT EXISTS "aws_redshift_snapshot_accounts_with_restore_access" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"snapshot_cq_id" uuid,
	"account_alias" text,
	"account_id" text,
	CONSTRAINT aws_redshift_snapshot_accounts_with_restore_access_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_redshift_snapshot_accounts_with_restore_access (cq_fetch_date, snapshot_cq_id);
SELECT setup_tsdb_child('aws_redshift_snapshot_accounts_with_restore_access', 'snapshot_cq_id', 'aws_redshift_snapshots', 'cq_id');
