\set framework 'cis_eks_v1.1.0'
\echo "Creating CIS EKS V1.1.0 Section 4 Views"
\ir ../views/project_policy_members.sql
\echo "Executing CIS EKS V1.1.0 Section 4"


\echo "4 Policies"
\echo "4.1 RBAC and Service Accounts"

\set check_id '4.1.1'
\echo "Executing check 4.1.1"
\echo "Ensure that the cluster-admin role is only used where required (Manual)"
\ir ../queries/manual.sql

\set check_id '4.1.2'
\echo "Executing check 4.1.2"
\echo "Minimize access to secrets (Manual)"
\ir ../queries/manual.sql

\set check_id '4.1.3'
\echo "Executing check 4.1.3"
\echo "Minimize wildcard use in Roles and ClusterRoles (Manual)"
\ir ../queries/manual.sql

\set check_id '4.1.4'
\echo "Executing check 4.1.4"
\echo "Minimize access to create pods (Manual)"
\ir ../queries/manual.sql

\set check_id '4.1.5'
\echo "Executing check 4.1.5"
\echo "Ensure that default service accounts are not actively used. (Manual)"
\ir ../queries/manual.sql

\set check_id '4.1.6'
\echo "Executing check 4.1.6"
\echo "Ensure that Service Account Tokens are only mounted where necessary (Manual)"
\ir ../queries/manual.sql



\echo "4.2 Pod Security Policies"
\set check_id '4.2.1'
\echo "Executing check 4.2.1"
\echo "Minimize the admission of privileged containers (Automated)"
\ir ../queries/manual.sql

\set check_id '4.2.2'
\echo "Executing check 4.2.2"
\echo "Minimize the admission of containers wishing to share the host process ID namespace (Automated)"
\ir ../queries/manual.sql

\set check_id '4.2.3'
\echo "Executing check 4.2.3"
\echo "Minimize the admission of containers wishing to share the host IPC namespace (Automated)"
\ir ../queries/manual.sql

\set check_id '4.2.4'
\echo "Executing check 4.2.4"
\echo "Minimize the admission of containers wishing to share the host network namespace (Automated)"
\ir ../queries/manual.sql

\set check_id '4.2.5'
\echo "Executing check 4.2.5"
\echo "Minimize the admission of containers with allowPrivilegeEscalation (Automated)"
\ir ../queries/manual.sql

\set check_id '4.2.6'
\echo "Executing check 4.2.6"
\echo "Minimize the admission of root containers (Automated) "
\ir ../queries/manual.sql

\set check_id '4.2.7'
\echo "Executing check 4.2.7"
\echo "Minimize the admission of containers with the NET_RAW capability (Automated)"
\ir ../queries/manual.sql

\set check_id '4.2.8'
\echo "Executing check 4.2.8"
\echo "Minimize the admission of containers with added capabilities (Manual)"
\ir ../queries/manual.sql

\set check_id '4.2.9'
\echo "Executing check 4.2.9"
\echo "Minimize the admission of containers with capabilities assigned (Manual)"
\ir ../queries/manual.sql



\echo "4.3 CNI Plugin"
\set check_id '4.3.1'
\echo "Executing check 4.3.1"
\echo "Ensure latest CNI version is used (Manual)"
\ir ../queries/manual.sql

\set check_id '4.3.2'
\echo "Executing check 4.3.2"
\echo "Ensure that all Namespaces have Network Policies defined (Manual)"
\ir ../queries/manual.sql


\echo "4.4 Secrets Management"
\set check_id '4.4.1'
\echo "Executing check 4.4.1"
\echo "Prefer using secrets as files over secrets as environment variables (Manual)"
\ir ../queries/manual.sql

\set check_id '4.4.2'
\echo "Executing check 4.4.2"
\echo "Consider external secret storage (Manual)"
\ir ../queries/manual.sql



\echo "4.5 Extensible Admission Control"
\echo "There are no checks in this section"
-- another empty section /shrug

\echo "4.6 General Policies"
\set check_id '4.6.1'
\echo "Executing check 4.6.1"
\echo "Create administrative boundaries between resources using namespaces (Manual)"
\ir ../queries/manual.sql

\set check_id '4.6.2'
\echo "Executing check 4.6.2"
\echo "Apply Security Context to Your Pods and Containers (Manual)"
\ir ../queries/manual.sql

\set check_id '4.6.3'
\echo "Executing check 4.6.3"
\echo "The default namespace should not be used (Manual)"
\ir ../queries/manual.sql
