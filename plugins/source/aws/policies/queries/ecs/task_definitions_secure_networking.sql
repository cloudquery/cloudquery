insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Amazon ECS task definitions should have secure networking modes and user definitions' as title,
    account_id,
    arn as resource_id,
    case when
        network_mode = 'host'
        and (c->>'Privileged')::boolean is distinct from true
        and (c->>'User' = 'root' or c->>'User' is null)
    then 'fail'
    else 'pass'
    end as status
from aws_ecs_task_definitions, jsonb_array_elements(aws_ecs_task_definitions.container_definitions) as c
