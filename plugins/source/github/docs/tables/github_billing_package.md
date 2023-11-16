# Table: github_billing_package

This table shows data for Github Billing Package.

The primary key for this table is **org**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|total_gigabytes_bandwidth_used|`int64`|
|total_paid_gigabytes_bandwidth_used|`int64`|
|included_gigabytes_bandwidth|`int64`|