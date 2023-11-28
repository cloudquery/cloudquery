\set ON_ERROR_STOP on
SET TIME ZONE 'UTC';
-- neat trick to set execution_time if not already set
-- https://stackoverflow.com/questions/32582600/only-set-variable-in-psql-script-if-not-specified-on-the-command-line
\set execution_time :execution_time
SELECT CASE 
  WHEN :'execution_time' = ':execution_time' THEN to_char(now(), 'YYYY-MM-dd HH24:MI:SS.US')
  ELSE :'execution_time'
END AS "execution_time"  \gset
\ir cis_v1.2.0/policy.sql
\ir pci_dss_v3.2.1/policy.sql
\ir foundational_security/policy.sql
\ir public_egress/policy.sql
\ir publicly_available/policy.sql
\ir imds_v2/policy.sql