# Table: aws_ec2_spot_fleet_requests

This table shows data for Amazon Elastic Compute Cloud (EC2) Spot Fleet Requests.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotFleetRequestConfig.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **spot_fleet_request_id**).
## Relations

The following tables depend on aws_ec2_spot_fleet_requests:
  - [aws_ec2_spot_fleet_instances](aws_ec2_spot_fleet_instances.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|activity_status|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|spot_fleet_request_config|`json`|
|spot_fleet_request_id|`utf8`|
|spot_fleet_request_state|`utf8`|