# Table: aws_ses_configuration_set_event_destinations

This table shows data for Amazon Simple Email Service (SES) Configuration Set Event Destinations.

https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_EventDestination.html

The composite primary key for this table is (**account_id**, **region**, **configuration_set_name**, **name**).

## Relations

This table depends on [aws_ses_configuration_sets](aws_ses_configuration_sets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|configuration_set_name (PK)|`utf8`|
|name (PK)|`utf8`|
|matching_event_types|`list<item: utf8, nullable>`|
|cloud_watch_destination|`json`|
|enabled|`bool`|
|kinesis_firehose_destination|`json`|
|pinpoint_destination|`json`|
|sns_destination|`json`|