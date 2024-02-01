# Table: aws_ec2_spot_fleet_instances

This table shows data for Amazon Elastic Compute Cloud (EC2) Spot Fleet Instances.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ActiveInstance.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

This table depends on [aws_ec2_spot_fleet_requests](aws_ec2_spot_fleet_requests.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|spot_fleet_request_id|`utf8`|
|instance_health|`utf8`|
|instance_id|`utf8`|
|instance_type|`utf8`|
|spot_instance_request_id|`utf8`|