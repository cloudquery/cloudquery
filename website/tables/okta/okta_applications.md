# Table: okta_applications

This table shows data for Okta Applications.

The primary key for this table is **id**.

## Relations

The following tables depend on okta_applications:
  - [okta_application_group_assignments](okta_application_group_assignments)
  - [okta_application_users](okta_application_users)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|accessibility|extension_type<storage=binary>|
|created|timestamp[us, tz=UTC]|
|features|list<item: utf8, nullable>|
|id (PK)|utf8|
|label|utf8|
|last_updated|timestamp[us, tz=UTC]|
|licensing|extension_type<storage=binary>|
|profile|extension_type<storage=binary>|
|sign_on_mode|utf8|
|status|utf8|
|visibility|extension_type<storage=binary>|
|_embedded|extension_type<storage=binary>|
|_links|extension_type<storage=binary>|
|additional_properties|extension_type<storage=binary>|