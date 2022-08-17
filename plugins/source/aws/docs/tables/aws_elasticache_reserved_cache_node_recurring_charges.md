
# Table: aws_elasticache_reserved_cache_node_recurring_charges
Contains the specific price and frequency of a recurring charges for a reserved cache node, or for a reserved cache node offering.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|reserved_cache_node_cq_id|uuid|Unique CloudQuery ID of aws_elasticache_reserved_cache_nodes table (FK)|
|recurring_charge_amount|float|The monetary amount of the recurring charge.|
|recurring_charge_frequency|text|The frequency of the recurring charge.|
