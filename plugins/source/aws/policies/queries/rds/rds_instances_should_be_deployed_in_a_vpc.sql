insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'RDS instances should be deployed in a VPC' as title,
    account_id,
    arn AS resource_id,
    case when db_subnet_group->>'VpcId' is null then 'fail' else 'pass' end as status
from aws_rds_instances
