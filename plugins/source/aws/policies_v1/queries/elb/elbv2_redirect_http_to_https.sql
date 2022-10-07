insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Application Load Balancer should be configured to redirect all HTTP requests to HTTPS' as title,
  account_id,
  arn as resource_id,
  case when
   protocol = 'HTTP' and (
        da->>'Type' != 'REDIRECT' or da->'RedirectConfig'->>'Protocol' != 'HTTPS')
    then 'fail'
    else 'pass'
  end as status
from aws_elbv2_listeners, JSONB_ARRAY_ELEMENTS(default_actions) AS da
