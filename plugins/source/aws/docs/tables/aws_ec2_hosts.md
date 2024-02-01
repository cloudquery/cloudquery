# Table: aws_ec2_hosts

This table shows data for Amazon Elastic Compute Cloud (EC2) Hosts.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Host.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|allocation_time|`timestamp[us, tz=UTC]`|
|allows_multiple_instance_types|`utf8`|
|asset_id|`utf8`|
|auto_placement|`utf8`|
|availability_zone|`utf8`|
|availability_zone_id|`utf8`|
|available_capacity|`json`|
|client_token|`utf8`|
|host_id|`utf8`|
|host_maintenance|`utf8`|
|host_properties|`json`|
|host_recovery|`utf8`|
|host_reservation_id|`utf8`|
|instances|`json`|
|member_of_service_linked_resource_group|`bool`|
|outpost_arn|`utf8`|
|owner_id|`utf8`|
|release_time|`timestamp[us, tz=UTC]`|
|state|`utf8`|