insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Lightsail instances should use IMDSv2' as title,
  account_id,
  arn as resource_id,
  case when
    metadata_options->>'HttpTokens' is distinct from 'required'
    then 'fail'
    else 'pass'
  end as status
from aws_lightsail_instances
