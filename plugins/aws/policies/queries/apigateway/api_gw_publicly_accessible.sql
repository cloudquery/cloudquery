insert into aws_policy_results
select
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Find all API Gateway instances that are publicly accessible' AS title,
    account_id,
    arn as resource_id,
    case
        when NOT '{PRIVATE}' = endpoint_configuration_types then 'fail'
        else 'pass'
        end as status
from
    aws_apigateway_rest_apis
