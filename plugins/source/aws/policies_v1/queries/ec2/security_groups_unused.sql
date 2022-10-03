insert into aws_policy_results
with interface_groups as (
    select distinct a->>'GroupId' as group_id from aws_ec2_instances, jsonb_array_elements(security_groups) as a
)
select :'execution_time'           as execution_time,
       :'framework'                as framework,
       :'check_id'                 as check_id,
       'Unused EC2 security group' as title,
       sg.account_id,
       sg.arn                      as resource_id,
       'fail'                      as status
from aws_ec2_security_groups sg
         left join interface_groups on interface_groups.group_id = sg.group_id
where interface_groups.group_id is null;
