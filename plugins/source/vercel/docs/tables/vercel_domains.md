# Table: vercel_domains

The primary key for this table is **id**.

## Relations

The following tables depend on vercel_domains:
  - [vercel_domain_records](vercel_domain_records.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|bought_at|Timestamp|
|cdn_enabled|Bool|
|config_verified_at|Timestamp|
|created_at|Timestamp|
|expires_at|Timestamp|
|id (PK)|String|
|intended_nameservers|StringArray|
|custom_nameservers|StringArray|
|name|String|
|nameservers|StringArray|
|ns_verified_at|Timestamp|
|ordered_at|Timestamp|
|renew|Bool|
|service_type|String|
|transfer_started_at|Timestamp|
|transferred_at|Timestamp|
|txt_verified_at|Timestamp|
|verification_record|String|
|verified|Bool|
|zone|Bool|
|creator|JSON|