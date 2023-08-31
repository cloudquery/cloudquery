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
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|accessibility|`json`|
|created|`timestamp[us, tz=UTC]`|
|features|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|label|`utf8`|
|last_updated|`timestamp[us, tz=UTC]`|
|licensing|`json`|
|profile|`json`|
|sign_on_mode|`utf8`|
|status|`utf8`|
|visibility|`json`|
|_embedded|`json`|
|_links|`json`|
|additional_properties|`json`|