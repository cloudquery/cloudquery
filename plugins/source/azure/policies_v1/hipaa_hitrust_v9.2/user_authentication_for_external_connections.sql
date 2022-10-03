\echo "User Authentication for External Connections"

\set check_id '1116_01j1organizational_145_01_j'
\echo "MFA should be enabled on accounts with owner permissions on your subscription"
\ir ../queries/security/mfa_should_be_enabled_on_accounts_with_owner_permissions_on_your_subscription.sql


\set check_id '1117_01j1organizational_23_01_j'
\echo "MFA should be enabled accounts with write permissions on your subscription"
\ir ../queries/security/mfa_should_be_enabled_accounts_with_write_permissions_on_your_subscription.sql


\set check_id '1118_01j2organizational_124_01_j'
\echo "MFA should be enabled on accounts with read permissions on your subscription"
\ir ../queries/security/mfa_should_be_enabled_on_accounts_with_read_permissions_on_your_subscription.sql


\set check_id '1119_01j2organizational_3_01_j'
\echo "Management ports of virtual machines should be protected with just-in-time network access control"
\ir ../queries/compute/virtual_machines_without_jit_network_access_policy.sql


\set check_id '1121_01j3organizational_2_01_j'
\echo "MFA should be enabled on accounts with owner permissions on your subscription"
\ir ../queries/security/mfa_should_be_enabled_on_accounts_with_owner_permissions_on_your_subscription.sql


\set check_id '1173_01j1organizational_6_01_j'
\echo "MFA should be enabled accounts with write permissions on your subscription"
\ir ../queries/security/mfa_should_be_enabled_accounts_with_write_permissions_on_your_subscription.sql


\set check_id '1174_01j1organizational_7_01_j'
\echo "MFA should be enabled on accounts with read permissions on your subscription"
\ir ../queries/security/mfa_should_be_enabled_on_accounts_with_read_permissions_on_your_subscription.sql


\set check_id '1175_01j1organizational_8_01_j'
\echo "Management ports of virtual machines should be protected with just-in-time network access control"
\ir ../queries/compute/virtual_machines_without_jit_network_access_policy.sql


\set check_id '1176_01j2organizational_5_01_j'
\echo "MFA should be enabled on accounts with owner permissions on your subscription"
\ir ../queries/security/mfa_should_be_enabled_on_accounts_with_owner_permissions_on_your_subscription.sql


\set check_id '1177_01j2organizational_6_01_j'
\echo "MFA should be enabled accounts with write permissions on your subscription"
\ir ../queries/security/mfa_should_be_enabled_accounts_with_write_permissions_on_your_subscription.sql


\set check_id '1178_01j2organizational_7_01_j'
\echo "MFA should be enabled on accounts with read permissions on your subscription"
\ir ../queries/security/mfa_should_be_enabled_on_accounts_with_read_permissions_on_your_subscription.sql


\set check_id '1179_01j3organizational_1_01_j'
\echo "Management ports of virtual machines should be protected with just-in-time network access control"
\ir ../queries/compute/virtual_machines_without_jit_network_access_policy.sql
      
    
  