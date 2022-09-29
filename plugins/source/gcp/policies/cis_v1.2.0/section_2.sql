\set framework 'cis_v1.2.0'
\echo "Creating CIS V1.2.0 Section 2 Views"
\ir ../views/log_metric_filters.sql
\echo "Executing CIS V1.2.0 Section 2"
\set check_id '2.1'
\echo "Executing check 2.1"
\ir ../queries/logging/not_configured_across_services_and_users.sql
\set check_id '2.2'
\echo "Executing check 2.2"
\ir ../queries/logging/sinks_not_configured_for_all_log_entries.sql
\set check_id '2.3'
\echo "Executing check 2.3"
\ir ../queries/logging/log_buckets_retention_policy_disabled.sql
\set check_id '2.4'
\echo "Executing check 2.4"
\ir ../queries/logging/project_ownership_changes_without_log_metric_filter_alerts.sql
\set check_id '2.5'
\echo "Executing check 2.5"
\ir ../queries/logging/audit_config_changes_without_log_metric_filter_alerts.sql
\set check_id '2.6'
\echo "Executing check 2.6"
\ir ../queries/logging/custom_role_changes_without_log_metric_filter_alerts.sql
\set check_id '2.7'
\echo "Executing check 2.7"
\ir ../queries/logging/vpc_firewall_changes_without_log_metric_filter_alerts.sql
\set check_id '2.8'
\echo "Executing check 2.8"
\ir ../queries/logging/vpc_route_changes_without_log_metric_filter_alerts.sql
\set check_id '2.9'
\echo "Executing check 2.9"
\ir ../queries/logging/vpc_network_changes_without_log_metric_filter_alerts.sql
\set check_id '2.10'
\echo "Executing check 2.10"
\ir ../queries/logging/storage_iam_changes_without_log_metric_filter_alerts.sql
\set check_id '2.11'
\echo "Executing check 2.11"
\ir ../queries/logging/sql_instance_changes_without_log_metric_filter_alerts.sql
\set check_id '2.12'
\echo "Executing check 2.12"
\ir ../queries/logging/dns_logging_disabled.sql
