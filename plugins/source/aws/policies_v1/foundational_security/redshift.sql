\set check_id 'Redshift.1'
\echo "Executing check Redshift.1"
\ir ../queries/redshift/cluster_publicly_accessible.sql

\set check_id 'Redshift.2'
\echo "Executing check Redshift.2"
\ir ../queries/redshift/clusters_should_be_encrypted_in_transit.sql

\set check_id 'Redshift.3'
\echo "Executing check Redshift.3"
\ir ../queries/redshift/clusters_should_have_automatic_snapshots_enabled.sql

\set check_id 'Redshift.4'
\echo "Executing check Redshift.4"
\ir ../queries/redshift/clusters_should_have_audit_logging_enabled.sql

\set check_id 'Redshift.6'
\echo "Executing check Redshift.6"
\ir ../queries/redshift/clusters_should_have_automatic_upgrades_to_major_versions_enabled.sql

\set check_id 'Redshift.7'
\echo "Executing check Redshift.7"
\ir ../queries/redshift/clusters_should_use_enhanced_vpc_routing.sql
