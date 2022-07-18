\set framework 'cis_v1.2.0'
\echo "Creating CIS V1.2.0 Section 2 Views"
\i views/log_metric_filters.sql
\echo "Executing CIS V1.2.0 Section 2"
\set check_id "2.1"
\echo "Executiong check 2.1"
\i queries/logging/not_configured_across_services_and_users.sql
\set check_id "2.2"
\echo "Executiong check 2.2"
\i queries/logging/sinks_not_configured_for_all_log_entries.sql
\set check_id "2.3"
\echo "Executiong check 2.3"
\i queries/logging/log_buckets_retention_policy_disabled.sql
\set check_id "2.4"
\echo "Executiong check 2.4"
\i queries/logging/project_ownership_changes_without_log_metric_filter_alerts.sql
\set check_id "2.5"
\echo "Executiong check 2.5"
\i queries/logging/audit_config_changes_without_log_metric_filter_alerts.sql
\set check_id "2.6"
\echo "Executiong check 2.6"
\i queries/logging/custom_role_changes_without_log_metric_filter_alerts.sql
\set check_id "2.7"
\echo "Executiong check 2.7"
\i queries/logging/vpc_firewall_changes_without_log_metric_filter_alerts.sql
\set check_id "2.8"
\echo "Executiong check 2.8"
\i queries/logging/vpc_route_changes_without_log_metric_filter_alerts.sql
\set check_id "2.9"
\echo "Executiong check 2.9"
\i queries/logging/vpc_network_changes_without_log_metric_filter_alerts.sql
\set check_id "2.10"
\echo "Executiong check 2.10"
\i queries/logging/storage_iam_changes_without_log_metric_filter_alerts.sql
\set check_id "2.11"
\echo "Executiong check 2.11"
\i queries/logging/sql_instance_changes_without_log_metric_filter_alerts.sql
\set check_id "2.12"
\echo "Executiong check 2.12"
\i queries/logging/dns_logging_disabled.sql
