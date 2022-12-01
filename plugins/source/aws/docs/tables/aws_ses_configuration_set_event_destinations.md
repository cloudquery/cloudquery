# Table: aws_ses_configuration_set_event_destinations

https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_EventDestination.html

The composite primary key for this table is (**account_id**, **region**, **configuration_set_name**, **name**).

## Relations
This table depends on [aws_ses_configuration_sets](aws_ses_configuration_sets.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|configuration_set_name (PK)|String|
|name (PK)|String|
|matching_event_types|StringArray|
|cloud_watch_destination|JSON|
|enabled|Bool|
|kinesis_firehose_destination|JSON|
|pinpoint_destination|JSON|
|sns_destination|JSON|