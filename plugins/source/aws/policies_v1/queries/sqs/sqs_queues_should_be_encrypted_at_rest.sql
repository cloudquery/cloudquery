insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Amazon SQS queues should be encrypted at rest' as title,
    account_id,
    arn as resource_id,
    case when
        (kms_master_key_id is null or kms_master_key_id = '') and sqs_managed_sse_enabled = false
    then 'fail' else 'pass' end as status
from aws_sqs_queues
