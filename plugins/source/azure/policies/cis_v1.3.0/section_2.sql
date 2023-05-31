\echo "Executing CIS V1.3.0 Section 2"
\set check_id "2.1"
\echo "Executing check 2.1"
\ir ../queries/security/defender_on_for_servers.sql
\set check_id "2.2"
\echo "Executing check 2.2"
\ir ../queries/security/defender_on_for_app_service.sql
\set check_id "2.3"
\echo "Executing check 2.3"
\ir ../queries/security/defender_on_for_sql_servers.sql
\set check_id "2.4"
\echo "Executing check 2.4"
\ir ../queries/security/defender_on_for_sql_servers_on_machines.sql
\set check_id "2.5"
\echo "Executing check 2.5"
\ir ../queries/security/defender_on_for_storage.sql
\set check_id "2.6"
\echo "Executing check 2.6"
\ir ../queries/security/defender_on_for_k8s.sql
\set check_id "2.7"
\echo "Executing check 2.7"
\ir ../queries/security/defender_on_for_container_registeries.sql
\set check_id "2.8"
\echo "Executing check 2.8"
\ir ../queries/security/defender_on_for_key_vault.sql
-- security settings does not have "enabled" property
-- \set check_id "2.10"
-- \echo "Executing check 2.10"
-- \ir ../queries/security/mcas_integration_with_security_center_enabled.sql
\set check_id "2.11"
\echo "Executing check 2.11"
\ir ../queries/security/auto_provisioning_monitoring_agent_enabled.sql
\set check_id "2.12"
\echo "Executing check 2.12"
\ir ../queries/security/default_policy_disabled.sql
-- security contacts api is broken
-- \set check_id "2.13"
-- \echo "Executing check 2.13"
-- \ir ../queries/security/security_email_configured.sql
-- \set check_id "2.14"
-- \echo "Executing check 2.14"
-- \ir ../queries/security/notify_high_severity_alerts.sql
