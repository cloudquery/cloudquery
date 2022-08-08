insert into aws_policy_results
-- Find all AWS instances that have a security group that allows unrestricted egress
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'All ec2 instances that have unrestricted access to the internet via a security group' as title,
    account_id,
    id as resource_id,
    'fail' as status -- TODO FIXME
from aws_ec2_instances
where cq_id in
    -- 	Find all instances that have egress rule that allows access to all ip addresses
    (select instance_cq_id
        from aws_ec2_instance_security_groups
        inner join view_aws_security_group_egress_rules on group_id = id
        where (ip = '0.0.0.0/0' or ip = '::/0'))
