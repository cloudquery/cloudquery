insert into aws_policy_results
select distinct
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Find all lambda functions that have unrestricted access to the internet' AS title,
    account_id,
    arn AS resource_id,
    'fail' AS status -- TODO FIXME
from aws_lambda_functions,
    JSONB_ARRAY_ELEMENTS_TEXT(configuration->'VpcConfig'->'SecurityGroupIds') as sgs,
     JSONB_ARRAY_ELEMENTS_TEXT(configuration->'VpcConfig'->' SubnetIds') as sns
where sns in
    --  Find all subnets that include a route table that inclues a catchall route
    (select a->>'SubnetId'
        from public.aws_ec2_route_tables, jsonb_array_elements(associations) a, jsonb_array_elements(routes) r
        where r->>'DestinationCidrBlock' = '0.0.0.0/0' OR r->>'DestinationIpv6CidrBlock' = '::/0'
    )
    and sgs in
    -- 	Find all functions that have egress rule that allows access to all ip addresses
    (select id from view_aws_security_group_egress_rules where ip = '0.0.0.0/0' or ip6 = '::/0')
union
-- Find all Lambda functions that do not run in a VPC
select distinct
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Find all lambda functions that have unrestricted access to the internet' AS title,
    account_id,
    arn AS resource_id,
    'fail' AS status -- TODO FIXME
from aws_lambda_functions
where configuration->'VpcConfig'->>'VpcId' is null

-- Note: We do not restrict the search to specific Runtimes
