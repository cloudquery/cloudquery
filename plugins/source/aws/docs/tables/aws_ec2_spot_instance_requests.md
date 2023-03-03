# Table: aws_ec2_spot_instance_requests

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotInstanceRequest.html

The composite primary key for this table is (**account_id**, **region**, **spot_instance_request_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|tags|JSON|
|actual_block_hourly_price|String|
|availability_zone_group|String|
|block_duration_minutes|Int|
|create_time|Timestamp|
|fault|JSON|
|instance_id|String|
|instance_interruption_behavior|String|
|launch_group|String|
|launch_specification|JSON|
|launched_availability_zone|String|
|product_description|String|
|spot_instance_request_id (PK)|String|
|spot_price|String|
|state|String|
|status|JSON|
|type|String|
|valid_from|Timestamp|
|valid_until|Timestamp|