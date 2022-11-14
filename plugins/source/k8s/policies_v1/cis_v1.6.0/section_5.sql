\set framework 'cis_v1.6.0'
\echo "Creating CIS V1.6.0 Section 5 Views"
\echo "Executing CIS V1.6.0 Section 5"
\echo "5 Policies"

\echo "5.1 RBAC and Service Accounts"
\set check_id '5.1.1' 
\echo "Executing check 5.1.1"
\echo "Ensure that the cluster-admin role is only used where required (Manual)"
\ir ../queries/manual.sql 

\set check_id '5.1.2' 
\echo "Executing check 5.1.2"
\echo "Minimize access to secrets (Manual)"
\ir ../queries/manual.sql 

\set check_id '5.1.3' 
\echo "Executing check 5.1.3"
\echo "Minimize wildcard use in Roles and ClusterRoles (Manual)"
\ir ../queries/manual.sql 

\set check_id '5.1.4' 
\echo "Executing check 5.1.4"
\echo "Minimize access to create pods (Manual)"
\ir ../queries/manual.sql 

\set check_id '5.1.5' 
\echo "Executing check 5.1.5"
\echo "Ensure that default service accounts are not actively used. (Manual)"
\ir ../queries/manual.sql 

\set check_id '5.1.6' 
\echo "Executing check 5.1.6"
\echo "Ensure that Service Account Tokens are only mounted where necessary (Manual)"
\ir ../queries/pod_security/pod_service_account_token_disabled.sql



\echo "5.2 Pod Security Policies"
\set check_id '5.2.1' 
\echo "Executing check 5.2.1"
\echo "Minimize the admission of privileged containers (Manual)"
\ir ../queries/pod_security/pod_container_privilege_disabled.sql

\set check_id '5.2.2' 
\echo "Executing check 5.2.2"
\echo "Minimize the admission of containers wishing to share the host process ID namespace (Manual)"
\ir ../queries/pod_security/pod_hostpid_hostipc_sharing_disabled.sql

\set check_id '5.2.3' 
\echo "Executing check 5.2.3"
\echo "Minimize the admission of containers wishing to share the host IPC namespace (Manual)"
\ir ../queries/pod_security/pod_hostpid_hostipc_sharing_disabled.sql

\set check_id '5.2.4' 
\echo "Executing check 5.2.4"
\echo "Minimize the admission of containers wishing to share the host network namespace (Manual)"
\ir ../queries/pod_security/pod_host_network_access_disabled.sql

\set check_id '5.2.5' 
\echo "Executing check 5.2.5"
\echo "Minimize the admission of containers with allowPrivilegeEscalation (Manual)"
\ir ../queries/pod_security/pod_container_privilege_escalation_disabled.sql

\set check_id '5.2.6' 
\echo "Executing check 5.2.6"
\echo "Minimize the admission of root containers (Manual)"
\ir ../queries/pod_security/pod_non_root_container.sql
\set check_id '5.2.7' 
\echo "Executing check 5.2.7"
\echo "Minimize the admission of containers with the NET_RAW capability (Manual)"
\ir ../queries/manual.sql

\set check_id '5.2.8' 
\echo "Executing check 5.2.8"
\echo "Minimize the admission of containers with added capabilities (Manual)"
\ir ../queries/manual.sql

\set check_id '5.2.9' 
\echo "Executing check 5.2.9"
\echo "Minimize the admission of containers with capabilities assigned (Manual)"
\ir ../queries/manual.sql


\echo "5.3 Network Policies and CNI"
\set check_id '5.3.1'
\echo "Executing check 5.3.1"
\echo "Ensure that the CNI in use supports Network Policies (Manual)"
\ir ../queries/manual.sql

\set check_id '5.3.2'
\echo "Executing check 5.3.2"
\echo "Ensure that all Namespaces have Network Policies defined (Manual)"
\set check_id '5.3.2.1'
\echo "Executing check 5.3.2.1"
\ir ../queries/network_hardening/network_policy_default_deny_ingress.sql
\set check_id '5.3.2.2'
\echo "Executing check 5.3.2.2"
\ir ../queries/network_hardening/network_policy_default_deny_egress.sql



\echo "5.4 Secrets Management"
\set check_id '5.4.1' 
\echo "Executing check 5.4.1"
\echo "Prefer using secrets as files over secrets as environment variables (Manual)"
\ir ../queries/manual.sql 

\set check_id '5.4.2' 
\echo "Executing check 5.4.2"
\echo "Consider external secret storage (Manual)"
\ir ../queries/manual.sql

\echo "5.5 Extensible Admission Control"
\set check_id '5.5.1'
\echo "Executing check 5.5.1"
\echo "Configure Image Provenance using ImagePolicyWebhook admission controller (Manual)"
\ir ../queries/manual.sql

-- for some reason, there is no section 5.6 in the CIS policy

\echo "5.7 General Policies"
\set check_id '5.7.1'
\echo "Executing check 5.7.1"
\echo "Create administrative boundaries between resources using namespaces (Manual)"
\ir ../queries/manual.sql

\set check_id '5.7.2'
\echo "Executing check 5.7.2"
\echo "Ensure that the seccomp profile is set to docker/default in your pod definitions (Manual)"
\set check_id '5.7.2.1'
\echo "Executing check 5.7.2.1"
\ir ../queries/pod_security/statefulset_default_seccomp_profile_enabled.sql
\set check_id '5.7.2.2'
\echo "Executing check 5.7.2.2"
\ir ../queries/pod_security/replicaset_default_seccomp_profile_enabled.sql
\set check_id '5.7.2.3'
\echo "Executing check 5.7.2.3"
\ir ../queries/pod_security/pod_container_default_seccomp_profile_enabled.sql
\set check_id '5.7.2.4'
\echo "Executing check 5.7.2.4"
\ir ../queries/pod_security/job_container_default_seccomp_profile_enabled.sql
\set check_id '5.7.2.5'
\echo "Executing check 5.7.2.5"
\ir ../queries/pod_security/deployment_container_default_seccomp_profile_enabled.sql
\set check_id '5.7.2.6'
\echo "Executing check 5.7.2.6"
\ir ../queries/pod_security/daemonset_default_seccomp_profile_enabled.sql


\set check_id '5.7.3'
\echo "Executing check 5.7.3"
\echo "Apply Security Context to Your Pods and Containers (Manual)"
\ir ../queries/manual.sql

\set check_id '5.7.4'
\echo "Executing check 5.7.4"
\echo "The default namespace should not be used (Manual)"
\set check_id '5.7.4.1'
\echo "Executing check 5.7.4.1"
\ir ../queries/pod_security/daemonset_default_namespace.sql
\set check_id '5.7.4.2'
\echo "Executing check 5.7.4.2"
\ir ../queries/pod_security/deployment_container_uses_default_namespace.sql
\set check_id '5.7.4.3'
\echo "Executing check 5.7.4.3"
\ir ../queries/pod_security/pod_container_uses_default_namespace.sql
\set check_id '5.7.4.4'
\echo "Executing check 5.7.4.4"
\ir ../queries/pod_security/job_container_uses_default_namespace.sql
\set check_id '5.7.4.5'
\echo "Executing check 5.7.4.5"
\ir ../queries/pod_security/replicaset_uses_default_namespace.sql
\set check_id '5.7.4.6'
\echo "Executing check 5.7.4.6"
\ir ../queries/pod_security/statefulset_uses_default_namespace.sql