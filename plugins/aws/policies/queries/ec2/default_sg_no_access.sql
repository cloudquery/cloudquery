insert into aws_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'The VPC default security group should not allow inbound and outbound traffic',
  account_id,
  arn,
  case when
      group_name = 'default'
      then 'fail'
      else 'pass'
  end
from
    aws_ec2_security_groups
inner join
    aws_ec2_security_group_ip_permissions on
        aws_ec2_security_groups.cq_id
        = aws_ec2_security_group_ip_permissions.security_group_cq_id
