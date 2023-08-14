# Table: aws_ec2_capacity_reservations

This table shows data for Amazon Elastic Compute Cloud (EC2) Capacity Reservations.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeCapacityReservations.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|availability_zone|`utf8`|
|availability_zone_id|`utf8`|
|available_instance_count|`int64`|
|capacity_allocations|`json`|
|capacity_reservation_arn|`utf8`|
|capacity_reservation_fleet_id|`utf8`|
|capacity_reservation_id|`utf8`|
|create_date|`timestamp[us, tz=UTC]`|
|ebs_optimized|`bool`|
|end_date|`timestamp[us, tz=UTC]`|
|end_date_type|`utf8`|
|ephemeral_storage|`bool`|
|instance_match_criteria|`utf8`|
|instance_platform|`utf8`|
|instance_type|`utf8`|
|outpost_arn|`utf8`|
|owner_id|`utf8`|
|placement_group_arn|`utf8`|
|start_date|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|tenancy|`utf8`|
|total_instance_count|`int64`|