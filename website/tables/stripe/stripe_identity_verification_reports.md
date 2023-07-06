# Table: stripe_identity_verification_reports

This table shows data for Stripe Identity Verification Reports.

https://stripe.com/docs/api/identity/verification_reports

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|document|`json`|
|id_number|`json`|
|livemode|`bool`|
|object|`utf8`|
|options|`json`|
|selfie|`json`|
|type|`utf8`|
|verification_session|`utf8`|