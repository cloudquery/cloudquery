# Table: okta_applications

The primary key for this table is **id**.

## Relations

The following tables depend on okta_applications:
  - [okta_application_group_assignments](okta_application_group_assignments.md)
  - [okta_application_users](okta_application_users.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|accessibility|JSON|
|created|Timestamp|
|features|StringArray|
|label|String|
|last_updated|Timestamp|
|licensing|JSON|
|profile|JSON|
|sign_on_mode|String|
|status|String|
|visibility|JSON|
|_embedded|JSON|
|_links|JSON|
|additional_properties|JSON|