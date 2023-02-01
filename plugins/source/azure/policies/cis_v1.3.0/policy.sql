SET TIME ZONE 'UTC';
-- trick to set execution_time if not already set
-- https://stackoverflow.com/questions/32582600/only-set-variable-in-psql-script-if-not-specified-on-the-command-line
\set execution_time :execution_time
SELECT CASE 
  WHEN :'execution_time' = ':execution_time' THEN to_char(now(), 'YYYY-MM-dd HH24:MI:SS.US')
  ELSE :'execution_time'
END AS "execution_time"  \gset
\set framework 'cis_v1.3.0'
\echo "Creating azure_policy_results table if not exist"
\ir ../create_azure_policy_results.sql
\echo "Creating view view_azure_security_policy_parameters"
\ir ../views/policy_assignment_parameters.sql
\echo "Executing CIS V1.3.0 Section 1 (Manual)"
\echo "Executing CIS V1.3.0 Section 2"
\set check_id "2.1"
\echo "Executing check 2.1"
\ir ../queries/security/defender_on_for_servers.sql
\set check_id "2.2"
\echo "Executing check 2.2"
\ir ../queries/security/defender_on_for_app_service.sql
\set check_id "2.3"
\echo "Executing check 2.3"
\ir ../queries/security/defender_on_for_sql_servers.sql
\set check_id "2.4"
\echo "Executing check 2.4"
\ir ../queries/security/defender_on_for_sql_servers_on_machines.sql
\set check_id "2.5"
\echo "Executing check 2.5"
\ir ../queries/security/defender_on_for_storage.sql
\set check_id "2.6"
\echo "Executing check 2.6"
\ir ../queries/security/defender_on_for_k8s.sql
\set check_id "2.7"
\echo "Executing check 2.7"
\ir ../queries/security/defender_on_for_container_registeries.sql
\set check_id "2.8"
\echo "Executing check 2.8"
\ir ../queries/security/defender_on_for_key_vault.sql
-- security settings does not have "enabled" property
-- \set check_id "2.10"
-- \echo "Executing check 2.10"
-- \ir ../queries/security/mcas_integration_with_security_center_enabled.sql
\set check_id "2.11"
\echo "Executing check 2.11"
\ir ../queries/security/auto_provisioning_monitoring_agent_enabled.sql
\set check_id "2.12"
\echo "Executing check 2.12"
\ir ../queries/security/default_policy_disabled.sql
-- security contacts api is broken
-- \set check_id "2.13"
-- \echo "Executing check 2.13"
-- \ir ../queries/security/security_email_configured.sql
-- \set check_id "2.14"
-- \echo "Executing check 2.14"
-- \ir ../queries/security/notify_high_severity_alerts.sql

\echo "Executing CIS V1.3.0 Section 4"
\set check_id "4.1.1"
\echo "Executing check 4.1.1"
\ir ../queries/sql/auditing_off.sql
\set check_id "4.1.2"
\echo "Executing check 4.1.2"
\ir ../queries/sql/data_encryption_off.sql
\set check_id "4.1.3"
\echo "Executing check 4.1.3"
\ir ../queries/sql/auditing_retention_less_than_90_days.sql
\set check_id "4.2.1"
\echo "Executing check 4.2.1"
\ir ../queries/sql/atp_on_sql_server_disabled.sql
\set check_id "4.2.2"
\echo "Executing check 4.2.2"
\ir ../queries/sql/va_is_enabled_on_sql_server_by_storage_account.sql
\set check_id "4.2.3"
\echo "Executing check 4.2.3"
\ir ../queries/sql/va_periodic_scans_enabled_on_sql_server.sql
\set check_id "4.2.4"
\echo "Executing check 4.2.4"
\ir ../queries/sql/va_send_scan_report_enabled_on_sql_server.sql
\set check_id "4.2.5"
\echo "Executing check 4.2.5"
\ir ../queries/sql/va_send_email_to_admins_and_owners_enabled.sql
\set check_id "4.3.1"
\echo "Executing check 4.3.1"
\ir ../queries/sql/postgresql_ssl_enforcment_disabled.sql
\set check_id "4.3.2"
\echo "Executing check 4.3.2"
\ir ../queries/sql/mysql_ssl_enforcment_disabled.sql
\set check_id "4.3.3"
\echo "Executing check 4.3.3"
\ir ../queries/sql/postgresql_log_checkpoints_disabled.sql
\set check_id "4.3.4"
\echo "Executing check 4.3.4"
\ir ../queries/sql/postgresql_log_connections_disabled.sql
\set check_id "4.3.5"
\echo "Executing check 4.3.5"
\ir ../queries/sql/postgresql_log_disconnections_disabled.sql
\set check_id "4.3.6"
\echo "Executing check 4.3.6"
\ir ../queries/sql/postgresql_connection_throttling_disabled.sql
\set check_id "4.3.7"
\echo "Executing check 4.3.7"
\ir ../queries/sql/postgresql_log_retention_days_less_than_3_days.sql
\set check_id "4.3.8"
\echo "Executing check 4.3.8"
\ir ../queries/sql/postgresql_allow_access_to_azure_services_enabled.sql
\set check_id "4.4"
\echo "Executing check 4.4"
\ir ../queries/sql/ad_admin_configured.sql
\set check_id "4.5"
\echo "Executing check 4.5"
\ir ../queries/sql/sqlserver_tde_not_encrypted_with_cmek.sql


\ir section_7.sql
\ir section_8.sql
\ir section_9.sql
