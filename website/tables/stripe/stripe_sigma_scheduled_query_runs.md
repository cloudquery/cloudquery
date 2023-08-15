# Table: stripe_sigma_scheduled_query_runs

This table shows data for Stripe Sigma Scheduled Query Runs.

https://stripe.com/docs/api/sigma/scheduled_queries

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|data_load_time|`int64`|
|error|`json`|
|file|`json`|
|livemode|`bool`|
|object|`utf8`|
|result_available_until|`int64`|
|sql|`utf8`|
|status|`utf8`|
|title|`utf8`|