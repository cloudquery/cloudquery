insert into aws_policy_results
select distinct
    :execution_time as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CodeBuild project environment variables should not contain clear text credentials' as title,
    account_id,
    arn as resource_id,
    case when
            aws_codebuild_project_environment_variables.type = 'PLAINTEXT'
            and (
                UPPER(
                    aws_codebuild_project_environment_variables.name
                ) like '%ACCESS_KEY%' or UPPER(
                    aws_codebuild_project_environment_variables.name
                ) like '%SECRET%' or UPPER(
                    aws_codebuild_project_environment_variables.name
                ) like '%PASSWORD%'
            )
            then 'fail'
        else 'pass'
    end as status
from aws_codebuild_projects
     inner join aws_codebuild_project_environment_variables on
    aws_codebuild_projects.cq_id = aws_codebuild_project_environment_variables.project_cq_id
