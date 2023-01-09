# Table: aws_ses_configuration_sets

https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetConfigurationSet.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ses_configuration_sets:
  - [aws_ses_configuration_set_event_destinations](aws_ses_configuration_set_event_destinations.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|name|String|
|delivery_options|JSON|
|reputation_options|JSON|
|sending_options|JSON|
|suppression_options|JSON|
|tracking_options|JSON|
|vdm_options|JSON|