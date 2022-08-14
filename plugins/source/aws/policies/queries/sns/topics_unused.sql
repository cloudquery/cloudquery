insert into aws_policy_results
with subscription as (select distinct topic_arn from aws_sns_subscriptions)
select :'execution_time'  as execution_time,
       :'framework'       as framework,
       :'check_id'        as check_id,
       'Unused SNS topic' as title,
       topic.account_id,
       topic.arn          as resource_id,
       'fail'             as status
from aws_sns_topics topic
         left join subscription on subscription.topic_arn = topic.arn
where subscription.topic_arn is null