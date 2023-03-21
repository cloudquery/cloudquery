# Table: aws_ec2_spot_fleet_requests

This table shows data for Amazon Elastic Compute Cloud (EC2) Spot Fleet Requests.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotFleetRequestConfig.html

The composite primary key for this table is (**account_id**, **region**, **spot_fleet_request_id**).

## Relations

The following tables depend on aws_ec2_spot_fleet_requests:
  - [aws_ec2_spot_fleet_instances](aws_ec2_spot_fleet_instances)

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
|activity_status|String|
|create_time|Timestamp|
|spot_fleet_request_config|JSON|
|spot_fleet_request_id (PK)|String|
|spot_fleet_request_state|String|