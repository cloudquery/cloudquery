# Table: aws_ses_configuration_sets

This table shows data for Amazon Simple Email Service (SES) Configuration Sets.

https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetConfigurationSet.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ses_configuration_sets:
  - [aws_ses_configuration_set_event_destinations](aws_ses_configuration_set_event_destinations)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|name|`utf8`|
|delivery_options|`json`|
|reputation_options|`json`|
|sending_options|`json`|
|suppression_options|`json`|
|tracking_options|`json`|
|vdm_options|`json`|