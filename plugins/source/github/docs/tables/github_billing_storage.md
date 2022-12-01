# Table: github_billing_storage



The primary key for this table is **org**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|days_left_in_billing_cycle|Int|
|estimated_paid_storage_for_month|Float|
|estimated_storage_for_month|Int|