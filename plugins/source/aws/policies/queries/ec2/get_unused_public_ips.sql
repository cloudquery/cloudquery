insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Unused EC2 EIPs should be removed' as title,
    account_id,
    public_ip as resource_id,
    case when
        instance_id is null
        then 'fail'
        else 'pass'
    end as status
from aws_ec2_eips
