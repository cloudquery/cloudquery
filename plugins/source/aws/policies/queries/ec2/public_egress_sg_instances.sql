insert into aws_policy_results
-- Find all AWS instances that have a security group that allows unrestricted egress
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'All ec2 instances that have unrestricted access to the internet via a security group' as title,
    aws_ec2_instances.account_id,
    instance_id as resource_id,
    'fail' as status -- TODO FIXME
from aws_ec2_instances, jsonb_array_elements(security_groups) sg
    -- 	Find all instances that have egress rule that allows access to all ip addresses
    inner join view_aws_security_group_egress_rules on id = sg->>'GroupId'
where (ip = '0.0.0.0/0' or ip6 = '::/0')
