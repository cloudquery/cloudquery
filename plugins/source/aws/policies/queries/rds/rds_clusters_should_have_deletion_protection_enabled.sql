insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'RDS clusters should have deletion protection enabled' as title,
    account_id,
    arn AS resource_id,
    case when deletion_protection is not TRUE then 'fail' else 'pass' end as status
from aws_rds_clusters
