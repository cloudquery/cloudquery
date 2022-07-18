\echo "Executing Control of Technical Vulnerabilities"
\set check_id "0709.10m1Organizational.1 - 10.m - 1"
\echo :check_id
\i queries/compute/machines_without_vulnerability_assessment_extension.sql
\set check_id "0709.10m1Organizational.1 - 10.m - 2"
\echo :check_id
\i queries/sql/sql_databases_with_unresolved_vulnerability_findings.sql
\set check_id "0709.10m1Organizational.1 - 10.m - 6"
\echo :check_id
\i queries/sql/managed_instances_without_vulnerability_assessments.sql
\set check_id "0709.10m1Organizational.1 - 10.m - 7"
\echo :check_id
\i queries/sql/servers_without_vulnerability_assessments.sql
\set check_id "0710.10m2Organizational.1 - 10.m"
\echo :check_id
\i queries/sql/managed_instances_without_vulnerability_assessments.sql
\set check_id "0711.10m2Organizational.23 - 10.m "
\echo :check_id
\i queries/compute/machines_without_vulnerability_assessment_extension.sql
\set check_id "0716.10m3Organizational.1 - 10.m - 1"
\echo :check_id
\i queries/sql/sql_databases_with_unresolved_vulnerability_findings.sql
\set check_id "0719.10m3Organizational.5 - 10.m - 1"
\echo :check_id
\i queries/sql/managed_instances_without_vulnerability_assessments.sql
