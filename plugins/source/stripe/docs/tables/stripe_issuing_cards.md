# Table: stripe_issuing_cards

https://stripe.com/docs/api/issuing_cards

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created (Incremental Key)|Timestamp|
|brand|String|
|cancellation_reason|String|
|cardholder|JSON|
|currency|String|
|cvc|String|
|exp_month|Int|
|exp_year|Int|
|financial_account|String|
|last4|String|
|livemode|Bool|
|metadata|JSON|
|number|String|
|object|String|
|replaced_by|JSON|
|replacement_for|JSON|
|replacement_reason|String|
|shipping|JSON|
|spending_controls|JSON|
|status|String|
|type|String|
|wallets|JSON|