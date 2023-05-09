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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|config_snapshot_delivery_properties|JSON|
|name (PK)|String|
|s3_bucket_name|String|
|s3_key_prefix|String|
|s3_kms_key_arn|String|
|sns_topic_arn|String|