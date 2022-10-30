insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'API Gateway REST API stages should have AWS X-Ray tracing enabled' as title,
    account_id,
    arn as resource_id,
    case
        when (stage_data_trace_enabled is not true or caching_enabled is not true) then 'fail'
        else 'pass'
        end as status
from
    view_aws_apigateway_method_settings
