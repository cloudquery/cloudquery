# Table: github_billing_storage


The primary key for this table is **org**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|org (PK)|String|
|days_left_in_billing_cycle|Int|
|estimated_paid_storage_for_month|Float|
|estimated_storage_for_month|Int|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|