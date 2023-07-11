# Table: aws_emr_cluster_instance_fleets

This table shows data for Amazon EMR Cluster Instance Fleets.

https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleet.html

The composite primary key for this table is (**cluster_arn**, **id**).

## Relations

This table depends on [aws_emr_clusters](aws_emr_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cluster_arn (PK)|`utf8`|
|id (PK)|`utf8`|
|instance_fleet_type|`utf8`|
|instance_type_specifications|`json`|
|launch_specifications|`json`|
|name|`utf8`|
|provisioned_on_demand_capacity|`int64`|
|provisioned_spot_capacity|`int64`|
|resize_specifications|`json`|
|status|`json`|
|target_on_demand_capacity|`int64`|
|target_spot_capacity|`int64`|