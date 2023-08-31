# Table: aws_redshift_events

This table shows data for Redshift Events.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_Event.html.

Only events occurred in the last 14 days are returned.

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|date|`timestamp[us, tz=UTC]`|
|event_categories|`list<item: utf8, nullable>`|
|event_id|`utf8`|
|message|`utf8`|
|severity|`utf8`|
|source_identifier|`utf8`|
|source_type|`utf8`|