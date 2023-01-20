# Table: stripe_terminal_configurations

https://stripe.com/docs/api/terminal_configurations

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|bbpos_wisepos_e|JSON|
|deleted|Bool|
|is_account_default|Bool|
|livemode|Bool|
|object|String|
|tipping|JSON|
|verifone_p400|JSON|