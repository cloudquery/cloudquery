# Table: digitalocean_billing_history

This table shows data for DigitalOcean Billing History.

https://docs.digitalocean.com/reference/api/api-reference/#operation/billingHistory_list

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|description|`utf8`|
|amount|`utf8`|
|invoice_id|`utf8`|
|invoice_uuid|`utf8`|
|date|`timestamp[us, tz=UTC]`|
|type|`utf8`|