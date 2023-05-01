# Table: aws_costexplorer_cost_30d

This table shows data for AWS Cost Explorer Cost 30d.

https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostAndUsage.html
To sync this table you must set the 'use_paid_apis' option to 'true' in the AWS provider configuration. 

The composite primary key for this table is (**account_id**, **start_date**, **end_date**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|start_date (PK)|String|
|end_date (PK)|String|
|estimated|Bool|
|groups|JSON|
|time_period|JSON|
|total|JSON|