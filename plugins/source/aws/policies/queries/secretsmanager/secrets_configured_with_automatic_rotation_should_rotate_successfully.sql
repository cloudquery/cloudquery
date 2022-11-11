insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Secrets Manager secrets configured with automatic rotation should rotate successfully' as title,
    account_id,
    arn as resource_id,
    case when
        (last_rotated_date is null and created_date > now() - INTERVAL '1 day' * (rotation_rules->>'AutomaticallyAfterDays')::integer)
        or (last_rotated_date is not null and last_rotated_date > now() - INTERVAL '1 day' * (rotation_rules->>'AutomaticallyAfterDays')::integer)
    then 'fail' else 'pass' end as status
from aws_secretsmanager_secrets
