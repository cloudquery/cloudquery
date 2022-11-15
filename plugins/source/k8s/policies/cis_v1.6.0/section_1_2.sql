\set framework 'cis_v1.6.0'
\echo "Creating CIS V1.6.0 Section 1 Views"
-- \ir ../views/project_policy_members.sql
\echo "Executing CIS V1.6.0 Section 1"
\echo "Control Plane Components"
\echo "1.2 API SERVER"

\set check_id  '1.2.1'
\echo "Executing check 1.2.1"
\echo "Ensure that the --anonymous-auth argument is set to false (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.2'
\echo "Executing check 1.2.2"
\echo "Ensure that the --basic-auth-file argument is not set (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.3'
\echo "Executing check 1.2.3"
\echo "Ensure that the --token-auth-file parameter is not set (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.4'
\echo "Executing check 1.2.4"
\echo "Ensure that the --kubelet-https argument is set to true (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.5'
\echo "Executing check 1.2.5"
\echo "Ensure that the --kubelet-client-certificate and --kubelet-client-key arguments are set as appropriate (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.6'
\echo "Executing check 1.2.6"
\echo "Ensure that the --kubelet-certificate-authority argument is set as appropriate (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.7'
\echo "Executing check 1.2.7"
\echo "Ensure that the --authorization-mode argument is not set to AlwaysAllow (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.8'
\echo "Executing check 1.2.8"
\echo "Ensure that the --authorization-mode argument includes Node (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.9'
\echo "Executing check 1.2.9"
\echo "Ensure that the --authorization-mode argument includes RBAC (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.10'
\echo "Executing check 1.2.10"
\echo "Ensure that the admission control plugin EventRateLimit is set (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.11'
\echo "Executing check 1.2.11"
\echo "Ensure that the admission control plugin AlwaysAdmit is not set (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.12'
\echo "Executing check 1.2.12"
\echo "Ensure that the admission control plugin AlwaysPullImages is set (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.13'
\echo "Executing check 1.2.13"
\echo "Ensure that the admission control plugin SecurityContextDeny is set if PodSecurityPolicy is not used (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.14'
\echo "Executing check 1.2.14"
\echo "Ensure that the admission control plugin ServiceAccount is set (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.15'
\echo "Executing check 1.2.15"
\echo "Ensure that the admission control plugin NamespaceLifecycle is set (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.16'
\echo "Executing check 1.2.16"
\echo "Ensure that the admission control plugin PodSecurityPolicy is set (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.17'
\echo "Executing check 1.2.17"
\echo "Ensure that the admission control plugin NodeRestriction is set (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.18'
\echo "Executing check 1.2.18"
\echo "Ensure that the --insecure-bind-address argument is not set (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.19'
\echo "Executing check 1.2.19"
\echo "Ensure that the --insecure-port argument is set to 0 (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.20'
\echo "Executing check 1.2.20"
\echo "Ensure that the --secure-port argument is not set to 0 (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.21'
\echo "Executing check 1.2.21"
\echo "Ensure that the --profiling argument is set to false (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.22'
\echo "Executing check 1.2.22"
\echo "Ensure that the --audit-log-path argument is set (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.23'
\echo "Executing check 1.2.23"
\echo "Ensure that the --audit-log-maxage argument is set to 30 or as appropriate (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.24'
\echo "Executing check 1.2.24"
\echo "Ensure that the --audit-log-maxbackup argument is set to 10 or as appropriate (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.25'
\echo "Executing check 1.2.25"
\echo "Ensure that the --audit-log-maxsize argument is set to 100 or as appropriate (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.26'
\echo "Executing check 1.2.26"
\echo "Ensure that the --request-timeout argument is set as appropriate (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.27'
\echo "Executing check 1.2.27"
\echo "Ensure that the --service-account-lookup argument is set to true (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.28'
\echo "Executing check 1.2.28"
\echo "Ensure that the --service-account-key-file argument is set as appropriate (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.29'
\echo "Executing check 1.2.29"
\echo "Ensure that the --etcd-certfile and --etcd-keyfile arguments are set as appropriate (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.30'
\echo "Executing check 1.2.30"
\echo "Ensure that the --tls-cert-file and --tls-private-key-file arguments are set as appropriate (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.31'
\echo "Executing check 1.2.31"
\echo "Ensure that the --client-ca-file argument is set as appropriate (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.32'
\echo "Executing check 1.2.32"
\echo "Ensure that the --etcd-cafile argument is set as appropriate (Automated)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.33'
\echo "Executing check 1.2.33"
\echo "Ensure that the --encryption-provider-config argument is set as appropriate (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.34'
\echo "Executing check 1.2.34"
\echo "Ensure that encryption providers are appropriately configured (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information

\set check_id  '1.2.35'
\echo "Executing check 1.2.35"
\echo "Ensure that the API Server only makes use of Strong Cryptographic Ciphers (Manual)"
\ir ../queries/manual.sql
-- need runtime/filesystem information