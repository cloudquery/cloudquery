\set framework 'cis_v1.6.0'
\echo "Creating CIS V1.6.0 Section 1 Views"
-- \ir ../views/project_policy_members.sql
\echo "Executing CIS V1.6.0 Section 1"
\echo "Control Plane Components"
\echo "1.4 Scheduler"

\set check_id '1.4.1'
\echo "Executing check 1.4.1"
\echo "Ensure that the --profiling argument is set to false (Automated)"
\ir ../queries/manual.sql

\set check_id '1.4.1'
\echo "Executing check 1.4.1"
\echo "Ensure that the --bind-address argument is set to 127.0.0.1 (Automated)"
\ir ../queries/manual.sql