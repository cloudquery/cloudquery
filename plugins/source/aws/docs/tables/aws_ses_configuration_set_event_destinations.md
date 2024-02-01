# Table: aws_ses_configuration_set_event_destinations

This table shows data for Amazon Simple Email Service (SES) Configuration Set Event Destinations.

https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_EventDestination.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **configuration_set_name**, **name**).
## Relations

This table depends on [aws_ses_configuration_sets](aws_ses_configuration_sets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|configuration_set_name|`utf8`|
|name|`utf8`|
|matching_event_types|`list<item: utf8, nullable>`|
|cloud_watch_destination|`json`|
|enabled|`bool`|
|kinesis_firehose_destination|`json`|
|pinpoint_destination|`json`|
|sns_destination|`json`|