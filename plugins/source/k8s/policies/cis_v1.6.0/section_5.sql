\set framework 'cis_v1.6.0'

\set check_id '5.1.6' 
\ir ../queries/pod_security/pod_service_account_token_disabled.sql

\set check_id '5.2.1' 
\ir ../queries/pod_security/pod_container_privilege_disabled.sql

\set check_id '5.2.2' 
\ir ../queries/pod_security/pod_hostpid_hostipc_sharing_disabled.sql

\set check_id '5.2.3' 
\ir ../queries/pod_security/pod_hostpid_hostipc_sharing_disabled.sql

\set check_id '5.2.4' 
\ir ../queries/pod_security/pod_host_network_access_disabled.sql

\set check_id '5.2.5' 
\ir ../queries/pod_security/pod_container_privilege_escalation_disabled.sql

\set check_id '5.2.6' 
\ir ../queries/pod_security/pod_non_root_container.sql

\set check_id '5.3.2.1'
\ir ../queries/network_hardening/network_policy_default_deny_ingress.sql
\set check_id '5.3.2.2'
\ir ../queries/network_hardening/network_policy_default_deny_egress.sql

\set check_id '5.7.2.1'
\ir ../queries/pod_security/statefulset_default_seccomp_profile_enabled.sql
\set check_id '5.7.2.2'
\ir ../queries/pod_security/replicaset_default_seccomp_profile_enabled.sql
\set check_id '5.7.2.3'
\ir ../queries/pod_security/pod_container_default_seccomp_profile_enabled.sql
\set check_id '5.7.2.4'
\ir ../queries/pod_security/job_container_default_seccomp_profile_enabled.sql
\set check_id '5.7.2.5'
\ir ../queries/pod_security/deployment_container_default_seccomp_profile_enabled.sql
\set check_id '5.7.2.6'
\ir ../queries/pod_security/daemonset_default_seccomp_profile_enabled.sql

\set check_id '5.7.4.1'
\ir ../queries/pod_security/daemonset_default_namespace.sql
\set check_id '5.7.4.2'
\ir ../queries/pod_security/deployment_container_uses_default_namespace.sql
\set check_id '5.7.4.3'
\ir ../queries/pod_security/pod_container_uses_default_namespace.sql
\set check_id '5.7.4.4'
\ir ../queries/pod_security/job_container_uses_default_namespace.sql
\set check_id '5.7.4.5'
\ir ../queries/pod_security/replicaset_uses_default_namespace.sql
\set check_id '5.7.4.6'
\ir ../queries/pod_security/statefulset_uses_default_namespace.sql