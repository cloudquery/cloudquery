insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Database logging should be enabled' as title,
    account_id,
    arn AS resource_id,
    case when
                 enabled_cloudwatch_logs_exports is null
                 or (engine in ('aurora', 'aurora-mysql', 'mariadb', 'mysql')
                 and not enabled_cloudwatch_logs_exports @> '{audit,error,general,slowquery}'
                     )
                 or (engine like '%postgres%'
                 and not enabled_cloudwatch_logs_exports @> '{postgresql,upgrade}')
                 or (engine like '%oracle%'
                 and not enabled_cloudwatch_logs_exports @> '{alert,audit,trace,listener}'
                     )
                 or (engine like '%sqlserver%'
                 and not enabled_cloudwatch_logs_exports @> '{error,agent}')
    then 'fail' else 'pass' end as status
from aws_rds_instances
