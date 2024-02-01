# Table: aws_config_delivery_channel_statuses

This table shows data for Config Delivery Channel Statuses.

https://docs.aws.amazon.com/config/latest/APIReference/API_DescribeDeliveryChannelStatus.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **name**).
## Relations

This table depends on [aws_config_delivery_channels](aws_config_delivery_channels.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|config_history_delivery_info|`json`|
|config_snapshot_delivery_info|`json`|
|config_stream_delivery_info|`json`|
|name|`utf8`|