\echo "Network Connection Control"
\set check_id "0809.01n2Organizational.1234 - 01.n - 2"
\echo :check_id
\i queries/web/api_app_should_only_be_accessible_over_https.sql
\set check_id "0809.01n2Organizational.1234 - 01.n - 3"
\echo :check_id
\i queries/mysql/enforce_ssl_connection_should_be_enabled_for_mysql_database_servers.sql
\set check_id "0809.01n2Organizational.1234 - 01.n - 4"
\echo :check_id
\i queries/postgresql/enforce_ssl_connection_should_be_enabled_for_postgresql_database_servers.sql
\set check_id "0809.01n2Organizational.1234 - 01.n - 5"
\echo :check_id
\i queries/web/function_app_should_only_be_accessible_over_https.sql
\set check_id "0809.01n2Organizational.1234 - 01.n - 6"
\echo :check_id
\i queries/compute/internet-facing_virtual_machines_should_be_protected_with_network_security_groups.sql
\set check_id "0809.01n2Organizational.1234 - 01.n - 7"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_api_app.sql
\set check_id "0809.01n2Organizational.1234 - 01.n - 8"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_function_app.sql
\set check_id "0809.01n2Organizational.1234 - 01.n - 9"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_web_app.sql
\set check_id "0809.01n2Organizational.1234 - 01.n - 10"
\echo :check_id
\i queries/redis/only_secure_connections_to_your_azure_cache_for_redis_should_be_enabled.sql
\set check_id "0809.01n2Organizational.1234 - 01.n - 11"
\echo :check_id
\i queries/storage/secure_transfer_to_storage_accounts_should_be_enabled.sql
\set check_id "0809.01n2Organizational.1234 - 01.n - 12"
\echo :check_id
\i queries/network/subnets_without_nsg_associated.sql
\set check_id "0809.01n2Organizational.1234 - 01.n - 13"
\echo :check_id
\i queries/compute/vms_without_approved_networks.sql
\set check_id "0809.01n2Organizational.1234 - 01.n - 14"
\echo :check_id
\i queries/web/web_application_should_only_be_accessible_over_https.sql

\set check_id "0810.01n2Organizational.5 - 01.n - 2"
\echo :check_id
\i queries/web/api_app_should_only_be_accessible_over_https.sql
\set check_id "0810.01n2Organizational.5 - 01.n - 3"
\echo :check_id
\i queries/mysql/enforce_ssl_connection_should_be_enabled_for_mysql_database_servers.sql
\set check_id "0810.01n2Organizational.5 - 01.n - 4"
\echo :check_id
\i queries/postgresql/enforce_ssl_connection_should_be_enabled_for_postgresql_database_servers.sql
\set check_id "0810.01n2Organizational.5 - 01.n - 5"
\echo :check_id
\i queries/web/function_app_should_only_be_accessible_over_https.sql
\set check_id "0810.01n2Organizational.5 - 01.n - 6"
\echo :check_id
\i queries/compute/internet-facing_virtual_machines_should_be_protected_with_network_security_groups.sql
\set check_id "0810.01n2Organizational.5 - 01.n - 7"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_api_app.sql
\set check_id "0810.01n2Organizational.5 - 01.n - 8"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_function_app.sql
\set check_id "0810.01n2Organizational.5 - 01.n - 9"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_web_app.sql
\set check_id "0810.01n2Organizational.5 - 01.n - 10"
\echo :check_id
\i queries/redis/only_secure_connections_to_your_azure_cache_for_redis_should_be_enabled.sql
\set check_id "0810.01n2Organizational.5 - 01.n - 11"
\echo :check_id
\i queries/storage/secure_transfer_to_storage_accounts_should_be_enabled.sql
\set check_id "0810.01n2Organizational.5 - 01.n - 12"
\echo :check_id
\i queries/network/subnets_without_nsg_associated.sql
\set check_id "0810.01n2Organizational.5 - 01.n - 13"
\echo :check_id
\i queries/compute/vms_without_approved_networks.sql
\set check_id "0810.01n2Organizational.5 - 01.n - 14"
\echo :check_id
\i queries/web/web_application_should_only_be_accessible_over_https.sql


\set check_id "0811.01n2Organizational.6 - 01.n - 2"
\echo :check_id
\i queries/web/api_app_should_only_be_accessible_over_https.sql
\set check_id "0811.01n2Organizational.6 - 01.n - 2"
\echo :check_id
\i queries/mysql/enforce_ssl_connection_should_be_enabled_for_mysql_database_servers.sql
\set check_id "0811.01n2Organizational.6 - 01.n - 3"
\echo :check_id
\i queries/mysql/enforce_ssl_connection_should_be_enabled_for_mysql_database_servers.sql
\set check_id "0811.01n2Organizational.6 - 01.n - 4"
\echo :check_id
\i queries/postgresql/enforce_ssl_connection_should_be_enabled_for_postgresql_database_servers.sql
\set check_id "0811.01n2Organizational.6 - 01.n - 5"
\echo :check_id
\i queries/web/function_app_should_only_be_accessible_over_https.sql
\set check_id "0811.01n2Organizational.6 - 01.n - 6"
\echo :check_id
\i queries/compute/internet-facing_virtual_machines_should_be_protected_with_network_security_groups.sql
\set check_id "0811.01n2Organizational.6 - 01.n - 7"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_api_app.sql
\set check_id "0811.01n2Organizational.6 - 01.n - 8"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_function_app.sql
\set check_id "0811.01n2Organizational.6 - 01.n - 9"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_web_app.sql
\set check_id "0811.01n2Organizational.6 - 01.n - 10"
\echo :check_id
\i queries/redis/only_secure_connections_to_your_azure_cache_for_redis_should_be_enabled.sql
\set check_id "0811.01n2Organizational.6 - 01.n - 12"
\echo :check_id
\i queries/network/subnets_without_nsg_associated.sql
\set check_id "0811.01n2Organizational.6 - 01.n - 13"
\echo :check_id
\i queries/compute/vms_without_approved_networks.sql
\set check_id "0811.01n2Organizational.6 - 01.n - 14"
\echo :check_id
\i queries/web/web_application_should_only_be_accessible_over_https.sql

