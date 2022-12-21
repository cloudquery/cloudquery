SET TIME ZONE 'UTC';
-- trick to set execution_time if not already set
-- https://stackoverflow.com/questions/32582600/only-set-variable-in-psql-script-if-not-specified-on-the-command-line
\set execution_time :execution_time
SELECT CASE 
  WHEN :'execution_time' = ':execution_time' THEN to_char(now(), 'YYYY-MM-dd HH24:MI:SS.US')
  ELSE :'execution_time'
END AS "execution_time"  \gset
\set framework 'cis_v1.5.0'
\echo "Executing CIS Microsoft Azure Foundations Benchmark v1.5.0"

\echo "Creating azure_policy_results table if not exist"
\ir ../create_azure_policy_results.sql
\echo "Creating view view_azure_security_policy_parameters"
\ir ../views/policy_assignment_parameters.sql

-- Sections left to do and move into section.sqls
\echo "Executing CIS V1.5.0 Section 3: Storage Accounts"
\echo "Executing CIS V1.5.0 Section 4: Database Services"
\echo "Executing CIS V1.5.0 Section 5: Logging and Monitoring"
\echo "Executing CIS V1.5.0 Section 6: Networking"
\echo "Executing CIS V1.5.0 Section 7: Virtual Machines"
\echo "Executing CIS V1.5.0 Section 8: Key Vault"
\echo "Executing CIS V1.5.0 Section 9: AppService"
\echo "Executing CIS V1.5.0 Section 10: Miscellaneous"

-- CIS V1.5.0 Section 1: Identity and Access Management
\ir section_1.sql

-- CIS V1.5.0 Section 2: Microsoft Defender for Cloud
\ir section_2.sql

\ir section_3.sql
\ir section_4.sql
\ir section_5.sql
\ir section_6.sql
\ir section_7.sql
\ir section_8.sql
\ir section_9.sql
\ir section_10.sql