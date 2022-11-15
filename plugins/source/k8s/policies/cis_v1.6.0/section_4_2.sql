\set framework 'cis_v1.6.0'
\echo "Creating CIS V1.6.0 Section 4 Views"
-- \ir ../views/project_policy_members.sql
\echo "Executing CIS V1.6.0 Section 4"
\echo "4 Worker Nodes"

\echo "4.2 Kubelet"
\set check_id '4.2.1'
\echo "Executing check 4.2.1"
\echo "Ensure that the --anonymous-auth argument is set to false (Automated)"
\ir ../queries/manual.sql

\set check_id '4.2.2'
\echo "Executing check 4.2.2"
\echo "Ensure that the --authorization-mode argument is not set to AlwaysAllow (Automated)"
\ir ../queries/manual.sql

\set check_id '4.2.3'
\echo "Executing check 4.2.3"
\echo "Ensure that the --client-ca-file argument is set as appropriate (Automated)"
\ir ../queries/manual.sql

\set check_id '4.2.4'
\echo "Executing check 4.2.4"
\echo "Verify that the --read-only-port argument is set to 0 (Manual)"
\ir ../queries/manual.sql

\set check_id '4.2.5'
\echo "Executing check 4.2.5"
\echo "Ensure that the --streaming-connection-idle-timeout argument is not set to 0 (Manual)"
\ir ../queries/manual.sql

\set check_id '4.2.6'
\echo "Executing check 4.2.6"
\echo "Ensure that the --protect-kernel-defaults argument is set to true (Automated)"
\ir ../queries/manual.sql

\set check_id '4.2.7'
\echo "Executing check 4.2.7"
\echo "Ensure that the --make-iptables-util-chains argument is set to true (Automated)"
\ir ../queries/manual.sql

\set check_id '4.2.8'
\echo "Executing check 4.2.8"
\echo "Ensure that the --hostname-override argument is not set (Manual)"
\ir ../queries/manual.sql

\set check_id '4.2.9'
\echo "Executing check 4.2.9"
\echo "Ensure that the --event-qps argument is set to 0 or a level which ensures appropriate event capture (Manual)"
\ir ../queries/manual.sql

\set check_id '4.2.10'
\echo "Executing check 4.2.10"
\echo "Ensure that the --tls-cert-file and --tls-private-key-file arguments are set as appropriate (Manual)"
\ir ../queries/manual.sql

\set check_id '4.2.11'
\echo "Executing check 4.2.11"
\echo "Ensure that the --rotate-certificates argument is not set to false (Manual)"
\ir ../queries/manual.sql

\set check_id '4.2.12'
\echo "Executing check 4.2.12"
\echo "Verify that the RotateKubeletServerCertificate argument is set to true (Manual)"
\ir ../queries/manual.sql

\set check_id '4.2.13'
\echo "Executing check 4.2.13"
\echo "Ensure that the Kubelet only makes use of Strong Cryptographic Ciphers (Manual)"
\ir ../queries/manual.sql