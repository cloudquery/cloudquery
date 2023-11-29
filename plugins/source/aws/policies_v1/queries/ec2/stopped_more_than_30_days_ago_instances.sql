insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Stopped EC2 instances should be removed after a specified time period' as title,
  account_id,
  instance_id as resource_id,
  case when
    state->>'Name' = 'stopped'
        AND NOW() - state_transition_reason_time > INTERVAL '30' DAY
    then 'fail'
    else 'pass'
  end
from aws_ec2_instances
