insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'The VPC default security group should not allow inbound and outbound traffic' as title,
  account_id,
  arn,
  case when
      group_name='default' 
      AND (jsonb_array_length(ip_permissions) > 0
      OR jsonb_array_length(ip_permissions_egress) > 0)
      then 'fail'
      else 'pass'
  end
from
    aws_ec2_security_groups;
