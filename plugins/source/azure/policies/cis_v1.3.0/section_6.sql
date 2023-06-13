\echo "Executing CIS V1.3.0 Section 5"
\ir ../views/nsg_rules_dest_ports.sql
\set check_id '6.1'
\ir ../queries/network/rdp_services_are_restricted_from_the_internet.sql
\set check_id '6.2'
\ir ../queries/network/ssh_services_are_restricted_from_the_internet.sql
\set check_id '6.3'
\ir ../queries/sql/no_sql_allow_ingress_from_any_ip.sql
\set check_id '6.4'
\ir ../queries/network/nsg_log_retention_period.sql
\set check_id '6.6'
\ir ../queries/network/udp_services_are_restricted_from_the_internet.sql
