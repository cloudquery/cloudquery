# Table: aws_redshift_events

This table shows data for Redshift Events.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_Event.html.

Only events occurred in the last 14 days are returned.

The composite primary key for this table is (**account_id**, **region**, **event_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|date|`timestamp[us, tz=UTC]`|
|event_categories|`list<item: utf8, nullable>`|
|event_id (PK)|`utf8`|
|message|`utf8`|
|severity|`utf8`|
|source_identifier|`utf8`|
|source_type|`utf8`|