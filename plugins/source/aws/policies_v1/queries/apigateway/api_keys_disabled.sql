insert into aws_policy_results
select :'execution_time'            as execution_time,
       :'framework'                 as framework,
       :'check_id'                  as check_id,
       'Unused API Gateway API key' as title,
       account_id,
       arn                          as resource_id,
       'fail'                       as status
from aws_apigateway_api_keys
where enabled = false