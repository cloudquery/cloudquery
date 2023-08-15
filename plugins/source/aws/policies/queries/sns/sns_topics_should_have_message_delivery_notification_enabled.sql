insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Logging of delivery status should be enabled for notification messages sent to a topic' as title,
    account_id,
    arn as resource_id,
    case when
        unknown_fields->'HTTPSuccessFeedbackRoleArn' is Null
        AND  unknown_fields->'FirehoseSuccessFeedbackRoleArn' is Null
        AND  unknown_fields->'LambdaSuccessFeedbackRoleArn' is Null
        AND  unknown_fields->'ApplicationSuccessFeedbackRoleArn' is Null
        AND  unknown_fields->'SQSSuccessFeedbackRoleArn' is Null
    then 'fail' else 'pass' end as status
from aws_sns_topics