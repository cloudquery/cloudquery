# Table: aws_budget_actions

This table shows data for Budget Actions.

https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Action.html

The composite primary key for this table is (**account_id**, **region**, **action_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|action_id (PK)|`utf8`|
|action_threshold|`json`|
|action_type|`utf8`|
|approval_model|`utf8`|
|budget_name|`utf8`|
|definition|`json`|
|execution_role_arn|`utf8`|
|notification_type|`utf8`|
|status|`utf8`|
|subscribers|`json`|