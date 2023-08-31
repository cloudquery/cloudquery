# Table: aws_config_delivery_channel_statuses

This table shows data for Config Delivery Channel Statuses.

https://docs.aws.amazon.com/config/latest/APIReference/API_DescribeDeliveryChannelStatus.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Relations

This table depends on [aws_config_delivery_channels](aws_config_delivery_channels).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|config_history_delivery_info|`json`|
|config_snapshot_delivery_info|`json`|
|config_stream_delivery_info|`json`|
|name (PK)|`utf8`|