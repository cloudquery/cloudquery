\set ON_ERROR_STOP on

SET TIME ZONE 'UTC';
-- trick to set execution_time if not already set
-- https://stackoverflow.com/questions/32582600/only-set-variable-in-psql-script-if-not-specified-on-the-command-line
\set execution_time '2022-07-20 16:16:30.298982+00'
SELECT CASE 
  WHEN :'execution_time' = '2022-07-20 16:16:30.298982+00'
  THEN now()
  ELSE :'execution_time'
END AS "execution_time"  \gset

\ir nsa_cisa_v1/policy.sql