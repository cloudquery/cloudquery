insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Secrets Manager secrets should have automatic rotation enabled' as title,
    account_id,
    arn as resource_id,
    case when
        rotation_enabled is distinct from TRUE
    then 'fail' else 'pass' end as status
from aws_secretsmanager_secrets
