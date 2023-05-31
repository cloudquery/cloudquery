\echo "Executing CIS V1.3.0 Section 5"
\ir ../views/nsg_rules_dest_ports.sql
\set check_id '6.1'
\ir ../queries/network/rdp_services_are_restricted_from_the_internet.sql
\set check_id '6.2'
\ir ../queries/network/ssh_services_are_restricted_from_the_internet.sql
-- \set check_id '6.3'
-- There is no firewal-rules table
-- https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/firewall-rules/list-by-server?tabs=HTTP
-- \ir ../queries/network/
-- \set check_id '6.4'
-- There is no flow-log-status table
-- https://learn.microsoft.com/en-us/rest/api/network-watcher/network-watchers/get-flow-log-status?tabs=HTTP
-- \ir ../queries/network/
\set check_id '6.6'
\ir ../queries/network/udp_services_are_restricted_from_the_internet.sql
