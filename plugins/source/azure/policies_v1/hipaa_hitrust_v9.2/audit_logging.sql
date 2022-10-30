\echo "Audit Logging"
\set check_id '1202.09aa1System.1 - 09.aa - 1'
\echo :check_id
\ir ../queries/datalake/datalake_storage_accounts_with_disabled_logging.sql
\set check_id '1203.09aa1System.2 - 09.aa - 1'
\echo :check_id
\ir ../queries/logic/app_workflow_logging_enabled.sql
\set check_id '1205.09aa2System.1 - 09.aa - 1'
\echo :check_id
\ir ../queries/batch/resource_logs_in_batch_accounts_should_be_enabled.sql
\set check_id '1206.09aa2System.23 - 09.aa - 1'
\echo :check_id
\ir ../queries/compute/virtual_machine_scale_sets_without_logs.sql
\set check_id '1207.09aa2System.4 - 09.aa - 1'
\echo :check_id
\ir ../queries/streamanalytics/resource_logs_in_azure_stream_analytics_should_be_enabled.sql
\set check_id '1207.09aa2System.4 - 09.aa - 2'
\echo :check_id
\ir ../queries/eventhub/namespaces_without_logging.sql
\set check_id '1208.09aa3System.1 - 09.aa - 1'
\echo :check_id
\ir ../queries/search/resource_logs_in_search_services_should_be_enabled.sql
\set check_id '1209.09aa3System.2 - 09.aa - 1'
\echo :check_id
\ir ../queries/web/apps_with_logging_disabled.sql
\set check_id '1211.09aa3System.4 - 09.aa - 1'
\echo :check_id
\ir ../queries/sql/sqlserverauditing_audit.sql
\set check_id '1211.09aa3System.4 - 09.aa - 2'
\echo :check_id
\ir ../queries/keyvault/hsms_without_logging.sql
\set check_id '1211.09aa3System.4 - 09.aa - 3'
\echo :check_id
\ir ../queries/keyvault/vaults_without_logging.sql
