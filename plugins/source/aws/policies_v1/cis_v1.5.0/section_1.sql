\echo  "Executing CIS V1.5.0 Section 1"
\set check_id '1.1'
    -- manual
\set check_id '1.2'
    -- manual
\set check_id '1.3'
    -- manual
\set check_id '1.4'
\echo "Executing check 1.4"
\ir ../queries/iam/root_user_no_access_keys.sql
\set check_id '1.5'
\echo "Executing check 1.5"
\ir ../queries/iam/mfa_enabled_for_root.sql
\set check_id '1.6'
\echo "Executing check 1.6"
\ir ../queries/iam/hardware_mfa_enabled_for_root.sql
\set check_id '1.7'
    -- todo credential report add password_last_used, access_key_1_last_used_date, access_key_2_last_used_date
\set check_id '1.8'
\echo "Executing check 1.8"
\ir ../queries/iam/password_policy_min_length.sql
\set check_id '1.9'
\echo "Executing check 1.9"
\ir ../queries/iam/password_policy_prevent_reuse.sql
\set check_id '1.10'
\echo "Executing check 1.10"
\ir ../queries/iam/mfa_enabled_for_console_access.sql
\set check_id '1.11'
    -- todo credential report  add access_key_1_last_used_date,access_key_2_last_used_date
\set check_id '1.12'
\echo "Executing check 1.12"
\ir ../queries/iam/unused_creds_disabled_45_days.sql
\set check_id '1.13'
\echo "Executing check 1.13"
\ir ../queries/iam/users_with_two_active_access_keys.sql
\set check_id '1.14'
\echo "Executing check 1.14"
\ir ../queries/iam/old_access_keys.sql
\set check_id '1.15'
\echo "Executing check 1.15"
\ir ../queries/iam/policies_attached_to_groups_roles.sql
\set check_id '1.16'
    -- todo svc.ListPolicies is not used (implement it and then do a check)
\set check_id '1.17'
    -- todo svc.ListPolicies is not used (implement it and then do a check)
\set check_id '1.18'
    -- manual
\set check_id '1.19'
\echo "Executing check 1.19"
\ir ../queries/iam/server_certificates_expired.sql
\set check_id '1.20'
\echo "Executing check 1.20"
-- \ir ../queries/accessanalyzer/regions_with_no_accessanalyzers.sql
\set check_id '1.21'
    -- manual
