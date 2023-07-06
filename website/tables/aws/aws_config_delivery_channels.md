# Table: aws_config_delivery_channels

This table shows data for Config Delivery Channels.

https://docs.aws.amazon.com/config/latest/APIReference/API_DeliveryChannel.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Relations

The following tables depend on aws_config_delivery_channels:
  - [aws_config_delivery_channel_statuses](aws_config_delivery_channel_statuses)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|config_snapshot_delivery_properties|`json`|
|name (PK)|`utf8`|
|s3_bucket_name|`utf8`|
|s3_key_prefix|`utf8`|
|s3_kms_key_arn|`utf8`|
|sns_topic_arn|`utf8`|