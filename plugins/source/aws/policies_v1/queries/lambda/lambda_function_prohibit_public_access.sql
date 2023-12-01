insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Lambda functions should prohibit public access' as title,
    account_id,
    arn as resource_id,
    'fail' as status -- TODO FIXME
from aws_lambda_functions,
    jsonb_array_elements(
        case jsonb_typeof(policy_document -> 'Statement')
            when
                'string' then jsonb_build_array(policy_document ->> 'Statement')
            when 'array' then policy_document -> 'Statement'
        end
    ) as statement
where statement ->> 'Effect' = 'Allow'
    and (
        statement ->> 'Principal' = '*'
        or statement -> 'Principal' ->> 'AWS' = '*'
         or ( case jsonb_typeof(statement -> 'Principal' -> 'AWS')
            when 'string' then jsonb_build_array(statement -> 'Principal' ->> 'AWS')
            when 'array' then (statement -> 'Principal' ->> 'AWS')::JSONB
	     end)::JSONB ? '*'
    )
