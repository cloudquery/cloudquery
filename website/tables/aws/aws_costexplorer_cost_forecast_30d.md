# Table: aws_costexplorer_cost_forecast_30d

This table shows data for AWS Cost Explorer cost forecast for the next 30 days.

https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostForecast.html
To sync this table you must set the 'use_paid_apis' option to 'true' in the AWS provider configuration. 

The composite primary key for this table is (**account_id**, **start_date**, **end_date**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|start_date (PK)|`utf8`|
|end_date (PK)|`utf8`|
|mean_value|`utf8`|
|prediction_interval_lower_bound|`utf8`|
|prediction_interval_upper_bound|`utf8`|
|time_period|`json`|