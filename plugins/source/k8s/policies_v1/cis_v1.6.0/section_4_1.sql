\set framework 'cis_v1.6.0'
\echo "Creating CIS V1.6.0 Section 4 Views"
\ir ../views/project_policy_members.sql
\echo "Executing CIS V1.6.0 Section 4"
\echo "4 Worker Nodes"

\echo "4.1 Worker Node Configuration Files"
\set check_id '4.1.1'
\echo "Executing check 4.1.1"\echo "Ensure that the kubelet service file permissions are set to 644 or more restrictive (Automated)................................................................................................................ 169"
\ir ../queries/manual.sql

\set check_id '4.1.2'
\echo "Executing check 4.1.2"\echo "Ensure that the kubelet service file ownership is set to root:root (Automated) ...................................................................................................................................... 171"
\ir ../queries/manual.sql

\set check_id '4.1.3'
\echo "Executing check 4.1.3"
echo " If proxy kubeconfig file exists ensure permissions are set to 644 or more restrictive (Manual) ....................................................................................................................... 173"
\ir ../queries/manual.sql

\set check_id '4.1.4'
\echo "Executing check 4.1.4"
echo " If proxy kubeconfig file exists ensure ownership is set to root:root (Manual) ................................................................................................................................................................. 175"
\ir ../queries/manual.sql

\set check_id '4.1.5'
\echo "Executing check 4.1.5"\echo "Ensure that the --kubeconfig kubelet.conf file permissions are set to 644 or more restrictive (Automated) ................................................................................................... 177"
\ir ../queries/manual.sql

\set check_id '4.1.6'
\echo "Executing check 4.1.6"\echo "Ensure that the --kubeconfig kubelet.conf file ownership is set to root:root (Manual).............................................................................................................................................. 179"
\ir ../queries/manual.sql

\set check_id '4.1.7'
\echo "Executing check 4.1.7"\echo "Ensure that the certificate authorities file permissions are set to 644 or more restrictive (Manual) ....................................................................................................................... 181"
\ir ../queries/manual.sql

\set check_id '4.1.8'
\echo "Executing check 4.1.8"\echo "Ensure that the client certificate authorities file ownership is set to root:root (Manual).............................................................................................................................................. 183"
\ir ../queries/manual.sql

\set check_id '4.1.9'
\echo "Executing check 4.1.9"\echo "Ensure that the kubelet --config configuration file has permissions set to 644 or more restrictive (Automated).............................................................................................. 185"
\ir ../queries/manual.sql

\set check_id '4.1.10'
\echo "Executing check 4.1.10"
\echo "Ensure that the kubelet --config configuration file ownership is set to root:root (Automated) .................................................................................................................. 187"
\ir ../queries/manual.sql

\set check_id '4.2.2'
\echo "Executing check 4.2.2"\echo "Ensure that the --authorization-mode argument is not set to AlwaysAllow (Automated) ...................................................................................................................................... 191"
\ir ../queries/manual.sql

\set check_id '4.2.3'
\echo "Executing check 4.2.3"\echo "Ensure that the --client-ca-file argument is set as appropriate (Automated) ................................................................................................................................................................. 193"
\ir ../queries/manual.sql

\set check_id '4.2.4'
\echo "Executing check 4.2.4"
echo " Verify that the --read-only-port argument is set to 0 (Manual) ...................... 195"
\ir ../queries/manual.sql

\set check_id '4.2.5'
\echo "Executing check 4.2.5"\echo "Ensure that the --streaming-connection-idle-timeout argument is not set to 0 (Manual) .......................................................................................................................................... 197"
\ir ../queries/manual.sql

\set check_id '4.2.6'
\echo "Executing check 4.2.6"\echo "Ensure that the --protect-kernel-defaults argument is set to true (Automated) ...................................................................................................................................... 199"
\ir ../queries/manual.sql

\set check_id '4.2.7'
\echo "Executing check 4.2.7"\echo "Ensure that the --make-iptables-util-chains argument is set to true (Automated) ...................................................................................................................................... 201"
\ir ../queries/manual.sql

\set check_id '4.2.8'
\echo "Executing check 4.2.8"\echo "Ensure that the --hostname-override argument is not set (Manual) ............ 203"
\ir ../queries/manual.sql

\set check_id '4.2.9'
\echo "Executing check 4.2.9"\echo "Ensure that the --event-qps argument is set to 0 or a level which ensures appropriate event capture (Manual) ...................................................................................... 205"
\ir ../queries/manual.sql

\set check_id '4.2.10'
\echo "Executing check 4.2.10"
\echo "Ensure that the --tls-cert-file and --tls-private-key-file arguments are set as appropriate (Manual).................................................................................................................... 207"
\ir ../queries/manual.sql

\set check_id '4.2.11'
\echo "Executing check 4.2.11"
\echo "Ensure that the --rotate-certificates argument is not set to false (Manual) ................................................................................................................................................................. 209"
\ir ../queries/manual.sql

\set check_id '4.2.12'
\echo "Executing check 4.2.12"
\echo " Verify that the RotateKubeletServerCertificate argument is set to true (Manual).............................................................................................................................................. 211"
\ir ../queries/manual.sql

\set check_id '4.2.13'
\echo "Executing check 4.2.13"
\echo "Ensure that the Kubelet only makes use of Strong Cryptographic Ciphers (Manual).............................................................................................................................................. 213"
\ir ../queries/manual.sql

