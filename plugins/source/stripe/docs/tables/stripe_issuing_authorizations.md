# Table: stripe_issuing_authorizations

https://stripe.com/docs/api/issuing_authorizations

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
|amount|Int|
|amount_details|JSON|
|approved|Bool|
|authorization_method|String|
|balance_transactions|JSON|
|card|JSON|
|cardholder|JSON|
|currency|String|
|livemode|Bool|
|merchant_amount|Int|
|merchant_currency|String|
|merchant_data|JSON|
|metadata|JSON|
|network_data|JSON|
|object|String|
|pending_request|JSON|
|request_history|JSON|
|status|String|
|transactions|JSON|
|treasury|JSON|
|verification_data|JSON|
|wallet|String|