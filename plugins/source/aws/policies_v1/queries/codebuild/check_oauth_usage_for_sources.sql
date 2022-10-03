insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CodeBuild GitHub or Bitbucket source repository URLs should use OAuth' as title,
    account_id,
    arn as resource_id,
    case when
        source->>'Type' IN ('GITHUB', 'BITBUCKET') AND source->'Auth'->>'Type' != 'OAUTH'
      then 'fail'
      else 'pass'
    end as status
from aws_codebuild_projects
