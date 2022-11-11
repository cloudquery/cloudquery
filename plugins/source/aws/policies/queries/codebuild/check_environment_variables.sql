insert into aws_policy_results
select distinct
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CodeBuild project environment variables should not contain clear text credentials' as title,
    account_id,
    arn as resource_id,
    case when
            e->>'Type' = 'PLAINTEXT'
            and (
                UPPER(e->>'Name') like '%ACCESS_KEY%' or
                UPPER(e->>'Name') like '%SECRET%' or
                UPPER(e->>'Name') like '%PASSWORD%'
            )
            then 'fail'
        else 'pass'
    end as status
from aws_codebuild_projects, JSONB_ARRAY_ELEMENTS(environment->'EnvironmentVariables') as e
