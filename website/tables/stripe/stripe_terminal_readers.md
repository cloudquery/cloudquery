# Table: stripe_terminal_readers

This table shows data for Stripe Terminal Readers.

https://stripe.com/docs/api/terminal/readers

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|action|`json`|
|deleted|`bool`|
|device_sw_version|`utf8`|
|device_type|`utf8`|
|ip_address|`utf8`|
|label|`utf8`|
|livemode|`bool`|
|location|`json`|
|metadata|`json`|
|object|`utf8`|
|serial_number|`utf8`|
|status|`utf8`|