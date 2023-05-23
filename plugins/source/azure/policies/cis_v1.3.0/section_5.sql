\echo "Executing CIS V1.3.0 Section 5"
\set check_id '5.1.1'
\echo "Executing check 5.1.1"
\ir ../queries/monitor/no_diagnostic_setting.sql
\set check_id '5.1.2'
\echo "Executing check 5.1.2"
\ir ../queries/monitor/insufficient_diagnostic_capturing_settings.sql
\set check_id '5.1.3'
\echo "Executing check 5.1.3"
\ir ../queries/storage/no_publicly_accessible_insights_activity_logs.sql
\set check_id '5.1.4'
\echo "Executing check 5.1.4"
\ir ../queries/storage/encrypt_with_cmk.sql
\set check_id '5.1.5'
\echo "Executing check 5.1.5"
\ir ../queries/monitor/logging_key_calut_is_enabled.sql
\set check_id '5.2.1'
\echo "Executing check 5.2.1"
\ir ../queries/monitor/log_alert_for_create_policy_assignment.sql
\set check_id '5.2.2'
\echo "Executing check 5.2.2"

\set check_id '5.2.3'
\echo "Executing check 5.2.3"

\set check_id '5.2.4'
\echo "Executing check 5.2.4"

\set check_id '5.2.5'
\echo "Executing check 5.2.5"

\set check_id '5.2.6'
\echo "Executing check 5.2.6"

\set check_id '5.2.7'
\echo "Executing check 5.2.7"

\set check_id '5.2.8'
\echo "Executing check 5.2.8"

\set check_id '5.2.9'
\echo "Executing check 5.2.9"

\set check_id '5.3'
\echo "Executing check 5.3"
