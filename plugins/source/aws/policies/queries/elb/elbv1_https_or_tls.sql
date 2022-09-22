insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Classic Load Balancer listeners should be configured with HTTPS or TLS termination' as title,
  lb.account_id,
  lb.arn as resource_id,
  case when
    li->'Listener'->>'Protocol' not in ('HTTPS', 'SSL')
    then 'fail'
    else 'pass'
  end as status
from aws_elbv1_load_balancers lb, jsonb_array_elements(lb.listener_descriptions) as li
