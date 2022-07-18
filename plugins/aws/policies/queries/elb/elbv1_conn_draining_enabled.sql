insert into aws_policy_results
select
  :execution_time as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Classic Load Balancers should have connection draining enabled' as title,
  account_id,
  aws_elbv1_load_balancers.arn as resource_id,
  case when
    attributes_connection_draining_enabled is not true
    then 'fail'
    else 'pass'
  end as status
from
    aws_elbv1_load_balancers
