insert into aws_policy_results
with endpoints as (
    select vpc_endpoint_id
    from aws_ec2_vpc_endpoints
    where vpc_endpoint_type = 'Interface'
        and service_name ~ CONCAT(
            'com.amazonaws.', region, '.ec2'
        )
)

select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Amazon EC2 should be configured to use VPC endpoints that are created for the Amazon EC2 service' as title,
    account_id,
    vpc_id as resource_id,
    case when
        endpoints.vpc_endpoint_id is null
        then 'fail'
        else 'pass'
    end as status
from aws_ec2_vpcs
left join endpoints
    on aws_ec2_vpcs.vpc_id = endpoints.vpc_endpoint_id
