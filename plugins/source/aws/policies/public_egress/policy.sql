\set ON_ERROR_STOP on
SET TIME ZONE 'UTC';
\set framework 'public_egress'
-- trick to set execution_time if not already set
-- https://stackoverflow.com/questions/32582600/only-set-variable-in-psql-script-if-not-specified-on-the-command-line
\set execution_time :execution_time
SELECT CASE 
  WHEN :'execution_time' = ':execution_time' THEN to_char(now(), 'YYYY-MM-dd HH24:MI:SS.US')
  ELSE :'execution_time'
END AS "execution_time"  \gset
\ir ../create_aws_policy_results.sql

\echo "Creating view_aws_security_group_egress_rules"
\ir ../views/security_group_egress_rules.sql

\set check_id 'ec2-all-instances-with-routes-and-security-groups'
\echo "Executing check ec2-all-instances-with-routes-and-security-groups"
\ir ../queries/ec2/public_egress_sg_and_routing_instances.sql

\set check_id 'ec2-instances'
\echo "Executing check ec2-instances"
\ir ../queries/ec2/public_egress_sg_instances.sql

\set check_id 'lambda-functions'
\echo "Executing check lambda-functions"
\ir ../queries/lambda/functions_with_public_egress.sql
