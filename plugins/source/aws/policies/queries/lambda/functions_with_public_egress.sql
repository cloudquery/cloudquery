insert into aws_policy_results
select distinct
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Find all ec2 instances that have unrestricted access to the internet' AS title,
    account_id,
    arn AS resource_id,
    'fail' AS status -- TODO FIXME
from aws_lambda_functions,
    UNNEST(vpc_config_security_group_ids) as sgs,
    UNNEST(vpc_config_subnet_ids) as sns
where sns in
    --  Find all subnets that include a route table that inclues a catchall route
    (select subnet_id
        from public.aws_ec2_route_tables
        inner join
            aws_ec2_route_table_associations on
                aws_ec2_route_table_associations.route_table_cq_id = aws_ec2_route_tables.cq_id
        where aws_ec2_route_tables.cq_id in
            --  Find all routes in any route table that contains a route to 0.0.0.0/0 or ::/0
            (select route_table_cq_id
                from public.aws_ec2_route_table_routes
                where destination_cidr_block = '0.0.0.0/0'
                    or destination_ipv6_cidr_block = '::/0'))
    and sgs in
    -- 	Find all functions that have egress rule that allows access to all ip addresses
    (select group_id
        from aws_ec2_instance_security_groups
        inner join view_aws_security_group_egress_rules on group_id = id
        where (ip = '0.0.0.0/0'
                or ip = '::/0') )
union
-- Find all Lambda functions that do not run in a VPC
select distinct
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Find all ec2 instances that have unrestricted access to the internet' AS title,
    account_id,
    arn AS resource_id,
    'fail' AS status -- TODO FIXME
from aws_lambda_functions
where vpc_config_vpc_id is null
    or vpc_config_vpc_id = ''

-- Note: We do not restrict the search to specific Runtimes