\set check_id "0812.01n2Organizational.8 - 01.n - 2"
\echo :check_id
\i queries/web/api_app_should_only_be_accessible_over_https.sql
\set check_id "0812.01n2Organizational.8 - 01.n - 3"
\echo :check_id
\i queries/mysql/enforce_ssl_connection_should_be_enabled_for_mysql_database_servers.sql
\set check_id "0812.01n2Organizational.8 - 01.n - 4"
\echo :check_id
\i queries/postgresql/enforce_ssl_connection_should_be_enabled_for_postgresql_database_servers.sql
\set check_id "0812.01n2Organizational.8 - 01.n - 5"
\echo :check_id
\i queries/web/function_app_should_only_be_accessible_over_https.sql
\set check_id "0812.01n2Organizational.8 - 01.n - 6"
\echo :check_id
\i queries/compute/internet-facing_virtual_machines_should_be_protected_with_network_security_groups.sql
\set check_id "0812.01n2Organizational.8 - 01.n - 7"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_api_app.sql
\set check_id "0812.01n2Organizational.8 - 01.n - 8"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_function_app.sql
\set check_id "0812.01n2Organizational.8 - 01.n - 9"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_web_app.sql
\set check_id "0812.01n2Organizational.8 - 01.n - 10"
\echo :check_id
\i queries/redis/only_secure_connections_to_your_azure_cache_for_redis_should_be_enabled.sql
\set check_id "0812.01n2Organizational.8 - 01.n - 11"
\echo :check_id
\i queries/storage/secure_transfer_to_storage_accounts_should_be_enabled.sql
\set check_id "0812.01n2Organizational.8 - 01.n - 12"
\echo :check_id
\i queries/network/subnets_without_nsg_associated.sql
\set check_id "0812.01n2Organizational.8 - 01.n - 13"
\echo :check_id
\i queries/compute/vms_without_approved_networks.sql
\set check_id "0812.01n2Organizational.8 - 01.n - 14"
\echo :check_id
\i queries/web/web_application_should_only_be_accessible_over_https.sql

\set check_id "0814.01n1Organizational.12 - 01.n - 2"
\echo :check_id
\i queries/web/api_app_should_only_be_accessible_over_https.sql
\set check_id "0814.01n1Organizational.12 - 01.n - 3"
\echo :check_id
\i queries/mysql/enforce_ssl_connection_should_be_enabled_for_mysql_database_servers.sql
\set check_id "0814.01n1Organizational.12 - 01.n - 4"
\echo :check_id
\i queries/postgresql/enforce_ssl_connection_should_be_enabled_for_postgresql_database_servers.sql
\set check_id "0814.01n1Organizational.12 - 01.n - 5"
\echo :check_id
\i queries/web/function_app_should_only_be_accessible_over_https.sql
\set check_id "0814.01n1Organizational.12 - 01.n - 6"
\echo :check_id
\i queries/compute/internet-facing_virtual_machines_should_be_protected_with_network_security_groups.sql
\set check_id "0814.01n1Organizational.12 - 01.n - 7"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_api_app.sql
\set check_id "0814.01n1Organizational.12 - 01.n - 8"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_function_app.sql
\set check_id "0814.01n1Organizational.12 - 01.n - 9"
\echo :check_id
\i queries/web/latest_tls_version_should_be_used_in_your_web_app.sql
\set check_id "0814.01n1Organizational.12 - 01.n - 10"
\echo :check_id
\i queries/redis/only_secure_connections_to_your_azure_cache_for_redis_should_be_enabled.sql
\set check_id "0814.01n1Organizational.12 - 01.n - 11"
\echo :check_id
\i queries/storage/secure_transfer_to_storage_accounts_should_be_enabled.sql
\set check_id "0814.01n1Organizational.12 - 01.n - 12"
\echo :check_id
\i queries/network/subnets_without_nsg_associated.sql
\set check_id "0814.01n1Organizational.12 - 01.n - 13"
\echo :check_id
\i queries/compute/vms_without_approved_networks.sql
\set check_id "0814.01n1Organizational.12 - 01.n - 14"
\echo :check_id
\i queries/web/web_application_should_only_be_accessible_over_https.sql
