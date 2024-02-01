# Table: aws_costexplorer_cost_forecast_30d

This table shows data for AWS Cost Explorer cost forecast for the next 30 days.

https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostForecast.html
To sync this table you must set the 'use_paid_apis' option to 'true' in the AWS provider configuration. 

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **start_date**, **end_date**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|start_date|`utf8`|
|end_date|`utf8`|
|mean_value|`utf8`|
|prediction_interval_lower_bound|`utf8`|
|prediction_interval_upper_bound|`utf8`|
|time_period|`json`|