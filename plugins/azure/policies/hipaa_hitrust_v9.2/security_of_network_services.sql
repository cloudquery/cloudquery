\echo "Security of Network Services"
\set check_id "0835.09n1Organizational.1 - 09.n - 1"
\echo :check_id
\i queries/compute/windows_machines_without_data_collection_agent.sql
\set check_id "0835.09n1Organizational.1 - 09.n - 2"
\echo :check_id
\i queries/compute/vms_no_resource_manager.sql
\set check_id "0836.09.n2Organizational.1 - 09.n - 1"
\echo :check_id
\i queries/compute/linux_machines_without_data_collection_agent.sql
\set check_id "0837.09.n2Organizational.2 - 09.n - 1"
\echo :check_id
\i queries/account/locations_without_network_watchers.sql
\set check_id "0885.09n2Organizational.3 - 09.n - 1"
\echo :check_id
\i queries/compute/linux_machines_without_data_collection_agent.sql
\set check_id "0886.09n2Organizational.4 - 09.n - 1"
\echo :check_id
\i queries/account/locations_without_network_watchers.sql
\set check_id "0887.09n2Organizational.5 - 09.n - 1"
\echo :check_id
\i queries/compute/windows_machines_without_data_collection_agent.sql
\set check_id "0888.09n2Organizational.6 - 09.n - 1"
\echo :check_id
\i queries/account/locations_without_network_watchers.sql
