\echo "Executing On-line Transactions"
\set check_id "0943.09y1Organizational.1 - 09.y - 1"
\echo :check_id
\i queries/storage/secure_transfer_to_storage_accounts_should_be_enabled.sql
\set check_id "0946.09y2Organizational.14 - 09.y - 1"
\echo :check_id
\i queries/redis/only_secure_connections_to_your_azure_cache_for_redis_should_be_enabled.sql
\set check_id "0947.09y2Organizational.2 - 09.y - 1"
\echo :check_id
\i queries/postgresql/enforce_ssl_connection_should_be_enabled_for_postgresql_database_servers.sql
\set check_id "0948.09y2Organizational.3 - 09.y - 1"
\echo :check_id
\i queries/mysql/enforce_ssl_connection_should_be_enabled_for_mysql_database_servers.sql
\set check_id "0949.09y2Organizational.5 - 09.y - 1"
\echo :check_id
\i queries/web/api_app_should_only_be_accessible_over_https.sql
\set check_id "0949.09y2Organizational.5 - 09.y - 2"
\echo :check_id
\i queries/web/function_app_should_only_be_accessible_over_https.sql
\set check_id "0949.09y2Organizational.5 - 09.y - 3"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_api_app.sql
\set check_id "0949.09y2Organizational.5 - 09.y - 4"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_function_app.sql
\set check_id "0949.09y2Organizational.5 - 09.y - 1"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_web_app.sql
\set check_id "0949.09y2Organizational.5 - 09.y - 1"
\echo :check_id
\i queries/web/web_application_should_only_be_accessible_over_https.sql
\set check_id "0915.09s2Organizational.2 - 09.s - 1"
\echo :check_id
\i queries/web/app_client_cert_disabled.sql
\set check_id "0916.09s2Organizational.4 - 09.s - 1"
\echo :check_id
\i queries/web/cors_should_not_allow_every_resource_to_access_your_web_applications.sql
\set check_id "0960.09sCSPOrganizational.1 - 09.s - 1"
\echo :check_id
\i queries/web/cors_should_not_allow_every_resource_to_access_your_function_apps.sql
\set check_id "1325.09s1Organizational.3 - 09.s - 1"
\echo :check_id
\i queries/web/remote_debugging_should_be_turned_off_for_function_apps.sql
