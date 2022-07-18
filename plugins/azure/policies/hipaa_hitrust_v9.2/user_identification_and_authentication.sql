\echo "User Identification and Authentication"
\set check_id "11109.01q1Organizational.57 - 01.q - 1"
\echo :check_id
\i queries/security/mfa_should_be_enabled_on_accounts_with_owner_permissions_on_your_subscription.sql
\set check_id "11110.01q1Organizational.6 - 01.q - 1"
\echo :check_id
\i queries/security/mfa_should_be_enabled_accounts_with_write_permissions_on_your_subscription.sql
\set check_id "11111.01q2System.4 - 01.q - 1"
\echo :check_id
\i queries/security/mfa_should_be_enabled_on_accounts_with_read_permissions_on_your_subscription.sql
\set check_id "11112.01q2Organizational.67 - 01.q - 1"
\echo :check_id
\i queries/authorization/subscriptions_with_more_than_3_owners.sql
\set check_id "11208.01q1Organizational.8 - 01.q - 1"
\echo :check_id
\i queries/authorization/subscriptions_with_less_than_2_owners.sql
