\echo "Executing Privilege Management Checks"
\set check_id '11180_01c3system_6_01_c'
\echo "Management ports of virtual machines should be protected with just-in-time network access control"
\ir ../queries/compute/virtual_machines_without_jit_network_access_policy.sql

\set check_id '1143_01c1system_123_01_c'
\echo "Management ports should be closed on your virtual machines"
\ir ../queries/network/security_groups_with_open_management_ports.sql

\set check_id '1144_01c1system_4_01_c'
\echo "A maximum of 3 owners should be designated for your subscription"
\ir ../queries/authorization/subscriptions_with_more_than_3_owners.sql

\set check_id '1145_01c2system_1_01_c'
\echo "There should be more than one owner assigned to your subscription"
\ir ../queries/authorization/subscriptions_with_less_than_2_owners.sql

\set check_id '1146_01c2system_23_01_c'
\echo "External accounts with owner permissions should be removed from your subscription"
\ir ../queries/security/external_accounts_with_owner_permissions_should_be_removed_from_your_subscription.sql

\set check_id '1147_01c2system_456_01_c'
\echo "Deprecated accounts with owner permissions should be removed from your subscription"
\ir ../queries/security/deprecated_accounts_with_owner_permissions_should_be_removed_from_your_subscription.sql

\set check_id '1148_01c2system_78_01_c_1'
\echo "Audit usage of custom RBAC rules"
\ir ../queries/authorization/custom_roles.sql

\set check_id '1148_01c2system_78_01_c_2'
\echo "Windows machines should meet requirements for 'Security Options_Accounts'"
\echo "Manual check"

\set check_id '1149_01c2system_9_01_c'
\echo "Role-Based Access Control (RBAC) should be used on Kubernetes Services"
\ir ../queries/container/aks_rbac_disabled.sql

\set check_id '1150_01c2system_10_01_c'
\echo "Management ports should be closed on your virtual machines"
\ir ../queries/network/security_groups_with_open_management_ports.sql

\set check_id '1151_01c3system_1_01_c'
\echo "A maximum of 3 owners should be designated for your subscription"
\ir ../queries/authorization/subscriptions_with_more_than_3_owners.sql

\set check_id '1152_01c3system_2_01_c'
\echo "There should be more than one owner assigned to your subscription"
\ir ../queries/authorization/subscriptions_with_less_than_2_owners.sql

\set check_id '1153_01c3system_35_01_c'
\echo "Role-Based Access Control (RBAC) should be used on Kubernetes Services"
\ir ../queries/container/aks_rbac_disabled.sql

\set check_id '1154_01c3system_4_01_c'
\echo "Contractors are provided with minimal system and physical access only after the organization assesses the contractor's ability to comply with its security requirements and the contractor agrees to comply."
\ir ../queries/authorization/subscriptions_with_more_than_3_owners.sql
