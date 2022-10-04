\set check_id 'RDS.1'
\echo "Executing check RDS.1"
\ir ../queries/rds/snapshots_should_prohibit_public_access.sql

\set check_id 'RDS.2'
\echo "Executing check RDS.2"
\ir ../queries/rds/rds_db_instances_should_prohibit_public_access.sql

\set check_id 'RDS.3'
\echo "Executing check RDS.3"
\ir ../queries/rds/rds_db_instances_should_have_encryption_at_rest_enabled.sql

\set check_id 'RDS.4'
\echo "Executing check RDS.4"
\ir ../queries/rds/rds_cluster_snapshots_and_database_snapshots_should_be_encrypted_at_rest.sql

\set check_id 'RDS.5'
\echo "Executing check RDS.5"
\ir ../queries/rds/rds_db_instances_should_be_configured_with_multiple_availability_zones.sql

\set check_id 'RDS.6'
\echo "Executing check RDS.6"
\ir ../queries/rds/enhanced_monitoring_should_be_configured_for_rds_db_instances_and_clusters.sql

\set check_id 'RDS.7'
\echo "Executing check RDS.7"
\ir ../queries/rds/rds_clusters_should_have_deletion_protection_enabled.sql

\set check_id 'RDS.8'
\echo "Executing check RDS.8"
\ir ../queries/rds/rds_db_instances_should_have_deletion_protection_enabled.sql

\set check_id 'RDS.9'
\echo "Executing check RDS.9"
\ir ../queries/rds/database_logging_should_be_enabled.sql

\set check_id 'RDS.10'
\echo "Executing check RDS.10"
\ir ../queries/rds/iam_authentication_should_be_configured_for_rds_instances.sql

\set check_id 'RDS.12'
\echo "Executing check RDS.12"
\ir ../queries/rds/iam_authentication_should_be_configured_for_rds_clusters.sql

\set check_id 'RDS.13'
\echo "Executing check RDS.13"
\ir ../queries/rds/rds_automatic_minor_version_upgrades_should_be_enabled.sql

\set check_id 'RDS.14'
\echo "Executing check RDS.14"
\ir ../queries/rds/amazon_aurora_clusters_should_have_backtracking_enabled.sql

\set check_id 'RDS.15'
\echo "Executing check RDS.15"
\ir ../queries/rds/rds_db_clusters_should_be_configured_for_multiple_availability_zones.sql

\set check_id 'RDS.16'
\echo "Executing check RDS.16"
\ir ../queries/rds/rds_db_clusters_should_be_configured_to_copy_tags_to_snapshots.sql

\set check_id 'RDS.17'
\echo "Executing check RDS.17"
\ir ../queries/rds/rds_db_instances_should_be_configured_to_copy_tags_to_snapshots.sql

\set check_id 'RDS.18'
\echo "Executing check RDS.18"
\ir ../queries/rds/rds_instances_should_be_deployed_in_a_vpc.sql

\set check_id 'RDS.19'
\echo "Executing check RDS.19"
\ir ../queries/rds/rds_event_notifications_subscription_should_be_configured_for_critical_cluster_events.sql

\set check_id 'RDS.20'
\echo "Executing check RDS.20"
\ir ../queries/rds/rds_event_notifications_subscription_should_be_configured_for_critical_database_instance_events.sql

\set check_id 'RDS.21'
\echo "Executing check RDS.21"
\ir ../queries/rds/rds_event_notifications_subscription_should_be_configured_for_critical_database_parameter_group_events.sql

\set check_id 'RDS.22'
\echo "Executing check RDS.22"
\ir ../queries/rds/rds_event_notifications_subscription_should_be_configured_for_critical_database_security_group_events.sql

\set check_id 'RDS.23'
\echo "Executing check RDS.23"
\ir ../queries/rds/rds_databases_and_clusters_should_not_use_a_database_engine_default_port.sql
