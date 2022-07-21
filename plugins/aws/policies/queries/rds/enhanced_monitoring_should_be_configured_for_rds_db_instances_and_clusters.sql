insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Enhanced monitoring should be configured for RDS DB instances and clusters' as title,
    account_id,
    arn AS resource_id,
    case when enhanced_monitoring_resource_arn is null then 'fail' else 'pass' end as status
from aws_rds_instances
