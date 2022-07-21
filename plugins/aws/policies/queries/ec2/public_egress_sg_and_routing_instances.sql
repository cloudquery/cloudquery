insert into aws_policy_results
-- Find all AWS instances that are in a subnet that includes a catchall route
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Find all ec2 instances that have unrestricted access to the internet with a wide open security group and routing' as title,
    account_id,
    id as resource_id,
    'fail' as status -- TODO FIXME
from aws_ec2_instances
where subnet_id in
    --  Find all subnets that include a route table that inclues a catchall route
    (select subnet_id
        from aws_ec2_route_tables
        inner join
            aws_ec2_route_table_associations on
                aws_ec2_route_table_associations.route_table_cq_id = aws_ec2_route_tables.cq_id
        where aws_ec2_route_tables.cq_id in
            --  Find all routes in any route table that contains a route to 0.0.0.0/0 or ::/0
            (select route_table_cq_id
                from aws_ec2_route_table_routes
                where destination_cidr_block = '0.0.0.0/0'
                    or destination_ipv6_cidr_block = '::/0'))
    and cq_id in
    -- 	Find all instances that have egress rule that allows access to all ip addresses
    (select instance_cq_id
        from aws_ec2_instance_security_groups
        inner join view_aws_security_group_egress_rules on group_id = id
        where (ip = '0.0.0.0/0' or ip = '::/0'))
