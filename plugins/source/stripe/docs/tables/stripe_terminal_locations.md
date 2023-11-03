# Table: stripe_terminal_locations

This table shows data for Stripe Terminal Locations.

https://stripe.com/docs/api/terminal/locations

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|address|`json`|
|configuration_overrides|`utf8`|
|deleted|`bool`|
|display_name|`utf8`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|