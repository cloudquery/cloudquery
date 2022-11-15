\set framework 'cis_v1.6.0'
\echo "Creating CIS V1.6.0 Section 1 Views"
-- \ir ../views/project_policy_members.sql
\echo "Executing CIS V1.6.0 Section 1"
\echo "Control Plane Components"
\echo "1.3 Controller Manager"

\set check_id '1.3.1'
\echo "Executing check 1.3.1"
\echo "Ensure that the --terminated-pod-gc-threshold argument is set as appropriate (Manual)"
\ir ../queries/manual.sql

\set check_id '1.3.2'
\echo "Executing check 1.3.2"
\echo "Ensure that the --profiling argument is set to false (Automated)"
\ir ../queries/manual.sql

\set check_id '1.3.3'
\echo "Executing check 1.3.3"
\echo "Ensure that the --use-service-account-credentials argument is set to true (Automated)"
\ir ../queries/manual.sql

\set check_id '1.3.4'
\echo "Executing check 1.3.4"
\echo "Ensure that the --service-account-private-key-file argument is set as appropriate (Automated)"
\ir ../queries/manual.sql

\set check_id '1.3.5'
\echo "Executing check 1.3.5"
\echo "Ensure that the --root-ca-file argument is set as appropriate (Automated)"
\ir ../queries/manual.sql

\set check_id '1.3.6'
\echo "Executing check 1.3.6"
\echo "Ensure that the RotateKubeletServerCertificate argument is set to true (Automated)"
\ir ../queries/manual.sql

\set check_id '1.3.7'
\echo "Executing check 1.3.7"
\echo "Ensure that the --bind-address argument is set to 127.0.0.1 (Automated)"
\ir ../queries/manual.sql