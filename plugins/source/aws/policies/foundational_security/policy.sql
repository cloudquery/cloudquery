\set ON_ERROR_STOP on
SET TIME ZONE 'UTC';
-- trick to set execution_time if not already set
-- https://stackoverflow.com/questions/32582600/only-set-variable-in-psql-script-if-not-specified-on-the-command-line
\set execution_time :execution_time
SELECT CASE 
  WHEN :'execution_time' = ':execution_time' THEN to_char(now(), 'YYYY-MM-dd HH24:MI:SS.US')
  ELSE :'execution_time'
END AS "execution_time"  \gset
\set framework 'foundational_security'
\ir ../create_aws_policy_results.sql
\ir ./account.sql
\ir ./acm.sql
\ir ./apigateway.sql
\ir ./autoscaling.sql
\ir ./awsconfig.sql
\ir ./cloudfront.sql
\ir ./cloudtrail.sql
\ir ./codebuild.sql
\ir ./dms.sql
\ir ./dynamodb.sql
\ir ./ec2.sql
\ir ./ecs.sql
\ir ./efs.sql
\ir ./elastic_beanstalk.sql
\ir ./elasticsearch.sql
\ir ./elb.sql
\ir ./elbv2.sql
\ir ./emr.sql
\ir ./guardduty.sql
\ir ./iam.sql
\ir ./kms.sql
\ir ./lambda.sql
\ir ./rds.sql
\ir ./redshift.sql
\ir ./s3.sql
\ir ./sagemaker.sql
\ir ./secretmanager.sql
\ir ./sns.sql
\ir ./sqs.sql
\ir ./ssm.sql
\ir ./waf.sql
