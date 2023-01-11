# Table: aws_savingsplans_plans

https://docs.aws.amazon.com/savingsplans/latest/APIReference/API_SavingsPlan.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|commitment|String|
|currency|String|
|description|String|
|ec2_instance_family|String|
|end|String|
|offering_id|String|
|payment_option|String|
|product_types|StringArray|
|recurring_payment_amount|String|
|region|String|
|savings_plan_arn|String|
|savings_plan_id|String|
|savings_plan_type|String|
|start|String|
|state|String|
|tags|JSON|
|term_duration_in_seconds|Int|
|upfront_payment_amount|String|