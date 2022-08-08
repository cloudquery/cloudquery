insert into aws_policy_results
with data as (
    select account_id, id, COUNT(aws_ec2_instance_network_interfaces.cq_id) as cnt
    from aws_ec2_instances
left join
    aws_ec2_instance_network_interfaces on
        aws_ec2_instances.cq_id = aws_ec2_instance_network_interfaces.instance_cq_id
group by account_id,
    region,
    id
)
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'EC2 instances should not use multiple ENIs' as title,
    account_id,
    id as resource_id,
    case when cnt > 1 then 'fail' else 'pass' end as status
from data
