# Table: stripe_disputes

https://stripe.com/docs/api/disputes

The primary key for this table is **id**.
It supports incremental syncs.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|amount|Int|
|balance_transactions|JSON|
|charge|JSON|
|created|Int|
|currency|String|
|evidence|JSON|
|evidence_details|JSON|
|is_charge_refundable|Bool|
|livemode|Bool|
|metadata|JSON|
|network_reason_code|String|
|object|String|
|payment_intent|JSON|
|reason|String|
|status|String|