# Table: aws_config_delivery_channel_statuses

This table shows data for Config Delivery Channel Statuses.

https://docs.aws.amazon.com/config/latest/APIReference/API_DescribeDeliveryChannelStatus.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Relations

This table depends on [aws_config_delivery_channels](aws_config_delivery_channels).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|config_history_delivery_info|JSON|
|config_snapshot_delivery_info|JSON|
|config_stream_delivery_info|JSON|
|name (PK)|String|