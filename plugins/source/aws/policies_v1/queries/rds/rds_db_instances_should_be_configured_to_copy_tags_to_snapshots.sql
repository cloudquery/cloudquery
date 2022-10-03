insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'RDS DB instances should be configured to copy tags to snapshots' as title,
    account_id,
    arn AS resource_id,
    case when copy_tags_to_snapshot is not TRUE then 'fail' else 'pass' end as status
from aws_rds_instances
