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


-- Resource: workspaces.directories
CREATE TABLE IF NOT EXISTS "aws_workspaces_directories" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
	"alias" text,
	"customer_user_name" text,
	"id" text,
	"name" text,
	"type" text,
    "arn" text,
	"dns_ip_addresses" text[],
	"iam_role_id" text,
	"ip_group_ids" text[],
	"registration_code" text,
	"change_compute_type" text,
	"increase_volume_size" text,
	"rebuild_workspace" text,
	"restart_workspace" text,
	"switch_running_mode" text,
	"state" text,
	"subnet_ids" text[],
	"tenancy" text,
	"device_type_android" text,
	"device_type_chrome_os" text,
	"device_type_ios" text,
	"device_type_linux" text,
	"device_type_osx" text,
	"device_type_web" text,
	"device_type_windows" text,
	"device_type_zero_client" text,
	"custom_security_group_id" text,
	"default_ou" text,
	"enable_internet_access" boolean,
	"enable_maintenance_mode" boolean,
	"enable_work_docs" boolean,
	"user_enabled_as_local_administrator" boolean,
	"workspace_security_group_id" text,
	CONSTRAINT aws_workspaces_directories_pk PRIMARY KEY(id),
	UNIQUE(cq_id)
);

-- Resource: workspaces.workspaces
CREATE TABLE IF NOT EXISTS "aws_workspaces_workspaces" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
    "arn" text,
	"bundle_id" text,
	"computer_name" text,
	"directory_id" text,
	"error_code" text,
	"error_message" text,
	"ip_address" text,
	"modification_states" jsonb,
	"root_volume_encryption_enabled" boolean,
	"state" text,
	"subnet_id" text,
	"user_name" text,
	"user_volume_encryption_enabled" boolean,
	"volume_encryption_key" text,
	"id" text,
	"compute_type_name" text,
	"root_volume_size_gib" integer,
	"running_mode" text,
	"running_mode_auto_stop_timeout_in_minutes" integer,
	"user_volume_size_gib" integer,
	CONSTRAINT aws_workspaces_workspaces_pk PRIMARY KEY(id),
	UNIQUE(cq_id)
);

-- Resource: redshift.event_subscriptions
CREATE TABLE IF NOT EXISTS "aws_redshift_event_subscriptions" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
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
	CONSTRAINT aws_redshift_event_subscriptions_pk PRIMARY KEY(account_id,id),
	UNIQUE(cq_id)
);

-- Resource: redshift.clusters
CREATE TABLE IF NOT EXISTS "aws_redshift_snapshots" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
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
	CONSTRAINT aws_redshift_snapshots_pk PRIMARY KEY(cluster_identifier,cluster_create_time),
	UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "aws_redshift_snapshot_accounts_with_restore_access" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"snapshot_cq_id" uuid,
	"account_alias" text,
	"account_id" text,
	CONSTRAINT aws_redshift_snapshot_accounts_with_restore_access_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (snapshot_cq_id) REFERENCES aws_redshift_snapshots(cq_id) ON DELETE CASCADE
);
