\set framework 'cis_v1.6.0'
\echo "Creating CIS V1.6.0 Section 2 Views"
-- \ir ../views/project_policy_members.sql
\echo "Executing CIS V1.6.0 Section 2"
\echo "2 etcd"

\set check_id '2.1'
\echo "Executing check 2.1"
\echo "Ensure that the --cert-file and --key-file arguments are set as appropriate (Automated)"
\ir ../queries/manual.sql

\set check_id '2.2'
\echo "Executing check 2.2"
\echo "Ensure that the --client-cert-auth argument is set to true (Automated)"
\ir ../queries/manual.sql

\set check_id '2.3'
\echo "Executing check 2.3"
\echo "Ensure that the --auto-tls argument is not set to true (Automated)"
\ir ../queries/manual.sql

\set check_id '2.4'
\echo "Executing check 2.4"
\echo "Ensure that the --peer-cert-file and --peer-key-file arguments are set as appropriate (Automated)"
\ir ../queries/manual.sql

\set check_id '2.5'
\echo "Executing check 2.5"
\echo "Ensure that the --peer-client-cert-auth argument is set to true (Automated)"
\ir ../queries/manual.sql

\set check_id '2.6'
\echo "Executing check 2.6"
\echo "Ensure that the --peer-auto-tls argument is not set to true (Automated)"
\ir ../queries/manual.sql

\set check_id '2.7'
\echo "Executing check 2.7"
\echo "Ensure that a unique Certificate Authority is used for etcd (Manual)"
\ir ../queries/manual.sql