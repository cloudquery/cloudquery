insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'EC2 subnets should not automatically assign public IP addresses' as title,
    owner_id as account_id,
    arn as resource_id,
    case when
        map_public_ip_on_launch is true
        then 'fail'
        else 'pass'
    end
from aws_ec2_subnets
