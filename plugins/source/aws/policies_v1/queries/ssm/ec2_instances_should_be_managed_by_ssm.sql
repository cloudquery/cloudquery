insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'EC2 instances should be managed by AWS Systems Manager' as title,
    aws_ec2_instances.account_id,
    aws_ec2_instances.arn as resource_id,
    case when
        aws_ssm_instances.instance_id is null
    then 'fail' else 'pass' end as status
from
    aws_ec2_instances
left outer join aws_ssm_instances on aws_ec2_instances.instance_id = aws_ssm_instances.instance_id
