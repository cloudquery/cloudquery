\echo "Executing K8S RBAC and Service Accounts CIS v1.7.0"

\echo "Policies - RBAC and Service Accounts"

\echo "Ensure that the cluster-admin role is only used where required"
\set check_id "cluster_admin_role"
\ir ../queries_cis_v1_7_0/rbac_and_service_accounts/cluster_admin_role.sql

\echo "Minimize access to secrets"
\set check_id "minimize_access_to_secrets"
\ir ../queries_cis_v1_7_0/rbac_and_service_accounts/minimize_access_to_secrets.sql

\echo "Minimize wildcard use in Roles and ClusterRoles"
\set check_id "minimize_wildcard_use"
\ir ../queries_cis_v1_7_0/rbac_and_service_accounts/minimize_wildcard_use.sql


