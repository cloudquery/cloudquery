# Table: aws_sns_subscriptions

This table shows data for Sns Subscriptions.

https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|delivery_policy|`json`|
|effective_delivery_policy|`json`|
|filter_policy|`json`|
|redrive_policy|`json`|
|endpoint|`utf8`|
|owner|`utf8`|
|protocol|`utf8`|
|subscription_arn|`utf8`|
|topic_arn|`utf8`|
|confirmation_was_authenticated|`bool`|
|pending_confirmation|`bool`|
|raw_message_delivery|`bool`|
|subscription_role_arn|`utf8`|
|unknown_fields|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Unused SNS topic

```sql
WITH
  subscription AS (SELECT DISTINCT topic_arn FROM aws_sns_subscriptions)
SELECT
  'Unused SNS topic' AS title,
  topic.account_id,
  topic.arn AS resource_id,
  'fail' AS status
FROM
  aws_sns_topics AS topic
  LEFT JOIN subscription ON subscription.topic_arn = topic.arn
WHERE
  subscription.topic_arn IS NULL;
```


