insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Remove unused Secrets Manager secrets' as title,
    account_id,
    arn as resource_id,
    case when
        (last_accessed_date is null and created_date > now() - INTERVAL '90 days')
        or (last_accessed_date is not null and last_accessed_date > now() - INTERVAL '90 days')
    then 'fail' else 'pass' end as status
from aws_secretsmanager_secrets
