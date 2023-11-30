insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Find all instances with a public IP address' AS title,
    account_id,
    arn as resource_id,
    case when public_ip_address is not null then 'fail' else 'pass' end as status
from
    aws_ec2_instances
