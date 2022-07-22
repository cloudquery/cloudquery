\set ON_ERROR_STOP on
SET TIME ZONE 'UTC';
-- neat trick to set execution_time if not already set
-- https://stackoverflow.com/questions/32582600/only-set-variable-in-psql-script-if-not-specified-on-the-command-line
\set execution_time :execution_time
SELECT CASE 
  WHEN :'execution_time' = ':execution_time' THEN to_char(now(), 'YYYY-MM-dd HH24:MI:SS.US')
  ELSE :'execution_time'
END AS "execution_time"  \gset

\set framework 'cis_v1.2.0'

\ir ../create_k8s_policy_results.sql
\ir ./network_hardening.sql
\ir ./pod_security.sql
