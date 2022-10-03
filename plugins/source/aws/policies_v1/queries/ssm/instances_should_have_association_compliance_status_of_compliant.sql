insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Instances managed by Systems Manager should have an association compliance status of COMPLIANT' as title,
    aws_ssm_instances.account_id,
    aws_ssm_instances.arn,
    case when
        aws_ssm_instance_compliance_items.compliance_type = 'Association'
        and aws_ssm_instance_compliance_items.status is distinct from 'COMPLIANT'
    then 'fail' else 'pass' end as status
from
    aws_ssm_instances
inner join aws_ssm_instance_compliance_items on aws_ssm_instances.arn = aws_ssm_instance_compliance_items.instance_arn
