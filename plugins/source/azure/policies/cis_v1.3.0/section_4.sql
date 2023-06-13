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
