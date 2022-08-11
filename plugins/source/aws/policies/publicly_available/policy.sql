\set ON_ERROR_STOP on
SET TIME ZONE 'UTC';
\set framework 'publicly_available'
-- trick to set execution_time if not already set
-- https://stackoverflow.com/questions/32582600/only-set-variable-in-psql-script-if-not-specified-on-the-command-line
\set execution_time :execution_time
SELECT CASE 
  WHEN :'execution_time' = ':execution_time' THEN to_char(now(), 'YYYY-MM-dd HH24:MI:SS.US')
  ELSE :'execution_time'
END AS "execution_time"  \gset
\ir ../create_aws_policy_results.sql

\set check_id 'API-Gateways'
\echo "Executing check API-Gateways"
\ir ../queries/apigateway/api_gw_publicly_accessible.sql

\set check_id 'API-Gateway-V2'
\echo "Executing check API-Gateway-V2"
\ir ../queries/apigateway/api_gw_v2_publicly_accessible.sql

\set check_id 'CloudFront-Distributions'
\echo "Executing check CloudFront-Distributions"
\ir ../queries/cloudfront/all_distributions.sql

\set check_id 'EC2-Public-Ips'
\echo "Executing check EC2-Public-Ips"
\ir ../queries/ec2/public_ips.sql

\set check_id 'ELB-Classic'
\echo "Executing check ELB-Classic"
\ir ../queries/elb/elbv1_internet_facing.sql

\set check_id 'ELB-V2'
\echo "Executing check ELB-V2"
\ir ../queries/elb/elbv2_internet_facing.sql

\set check_id 'Redshift'
\echo "Executing check Redshift"
\ir ../queries/redshift/cluster_publicly_accessible.sql

\set check_id 'RDS'
\echo "Executing check RDS"
\ir ../queries/rds/rds_db_instances_should_prohibit_public_access.sql
