# Table: aws_sns_topics

This table shows data for Sns Topics.

https://docs.aws.amazon.com/sns/latest/api/API_GetTopicAttributes.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|delivery_policy|`json`|
|policy|`json`|
|effective_delivery_policy|`json`|
|display_name|`utf8`|
|owner|`utf8`|
|subscriptions_confirmed|`int64`|
|subscriptions_deleted|`int64`|
|subscriptions_pending|`int64`|
|kms_master_key_id|`utf8`|
|fifo_topic|`bool`|
|content_based_deduplication|`bool`|
|unknown_fields|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### SNS topics should be encrypted at rest using AWS KMS

```sql
SELECT
  'SNS topics should be encrypted at rest using AWS KMS' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN kms_master_key_id IS NULL OR kms_master_key_id = '' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_sns_topics;
```

### Logging of delivery status should be enabled for notification messages sent to a topic

```sql
SELECT
  'Logging of delivery status should be enabled for notification messages sent to a topic'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN (unknown_fields->'HTTPSuccessFeedbackRoleArn') IS NULL
  AND (unknown_fields->'FirehoseSuccessFeedbackRoleArn') IS NULL
  AND (unknown_fields->'LambdaSuccessFeedbackRoleArn') IS NULL
  AND (unknown_fields->'ApplicationSuccessFeedbackRoleArn') IS NULL
  AND (unknown_fields->'SQSSuccessFeedbackRoleArn') IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_sns_topics;
```

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


