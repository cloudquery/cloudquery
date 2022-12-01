# Table: github_billing_action



The primary key for this table is **org**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|total_minutes_used|Int|
|total_paid_minutes_used|Float|
|included_minutes|Int|
|minutes_used_breakdown|JSON|