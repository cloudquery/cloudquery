\set framework 'cis_v1.2.0'
\echo "Executing CIS V1.2.0 Section 6"
-- MANUAL
\set check_id '6.1.1'
\echo "Executing check 6.1.1"
\echo "Ensure that a MySQL database instance does not allow anyone to connect with administrative privileges (Automated)"
\ir ../queries/manual.sql
\set check_id '6.1.2'
\echo "Executing check 6.1.2"
\ir ../queries/sql/mysql_skip_show_database_flag_off.sql
\set check_id '6.1.3'
\echo "Executing check 6.1.3"
\ir ../queries/sql/mysql_local_inline_flag_on.sql
\set check_id '6.2.1'
\echo "Executing check 6.2.1"
\ir ../queries/sql/postgresql_log_checkpoints_flag_off.sql
\set check_id '6.2.2'
\echo "Executing check 6.2.2"
\ir ../queries/sql/postgresql_log_error_verbosity_flag_not_strict.sql
\set check_id '6.2.3'
\echo "Executing check 6.2.3"
\ir ../queries/sql/postgresql_log_connections_flag_off.sql
\set check_id '6.2.4'
\echo "Executing check 6.2.4"
\ir ../queries/sql/postgresql_log_disconnections_flag_off.sql
\set check_id '6.2.5'
\echo "Executing check 6.2.5"
\ir ../queries/sql/postgresql_log_duration_flag_off.sql
\set check_id '6.2.6'
\echo "Executing check 6.2.6"
\ir ../queries/sql/postgresql_log_lock_waits_flag_off.sql
\set check_id '6.2.7'
\echo "Executing check 6.2.7"
\echo "Ensure 'log_statement' database flag for Cloud SQL PostgreSQL instance is set appropriately (Manual)"
\ir ../queries/manual.sql
\set check_id '6.2.8'
\echo "Executing check 6.2.8"
\ir ../queries/sql/postgresql_log_hostname_flag_off.sql
\set check_id '6.2.9'
\echo "Executing check 6.2.9"
\ir ../queries/sql/postgresql_log_parser_stats_flag_on.sql
\set check_id '6.2.10'
\echo "Executing check 6.2.10"
\ir ../queries/sql/postgresql_log_planner_stats_flag_on.sql
\set check_id '6.2.11'
\echo "Executing check 6.2.11"
\ir ../queries/sql/postgresql_log_executor_stats_flag_on.sql
\set check_id '6.2.12'
\echo "Executing check 6.2.12"
\ir ../queries/sql/postgresql_log_statement_stats_flag_on.sql
\set check_id '6.2.13'
\echo "Executing check 6.2.13"
\echo "Ensure that the 'log_min_messages' database flag for Cloud SQL PostgreSQL instance is set appropriately (Manual)"
\ir ../queries/manual.sql
\set check_id '6.2.14'
\echo "Executing check 6.2.14"
\ir ../queries/sql/postgresql_log_min_error_statement_flag_less_error.sql
\set check_id '6.2.15'
\echo "Executing check 6.2.15"
\ir ../queries/sql/postgresql_log_temp_files_flag_off.sql
\set check_id '6.2.16'
\echo "Executing check 6.2.16"
\ir ../queries/sql/postgresql_log_min_duration_statement_flag_on.sql
\set check_id '6.3.1'
\echo "Executing check 6.3.1"
\ir ../queries/sql/sqlserver_external_scripts_enabled_flag_on.sql
\set check_id '6.3.2'
\echo "Executing check 6.3.2"
\ir ../queries/sql/sqlserver_cross_db_ownership_chaining_flag_on.sql
\set check_id '6.3.3'
\echo "Executing check 6.3.3"
\ir ../queries/sql/sqlserver_user_connections_flag_not_set.sql
\set check_id '6.3.4'
\echo "Executing check 6.3.4"
\ir ../queries/sql/sqlserver_user_options_flag_set.sql
\set check_id '6.3.5'
\echo "Executing check 6.3.5"
\ir ../queries/sql/sqlserver_remote_access_flag_on.sql
\set check_id '6.3.6'
\echo "Executing check 6.3.6"
\ir ../queries/sql/sqlserver_trace_flag_on.sql
\set check_id '6.3.7'
\echo "Executing check 6.3.7"
\ir ../queries/sql/sqlserver_contained_database_authentication_flag_on.sql
\set check_id '6.4'
\echo "Executing check 6.4"
\ir ../queries/sql/db_instance_without_ssl.sql
\set check_id '6.5'
\echo "Executing check 6.5"
\ir ../queries/sql/db_instance_publicly_accessible.sql
\set check_id '6.6'
\echo "Executing check 6.6"
\ir ../queries/sql/db_instance_with_public_ip.sql
\set check_id '6.7'
\echo "Executing check 6.7"
\ir ../queries/sql/db_instances_without_backups.sql
