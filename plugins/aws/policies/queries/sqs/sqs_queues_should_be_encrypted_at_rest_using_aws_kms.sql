insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'SQS queues should be encrypted at rest using AWS KMS' as title,
    account_id,
    arn as resource_id,
    case when
        kms_master_key_id is null or kms_master_key_id = ''
    then 'fail' else 'pass' end as status
from aws_sqs_queues
