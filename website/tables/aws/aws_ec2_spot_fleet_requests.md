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
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|activity_status|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|spot_fleet_request_config|`json`|
|spot_fleet_request_id (PK)|`utf8`|
|spot_fleet_request_state|`utf8`|