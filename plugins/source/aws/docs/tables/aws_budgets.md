# Table: aws_budgets

This table shows data for Budgets.

https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html

The composite primary key for this table is (**account_id**, **region**, **budget_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|budget_name (PK)|`utf8`|
|budget_type|`utf8`|
|time_unit|`utf8`|
|auto_adjust_data|`json`|
|budget_limit|`json`|
|calculated_spend|`json`|
|cost_filters|`json`|
|cost_types|`json`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|planned_budget_limits|`json`|
|time_period|`json`|