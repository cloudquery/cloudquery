
# Table: gcp_compute_instance_scheduling_node_affinities
Node Affinity: the configuration of desired nodes onto which this Instance could be scheduled
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid|Unique ID of gcp_compute_instances table (FK)|
|key|text|Corresponds to the label key of Node resource|
|operator|text|Defines the operation of node selection Valid operators are IN for affinity and NOT_IN for anti-affinity|
|values|text[]|Corresponds to the label values of Node resource|
