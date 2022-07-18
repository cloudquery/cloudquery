\echo "Segregation in Networks"
\set check_id "0805.01m1Organizational.12 - 01.m - 1"
\echo :check_id
\i queries/container/containers_without_virtual_service_endpoint.sql
\set check_id "0805.01m1Organizational.12 - 01.m - 2"
\echo :check_id
\i queries/network/virtualnetworkserviceendpoint_appservice_auditifnotexists.sql
\set check_id "0805.01m1Organizational.12 - 01.m - 3"
\echo :check_id
\i queries/cosmosdb/cosmos_db_should_use_a_virtual_network_service_endpoint.sql
\set check_id "0805.01m1Organizational.12 - 01.m - 4"
\echo :check_id
\i queries/eventhub/event_hub_should_use_a_virtual_network_service_endpoint.sql
\set check_id "0805.01m1Organizational.12 - 01.m - 5"
\echo :check_id
\i queries/network/gateway_subnets_should_not_be_configured_with_a_network_security_group.sql
\set check_id "0805.01m1Organizational.12 - 01.m - 6"
\echo :check_id
\i queries/compute/internet-facing_virtual_machines_should_be_protected_with_network_security_groups.sql
\set check_id "0805.01m1Organizational.12 - 01.m - 7"
\echo :check_id
\i queries/keyvault/vaults_with_no_service_endpoint.sql
\set check_id "0805.01m1Organizational.12 - 01.m - 8"
\echo :check_id
\i queries/sql/sql_servers_with_no_service_endpoint.sql
\set check_id "0805.01m1Organizational.12 - 01.m - 9"
\echo :check_id
\i queries/storage/accounts_with_no_service_endpoint_associated.sql
\set check_id "0805.01m1Organizational.12 - 01.m - 10"
\echo :check_id
\i queries/network/subnets_without_nsg_associated.sql
\set check_id "0805.01m1Organizational.12 - 01.m - 11"
\echo :check_id
\i queries/compute/vms_without_approved_networks.sql

\set check_id "0806.01m2Organizational.12356 - 01.m - 1"
\echo :check_id
\i queries/container/containers_without_virtual_service_endpoint.sql
\set check_id "0806.01m2Organizational.12356 - 01.m - 2"
\echo :check_id
\i queries/network/virtualnetworkserviceendpoint_appservice_auditifnotexists.sql
\set check_id "0806.01m2Organizational.12356 - 01.m - 3"
\echo :check_id
\i queries/cosmosdb/cosmos_db_should_use_a_virtual_network_service_endpoint.sql
\set check_id "0806.01m2Organizational.12356 - 01.m - 4"
\echo :check_id
\i queries/eventhub/event_hub_should_use_a_virtual_network_service_endpoint.sql
\set check_id "0806.01m2Organizational.12356 - 01.m - 6"
\echo :check_id
\i queries/compute/internet-facing_virtual_machines_should_be_protected_with_network_security_groups.sql
\set check_id "0806.01m2Organizational.12356 - 01.m - 7"
\echo :check_id
\i queries/keyvault/vaults_with_no_service_endpoint.sql
\set check_id "0806.01m2Organizational.12356 - 01.m - 9"
\echo :check_id
\i queries/storage/accounts_with_no_service_endpoint_associated.sql
\set check_id "0806.01m2Organizational.12356 - 01.m - 10"
\echo :check_id
\i queries/network/subnets_without_nsg_associated.sql
\set check_id "0806.01m2Organizational.12356 - 01.m - 11"
\echo :check_id
\i queries/compute/vms_without_approved_networks.sql
