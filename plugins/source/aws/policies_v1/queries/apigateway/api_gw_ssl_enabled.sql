insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'API Gateway REST API stages should be configured to use SSL certificates for backend authentication' as title,
  account_id,
  arn as resource_id,
  case
    when cert is null then 'fail'
    else 'pass'
  end as status
from
    view_aws_apigateway_method_settings
