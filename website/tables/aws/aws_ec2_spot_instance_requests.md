# Table: aws_ec2_spot_instance_requests

This table shows data for Amazon Elastic Compute Cloud (EC2) Spot Instance Requests.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotInstanceRequest.html

The composite primary key for this table is (**account_id**, **region**, **spot_instance_request_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|actual_block_hourly_price|`utf8`|
|availability_zone_group|`utf8`|
|block_duration_minutes|`int64`|
|create_time|`timestamp[us, tz=UTC]`|
|fault|`json`|
|instance_id|`utf8`|
|instance_interruption_behavior|`utf8`|
|launch_group|`utf8`|
|launch_specification|`json`|
|launched_availability_zone|`utf8`|
|product_description|`utf8`|
|spot_instance_request_id (PK)|`utf8`|
|spot_price|`utf8`|
|state|`utf8`|
|status|`json`|
|type|`utf8`|
|valid_from|`timestamp[us, tz=UTC]`|
|valid_until|`timestamp[us, tz=UTC]`|