insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'AWS KMS keys should not be unintentionally deleted' as title,
    account_id,
    arn as resource_id,
    case when key_state = 'PendingDeletion' and key_manager = 'CUSTOMER' then 'fail' else 'pass' end as status
from aws_kms_keys
