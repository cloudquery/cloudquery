\echo "Executing Business Continuity and Risk Assessment"
\echo "Check 1634.12b1Organizational.1 - 12.b - 1"
\set check_id "1634.12b1Organizational.1 - 12.b - 1"
\i queries/compute/audit_virtual_machines_without_disaster_recovery_configured.sql

\set check_id "1635.12b1Organizational.2 - 12.b - 1"
\echo :check_id
\i queries/keyvault/azure_key_vault_managed_hsm_should_have_purge_protection_enabled.sql
\set check_id "1635.12b1Organizational.2 - 12.b - 2"
\echo :check_id
\i queries/keyvault/not_recoverable.sql

\set check_id "1638.12b2Organizational.345 - 12.b - 1"
\echo :check_id
\i queries/compute/audit_virtual_machines_without_disaster_recovery_configured.sql
