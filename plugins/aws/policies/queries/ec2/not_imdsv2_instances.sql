insert into aws_policy_results
select
  :execution_time as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'EC2 instances should use IMDSv2' as title,
  account_id,
  id as resource_id,
  case when
    metadata_options_http_tokens is distinct from 'required'
    then 'fail'
    else 'pass'
  end as status
from aws_ec2_instances
