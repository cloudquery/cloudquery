# Table: aws_lightsail_database_events


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_lightsail_databases`](aws_lightsail_databases.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|database_arn|String|
|created_at|Timestamp|
|event_categories|StringArray|
|message|String|
|resource|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|