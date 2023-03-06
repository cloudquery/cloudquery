# Table: stripe_identity_verification_reports

https://stripe.com/docs/api/identity_verification_reports

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created (Incremental Key)|Timestamp|
|document|JSON|
|id_number|JSON|
|livemode|Bool|
|object|String|
|options|JSON|
|selfie|JSON|
|type|String|
|verification_session|String|