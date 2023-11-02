# Table: github_billing_action

This table shows data for Github Billing Action.

The primary key for this table is **org**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|total_minutes_used|`int64`|
|total_paid_minutes_used|`float64`|
|included_minutes|`int64`|
|minutes_used_breakdown|`json`|