\set framework 'cis_v1.6.0'
\echo "Creating CIS V1.6.0 Section 4 Views"
-- \ir ../views/project_policy_members.sql
\echo "Executing CIS V1.6.0 Section 4"
\echo "4 Worker Nodes"

\echo "4.1 Worker Node Configuration Files"
\set check_id '4.1.1'
\echo "Executing check 4.1.1"
\echo "Ensure that the kubelet service file permissions are set to 644 or more restrictive (Automated)"
\ir ../queries/manual.sql

\set check_id '4.1.2'
\echo "Executing check 4.1.2"
\echo "Ensure that the kubelet service file ownership is set to root:root (Automated)"
\ir ../queries/manual.sql

\set check_id '4.1.3'
\echo "Executing check 4.1.3"
\echo "If proxy kubeconfig file exists ensure permissions are set to 644 or more restrictive (Manual)"
\ir ../queries/manual.sql

\set check_id '4.1.4'
\echo "Executing check 4.1.4"
\echo "If proxy kubeconfig file exists ensure ownership is set to root:root (Manual)"
\ir ../queries/manual.sql

\set check_id '4.1.5'
\echo "Executing check 4.1.5"
\echo "Ensure that the --kubeconfig kubelet.conf file permissions are set to 644 or more restrictive (Automated)"
\ir ../queries/manual.sql

\set check_id '4.1.6'
\echo "Executing check 4.1.6"
\echo "Ensure that the --kubeconfig kubelet.conf file ownership is set to root:root (Manual)"
\ir ../queries/manual.sql

\set check_id '4.1.7'
\echo "Executing check 4.1.7"
\echo "Ensure that the certificate authorities file permissions are set to 644 or more restrictive (Manual)"
\ir ../queries/manual.sql

\set check_id '4.1.8'
\echo "Executing check 4.1.8"
\echo "Ensure that the client certificate authorities file ownership is set to root:root (Manual)"
\ir ../queries/manual.sql

\set check_id '4.1.9'
\echo "Executing check 4.1.9"
\echo "Ensure that the kubelet --config configuration file has permissions set to 644 or more restrictive (Automated)"
\ir ../queries/manual.sql

\set check_id '4.1.10'
\echo "Executing check 4.1.10"
\echo "Ensure that the kubelet --config configuration file ownership is set to root:root (Automated)"
\ir ../queries/manual.sql