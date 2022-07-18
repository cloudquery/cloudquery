insert into aws_policy_results
select
    :execution_time as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Amazon ECS task definitions should have secure networking modes and user definitions' as title,
    account_id,
    aws_ecs_task_definitions.arn as resource_id,
    case when
        aws_ecs_task_definitions.network_mode = 'host'
        and (
            aws_ecs_task_definition_container_definitions.privileged is distinct from TRUE
            and (
                aws_ecs_task_definition_container_definitions."user" = 'root' or aws_ecs_task_definition_container_definitions."user" is null
            )
        )
        then 'fail'
        else 'pass'
    end as status
from aws_ecs_task_definitions
     inner join
    aws_ecs_task_definition_container_definitions on
        aws_ecs_task_definitions.cq_id = aws_ecs_task_definition_container_definitions.task_definition_cq_id
