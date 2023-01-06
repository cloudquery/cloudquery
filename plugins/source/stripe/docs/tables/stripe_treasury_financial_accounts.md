# Table: stripe_treasury_financial_accounts

https://stripe.com/docs/api/treasury_financial_accounts

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|active_features|StringArray|
|balance|JSON|
|country|String|
|created|Int|
|features|JSON|
|financial_addresses|JSON|
|livemode|Bool|
|metadata|JSON|
|object|String|
|pending_features|StringArray|
|platform_restrictions|JSON|
|restricted_features|StringArray|
|status|String|
|status_details|JSON|
|supported_currencies|StringArray|