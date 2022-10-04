insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Attached EBS volumes should be encrypted at rest' as title,
    account_id,
    arn as resource_id,
    case when
        encrypted is FALSE
        then 'fail'
        else 'pass'
    end as status
from aws_ec2_ebs_volumes
