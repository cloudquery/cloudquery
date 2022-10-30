\set ON_ERROR_STOP on
SET TIME ZONE 'UTC';
-- trick to set execution_time if not already set
-- https://stackoverflow.com/questions/32582600/only-set-variable-in-psql-script-if-not-specified-on-the-command-line
\set execution_time :execution_time
SELECT CASE 
  WHEN :'execution_time' = ':execution_time' THEN to_char(now(), 'YYYY-MM-dd HH24:MI:SS.US')
  ELSE :'execution_time'
END AS "execution_time"  \gset
\set framework 'hipaa_hitrust_v9.2'
\echo "Creating azure_policy_results table if not exist"
\ir ../create_azure_policy_results.sql
-- \ir privilege_management.sql
\ir user_authentication_for_external_connections.sql
\ir segregation_in_networks.sql
\ir network_connection_control.sql
\ir user_identification_and_authentication.sql
\ir identification_of_risks_related_to_external_parties.sql
\ir audit_logging.sql
\ir monitoring_system_use.sql
\ir administrator_and_operator_logs.sql
\ir segregation_of_duties.sql
\ir controls_against_malicious_code.sql
\ir back_up.sql
\ir network_controls.sql
\ir security_of_network_services.sql
\ir management_of_removable_media.sql
\ir information_exchange_policies_and_procedures.sql
\ir on_line_transactions.sql
-- \ir control_of_operational_software.sql
-- \ir change_control_procedures.sql
\ir control_of_technical_vulnerabilities.sql
\ir business_continuity_and_risk_assessment.sql
