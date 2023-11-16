# Table: digitalocean_balances

This table shows data for DigitalOcean Balances.

https://docs.digitalocean.com/reference/api/api-reference/#operation/balance_get

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|month_to_date_balance|`utf8`|
|account_balance|`utf8`|
|month_to_date_usage|`utf8`|
|generated_at|`timestamp[us, tz=UTC]`|