insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'API Gateway REST API cache data should be encrypted at rest' AS title,
  account_id,
  arn as resource_id,
  case
    when stage_caching_enabled is true
        or (
            caching_enabled is true
            and cache_data_encrypted is not true
        ) then 'pass'
    else 'fail'
  end as status
from
    view_aws_apigateway_method_settings
