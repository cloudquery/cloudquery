# Table: aws_redshift_events

This table shows data for Redshift Events.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_Event.html.

Only events occurred in the last hour are returned.

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|date|Timestamp|
|event_categories|StringArray|
|event_id|String|
|message|String|
|severity|String|
|source_identifier|String|
|source_type|String|