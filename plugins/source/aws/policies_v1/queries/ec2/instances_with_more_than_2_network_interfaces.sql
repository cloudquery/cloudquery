insert into aws_policy_results
with data as (
    select account_id, instance_id, COUNT(nics->>'Status') as cnt
    from aws_ec2_instances left join jsonb_array_elements(aws_ec2_instances.network_interfaces) as nics on true
group by account_id, instance_id
)
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'EC2 instances should not use multiple ENIs' as title,
    account_id,
    instance_id as resource_id,
    case when cnt > 1 then 'fail' else 'pass' end as status
from data
