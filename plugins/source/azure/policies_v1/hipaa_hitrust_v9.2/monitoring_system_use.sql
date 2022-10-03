\echo "Monitoring System Use"
\set check_id '1120.09ab3System.9 - 09.ab - 1'
\echo :check_id
\ir ../queries/monitor/azure_monitor_should_collect_activity_logs_from_all_regions.sql
\set check_id '12100.09ab2System.15 - 09.ab - 1'
\echo :check_id
\ir ../queries/compute/machines_without_log_analytics_agent.sql
\set check_id '12101.09ab1Organizational.3 - 09.ab - 1'
\echo :check_id
\ir ../queries/compute/scale_sets_without_log_analytics_agent.sql
\set check_id '12102.09ab1Organizational.4 - 09.ab - 1'
\echo :check_id
\ir ../queries/compute/guestconfiguration_windowsloganalyticsagentconnection_aine.sql
\set check_id '1212.09ab1System.1 - 09.ab - 1'
\echo :check_id
\ir ../queries/monitor/azure_monitor_log_profile_should_collect_logs_for_categories_write_delete_and_action.sql
\set check_id '1213.09ab2System.128 - 09.ab - 1'
\echo :check_id
\ir ../queries/security/asc_automatic_provisioning_log_analytics_monitoring_agent.sql
\set check_id '1214.09ab2System.3456 - 09.ab - 1'
\echo :check_id
\ir ../queries/monitor/azure_monitor_should_collect_activity_logs_from_all_regions.sql
\set check_id '1215.09ab2System.7 - 09.ab - 1'
\echo :check_id
\ir ../queries/compute/machines_without_log_analytics_agent.sql
\set check_id '1216.09ab3System.12 - 09.ab - 1'
\echo :check_id
\ir ../queries/compute/scale_sets_without_log_analytics_agent.sql
\set check_id '1217.09ab3System.3 - 09.ab - 1'
\echo :check_id
\ir ../queries/monitor/azure_monitor_log_profile_should_collect_logs_for_categories_write_delete_and_action.sql
\set check_id '1219.09ab3System.10 - 09.ab - 1'
\echo :check_id
\ir ../queries/monitor/azure_monitor_log_profile_should_collect_logs_for_categories_write_delete_and_action.sql
\set check_id '1220.09ab3System.56 - 09.ab - 1'
\echo :check_id
\ir ../queries/security/asc_automatic_provisioning_log_analytics_monitoring_agent.sql
