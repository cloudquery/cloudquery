insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'IAM authentication should be configured for RDS clusters' as title,
    account_id,
    arn AS resource_id,
    case when iam_database_authentication_enabled is not TRUE then 'fail' else 'pass' end as status
from aws_rds_clusters
