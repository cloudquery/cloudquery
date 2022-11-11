insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Classic Load Balancers with HTTPS/SSL listeners should use a predefined security policy that has strong configuration' as title,
  lb.account_id,
  lb.arn as resource_id,
  case when
    li->'Listener'->>'Protocol' in ('HTTPS', 'SSL')
    and 'ELBSecurityPolicy-TLS-1-2-2017-01' != any( ARRAY(SELECT JSONB_ARRAY_ELEMENTS_TEXT(lb.policies->'OtherPolicies')) )
    then 'fail'
    else 'pass'
  end as status
from aws_elbv1_load_balancers lb, jsonb_array_elements(lb.listener_descriptions) as li