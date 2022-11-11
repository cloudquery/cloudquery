\set framework 'cis_v1.2.0'
\echo "Creating CIS V1.2.0 Section 3 Views"
\ir ../views/firewall_allowed_rules.sql
\echo "Executing CIS V1.2.0 Section 3"
\set check_id '3.1'
\echo "Executing check 3.1"
\ir ../queries/compute/default_network_exist.sql
\set check_id '3.2'
\echo "Executing check 3.2"
\ir ../queries/compute/legacy_network_exist.sql
\set check_id '3.3'
\echo "Executing check 3.3"
\ir ../queries/dns/zones_with_dnssec_disabled.sql
\set check_id '3.4'
\echo "Executing check 3.4"
\ir ../queries/dns/key_signing_with_rsasha1.sql
\set check_id '3.5'
\echo "Executing check 3.5"
\ir ../queries/dns/zone_signing_with_rsasha1.sql
\set check_id '3.6'
\echo "Executing check 3.6"
\ir ../queries/compute/ssh_access_permitted.sql
\set check_id '3.7'
\echo "Executing check 3.7"
\ir ../queries/compute/rdp_access_permitted.sql
\set check_id '3.8'
\echo "Executing check 3.8"
\ir ../queries/compute/flow_logs_disabled_in_vpc.sql
\set check_id '3.9'
\echo "Executing check 3.9"
-- \ir ../queries/compute/ssl_proxy_with_weak_cipher.sql
\set check_id '3.10'
\echo "Executing check 3.10"
\ir ../queries/compute/allow_traffic_behind_iap.sql
