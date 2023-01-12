# Table: stripe_billing_portal_configurations

https://stripe.com/docs/api/billing_portal_configurations

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|active|Bool|
|application|JSON|
|business_profile|JSON|
|created|Timestamp|
|default_return_url|String|
|features|JSON|
|is_default|Bool|
|livemode|Bool|
|login_page|JSON|
|metadata|JSON|
|object|String|
|updated|Int|