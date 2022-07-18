\echo  "Executing CIS V1.2.0 Section 1"
\set check_id "1.1"
\echo "Executing check 1.1"
\i queries/iam/avoid_root_usage.sql
\set check_id "1.2"
\echo "Executing check 1.2"
\i queries/iam/mfa_enabled_for_console_access.sql
\set check_id "1.3"
\echo "Executing check 1.3"
\i queries/iam/unused_creds_disabled.sql
\set check_id "1.4"
\echo "Executing check 1.4"
\i queries/iam/old_access_keys.sql
\set check_id "1.5"
\echo "Executing check 1.5"
\i queries/iam/password_policy_min_uppercase.sql
\set check_id "1.6"
\echo "Executing check 1.6"
\i queries/iam/password_policy_min_lowercase.sql
\set check_id "1.7"
\echo "Executing check 1.7"
\i queries/iam/password_policy_min_one_symbol.sql
\set check_id "1.8"
\echo "Executing check 1.8"
\i queries/iam/password_policy_min_number.sql
\set check_id "1.9"
\echo "Executing check 1.9"
\i queries/iam/password_policy_min_length.sql
\set check_id "1.10"
\echo "Executing check 1.10"
\i queries/iam/password_policy_prevent_reuse.sql
\set check_id "1.11"
\echo "Executing check 1.11"
\i queries/iam/password_policy_expire_old_passwords.sql
\set check_id "1.12"
\echo "Executing check 1.12"
\i queries/iam/root_user_no_access_keys.sql
\set check_id "1.13"
\echo "Executing check 1.13"
\i queries/iam/mfa_enabled_for_root.sql
\set check_id "1.14"
\echo "Executing check 1.14"
\i queries/iam/hardware_mfa_enabled_for_root.sql
\set check_id "1.16"
\echo "Executing check 1.16"
\i queries/iam/policies_attached_to_groups_roles.sql
