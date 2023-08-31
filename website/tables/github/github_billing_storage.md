# Table: github_billing_storage

This table shows data for Github Billing Storage.

The primary key for this table is **org**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|days_left_in_billing_cycle|`int64`|
|estimated_paid_storage_for_month|`float64`|
|estimated_storage_for_month|`int64`|