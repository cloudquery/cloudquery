# Table: stripe_disputes

This table shows data for Stripe Disputes.

https://stripe.com/docs/api/disputes

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|amount|`int64`|
|balance_transactions|`json`|
|charge|`json`|
|currency|`utf8`|
|evidence|`json`|
|evidence_details|`json`|
|is_charge_refundable|`bool`|
|livemode|`bool`|
|metadata|`json`|
|network_reason_code|`utf8`|
|object|`utf8`|
|payment_intent|`json`|
|reason|`utf8`|
|status|`utf8`|