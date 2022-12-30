# Table: stripe_disputes

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|amount|Int|
|balance_transactions|JSON|
|charge|JSON|
|created|Int|
|currency|String|
|evidence|JSON|
|evidence_details|JSON|
|id (PK)|String|
|is_charge_refundable|Bool|
|livemode|Bool|
|metadata|JSON|
|network_reason_code|String|
|object|String|
|payment_intent|JSON|
|reason|String|
|status|String|