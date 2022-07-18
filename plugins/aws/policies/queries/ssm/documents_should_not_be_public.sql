insert into aws_policy_results
select
    :execution_time as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'SSM documents should not be public' as title,
    account_id,
    arn as resource_id,
    case when
        'all' = ANY(account_ids)
    then 'fail' else 'pass' end as status
from aws_ssm_documents
where owner in (select account_id from aws_accounts)
