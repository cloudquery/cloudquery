# Table: stripe_accounts

https://stripe.com/docs/api/accounts

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.
## Relations

The following tables depend on stripe_accounts:
  - [stripe_capabilities](stripe_capabilities.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created (Incremental Key)|Timestamp|
|business_profile|JSON|
|business_type|String|
|capabilities|JSON|
|charges_enabled|Bool|
|company|JSON|
|controller|JSON|
|country|String|
|default_currency|String|
|deleted|Bool|
|details_submitted|Bool|
|email|String|
|external_accounts|JSON|
|future_requirements|JSON|
|individual|JSON|
|metadata|JSON|
|object|String|
|payouts_enabled|Bool|
|requirements|JSON|
|settings|JSON|
|tos_acceptance|JSON|
|type|String|