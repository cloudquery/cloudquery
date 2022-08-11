insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Amazon Redshift clusters should have audit logging enabled' as title,
    account_id,
    arn as resource_id,
    case when
     jsonb_typeof(logging_status -> 'LoggingEnabled') is null
     or (
             jsonb_typeof(logging_status -> 'LoggingEnabled') is not null
             and (logging_status ->> 'LoggingEnabled')::BOOLEAN is FALSE
         )
    then 'fail' else 'pass' end as status
from aws_redshift_clusters
