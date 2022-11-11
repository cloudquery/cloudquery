\set ON_ERROR_STOP on
SET TIME ZONE 'UTC';
\set framework 'imds_v2'
-- trick to set execution_time if not already set
-- https://stackoverflow.com/questions/32582600/only-set-variable-in-psql-script-if-not-specified-on-the-command-line
\set execution_time :execution_time
SELECT CASE 
  WHEN :'execution_time' = ':execution_time' THEN to_char(now(), 'YYYY-MM-dd HH24:MI:SS.US')
  ELSE :'execution_time'
END AS "execution_time"  \gset
\ir ../create_aws_policy_results.sql

\set check_id 'EC2-IMDSv2'
\echo "Executing check EC2 IMDSv2"
\ir ../queries/ec2/not_imdsv2_instances.sql

\set check_id 'Lightsail-IMDSv2'
\echo "Executing check Lightsail IMDSv2"
\ir ../queries/lightsail/not_imdsv2_instances.sql

\set check_id 'AMIs-IMDSv2'
\echo "Executing check EC2 AMI IMDSv2"
\ir ../queries/ec2/images_imdsv2_required.sql