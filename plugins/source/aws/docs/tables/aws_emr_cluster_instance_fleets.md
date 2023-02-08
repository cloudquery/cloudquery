# Table: aws_emr_cluster_instance_fleets

https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleet.html

The composite primary key for this table is (**cluster_arn**, **id**).

## Relations

This table depends on [aws_emr_clusters](aws_emr_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|cluster_arn (PK)|String|
|id (PK)|String|
|instance_fleet_type|String|
|instance_type_specifications|JSON|
|launch_specifications|JSON|
|name|String|
|provisioned_on_demand_capacity|Int|
|provisioned_spot_capacity|Int|
|status|JSON|
|target_on_demand_capacity|Int|
|target_spot_capacity|Int|