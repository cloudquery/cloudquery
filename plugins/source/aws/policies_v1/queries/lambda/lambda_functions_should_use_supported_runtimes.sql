insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Lambda functions should use supported runtimes' as title,
    f.account_id,
    f.arn AS resource_id,
    case when r.name is null then 'fail'
    else 'pass' end AS status
from aws_lambda_functions f
left join aws_lambda_runtimes r on r.name=f.configuration->>'Runtime'
where f.configuration->>'PackageType' != 'Image'
