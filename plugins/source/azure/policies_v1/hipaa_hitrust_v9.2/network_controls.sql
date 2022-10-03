\echo "Network Controls"
\set check_id '0858.09m1Organizational.4 - 09.m - 1'
\echo :check_id
\ir ../queries/network/asc_unprotectedendpoints_audit.sql
\set check_id '0858.09m1Organizational.4 - 09.m - 2'
\echo :check_id
\ir ../queries/compute/virtual_machines_without_jit_network_access_policy.sql
\set check_id '0861.09m2Organizational.67 - 09.m - 1'
\echo :check_id
\ir ../queries/network/virtualnetworkserviceendpoint_appservice_auditifnotexists.sql
\set check_id '0863.09m2Organizational.910 - 09.m - 1'
\echo :check_id
\ir ../queries/eventhub/event_hub_should_use_a_virtual_network_service_endpoint.sql
\set check_id '0864.09m2Organizational.12 - 09.m - 1'
\echo :check_id
\ir ../queries/cosmosdb/cosmos_db_should_use_a_virtual_network_service_endpoint.sql
\set check_id '0865.09m2Organizational.13 - 09.m - 1'
\echo :check_id
\ir ../queries/keyvault/vaults_with_no_service_endpoint.sql
\set check_id '0866.09m3Organizational.1516 - 09.m - 1'
\echo :check_id
\ir ../queries/storage/accounts_with_unrestricted_access.sql
\set check_id '0867.09m3Organizational.17 - 09.m - 1'
\echo :check_id
\ir ../queries/storage/accounts_with_no_service_endpoint_associated.sql
\set check_id '0868.09m3Organizational.18 - 09.m - 1'
\echo :check_id
\ir ../queries/container/containers_without_virtual_service_endpoint.sql
\set check_id '0869.09m3Organizational.19 - 09.m - 1'
\echo :check_id
\ir ../queries/container/containers_without_virtual_service_endpoint.sql
\set check_id '0870.09m3Organizational.20 - 09.m - 1'
\echo :check_id
\ir ../queries/container/containers_without_virtual_service_endpoint.sql
\set check_id '0871.09m3Organizational.22 - 09.m - 1'
\echo :check_id
\ir ../queries/container/containers_without_virtual_service_endpoint.sql
