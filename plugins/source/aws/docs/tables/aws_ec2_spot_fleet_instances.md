# Table: aws_ec2_spot_fleet_instances

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ActiveInstance.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_ec2_spot_fleet_requests](aws_ec2_spot_fleet_requests.md).

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
|spot_fleet_request_id|String|
|instance_health|String|
|instance_id|String|
|instance_type|String|
|spot_instance_request_id|String|