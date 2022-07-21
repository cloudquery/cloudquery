insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'security group is not currently in use so it should be deleted' as title,
    account_id,
    arn as resource_id,
    case when aws_ec2_instance_network_interface_groups.cq_id is null then 'fail' else 'pass' end as status
from
    aws_ec2_security_groups
left join
    aws_ec2_instance_network_interface_groups on
        aws_ec2_security_groups.id = aws_ec2_instance_network_interface_groups.group_id
