\echo "Identification of Risks Related to External Parties"
\set check_id "1401.05i1Organizational.1239 - 05.i - 1"
\echo :check_id
\i queries/storage/secure_transfer_to_storage_accounts_should_be_enabled.sql
\set check_id "1402.05i1Organizational.45 - 05.i - 1"
\echo :check_id
\i queries/web/function_app_should_only_be_accessible_over_https.sql
\set check_id "1403.05i1Organizational.67 - 05.i - 1"
\echo :check_id
\i queries/web/web_application_should_only_be_accessible_over_https.sql
\set check_id "1418.05i1Organizational.8 - 05.i - 1"
\echo :check_id
\i queries/mysql/enforce_ssl_connection_should_be_enabled_for_mysql_database_servers.sql
\set check_id "1450.05i2Organizational.2 - 05.i - 1"
\echo :check_id
\i queries/postgresql/enforce_ssl_connection_should_be_enabled_for_postgresql_database_servers.sql
\set check_id "1451.05iCSPOrganizational.2 - 05.i - 1"
\echo :check_id
\i queries/redis/only_secure_connections_to_your_azure_cache_for_redis_should_be_enabled.sql