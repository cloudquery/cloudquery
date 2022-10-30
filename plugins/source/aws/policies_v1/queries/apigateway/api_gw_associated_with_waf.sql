insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'API Gateway should be associated with an AWS WAF web ACL' AS title,
  account_id,
  arn as resource_id,
  case
    when waf is null then 'fail'
    else 'pass'
  end as status
from
    view_aws_apigateway_method_settings
