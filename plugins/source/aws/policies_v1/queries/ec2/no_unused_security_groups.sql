insert into aws_policy_results
with interface_groups as (
    select distinct g->>'GroupId' as id from aws_ec2_instances i, jsonb_array_elements(network_interfaces) as a, jsonb_array_elements(a->'Groups') as g
)
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'security group is not currently in use so it should be deleted' as title,
    account_id,
    arn as resource_id,
    case when interface_groups.id is null then 'fail' else 'pass' end as status
from
    aws_ec2_security_groups
left join
    interface_groups on
        aws_ec2_security_groups.group_id = interface_groups.id;
