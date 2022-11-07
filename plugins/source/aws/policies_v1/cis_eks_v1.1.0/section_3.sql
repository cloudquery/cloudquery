\set framework 'cis_eks_v1.1.0'
\echo "Creating CIS EKS V1.1.0 Section 3 Views"
\ir ../views/project_policy_members.sql
\echo "Executing CIS EKS V1.1.0 Section 3"


\echo "3 Worker Nodes"
\echo "3.1 Worker Node Configuration Files"
\set check_id '3.1.1'
\echo "Executing check 3.1.1"
\echo "Ensure that the proxy kubeconfig file permissions are set to 644 or more restrictive (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id '3.1.2'
\echo "Executing check 3.1.2"
\echo "Ensure that the proxy kubeconfig file ownership is set to root:root (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id '3.1.3'
\echo "Executing check 3.1.3"
\echo "Ensure that the kubelet configuration file has permissions set to 644 or more restrictive (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id '3.1.4'
\echo "Executing check 3.1.4"
\echo "Ensure that the kubelet configuration file ownership is set to root:root (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information


\echo "3.2 Kubelet"

\set check_id '3.2.1'
\echo "Executing check 3.2.1"
\echo "Ensure that the --anonymous-auth argument is set to false (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id '3.2.2'
\echo "Executing check 3.2.2"
\echo "Ensure that the --authorization-mode argument is not set to AlwaysAllow (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id '3.2.3'
\echo "Executing check 3.2.3"
\echo "Ensure that the --client-ca-file argument is set as appropriate (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id '3.2.4'
\echo "Executing check 3.2.4"
\echo "Ensure that the --read-only-port is secured (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id '3.2.5'
\echo "Executing check 3.2.5"
\echo "Ensure that the --streaming-connection-idle-timeout argument is not set to 0 (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id '3.2.6'
\echo "Executing check 3.2.6"
\echo "Ensure that the --protect-kernel-defaults argument is set to true (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id '3.2.7'
\echo "Executing check 3.2.7"
\echo "Ensure that the --make-iptables-util-chains argument is set to true (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id '3.2.8'
\echo "Executing check 3.2.8"
\echo "Ensure that the --hostname-override argument is not set (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id '3.2.9'
\echo "Executing check 3.2.9"
\echo "Ensure that the --eventRecordQPS argument is set to 0 or a level which ensures appropriate event capture (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id '3.2.10'
\echo "Executing check 3.2.10"
\echo "Ensure that the --rotate-certificates argument is not set to false (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id '3.2.11'
\echo "Executing check 3.2.11"
\echo "Ensure that the RotateKubeletServerCertificate argument is set to true (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information


\echo "3.3 Container Optimized OS"
\set check_id '3.3.1'
\echo "Executing check 3.3.1"
\echo "Prefer using Container-Optimized OS when possible (Manual)"
\ir ../queries/manual.sql
-- no generic check