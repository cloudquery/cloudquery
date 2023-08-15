# Table: stripe_terminal_configurations

This table shows data for Stripe Terminal Configurations.

https://stripe.com/docs/api/terminal/configuration

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|bbpos_wisepos_e|`json`|
|deleted|`bool`|
|is_account_default|`bool`|
|livemode|`bool`|
|object|`utf8`|
|tipping|`json`|
|verifone_p400|`json`|