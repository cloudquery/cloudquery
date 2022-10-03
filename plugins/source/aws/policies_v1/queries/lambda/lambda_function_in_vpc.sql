insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Lambda functions should be in a VPC' AS title,
    account_id,
    arn as resource_id,
    case when configuration->'VpcConfig'->>'VpcId' is null or configuration->'VpcConfig'->>'VpcId' = '' then 'fail' else 'pass' end as status
from aws_lambda_functions
