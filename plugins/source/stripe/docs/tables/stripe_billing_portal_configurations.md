# Table: stripe_billing_portal_configurations

This table shows data for Stripe Billing Portal Configurations.

https://stripe.com/docs/api/customer_portal/configuration

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|active|`bool`|
|application|`json`|
|business_profile|`json`|
|created|`timestamp[us, tz=UTC]`|
|default_return_url|`utf8`|
|features|`json`|
|is_default|`bool`|
|livemode|`bool`|
|login_page|`json`|
|metadata|`json`|
|object|`utf8`|
|updated|`int64`|