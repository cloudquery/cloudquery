# Table: aws_costexplorer_cost_custom

This table shows data for AWS Cost Explorer costs based on custom inputs.

https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostAndUsage.html
To sync this table you must set the 'use_paid_apis' option to 'true' in the AWS provider configuration as well as specify the request parameters in the 'table_options' attribute. 

The composite primary key for this table is (**account_id**, **start_date**, **end_date**, **input_hash**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|start_date (PK)|`utf8`|
|end_date (PK)|`utf8`|
|input_hash (PK)|`utf8`|
|input_json|`json`|
|estimated|`bool`|
|groups|`json`|
|time_period|`json`|
|total|`json`|