SET TIME ZONE 'UTC';
-- trick to set execution_time if not already set
-- https://stackoverflow.com/questions/32582600/only-set-variable-in-psql-script-if-not-specified-on-the-command-line
\set execution_time :execution_time
SELECT CASE 
  WHEN :'execution_time' = ':execution_time' THEN to_char(now(), 'YYYY-MM-dd HH24:MI:SS.US')
  ELSE :'execution_time'
END AS "execution_time"  \gset
\set framework 'cis_v1.3.0'
\echo "Creating azure_policy_results table if not exist"
\ir ../create_azure_policy_results.sql
\echo "Creating view view_azure_security_policy_parameters"
\ir ../views/policy_assignment_parameters.sql

\ir section_1.sql
\ir section_2.sql
\ir section_3.sql
\ir section_4.sql
\ir section_5.sql
\ir section_6.sql
\ir section_7.sql
\ir section_8.sql
\ir section_9.sql
