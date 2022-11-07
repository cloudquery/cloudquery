# Table: okta_applications



The primary key for this table is **id**.

## Relations
The following tables depend on okta_applications:
  - [okta_application_users](okta_application_users.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|label|String|
|name|String|
|status|String|
|sign_on_mode|String|
|created|Timestamp|
|embedded|JSON|
|links|JSON|
|last_updated|Timestamp|
|accessibility|JSON|
|credentials|JSON|
|licensing|JSON|
|settings|JSON|
|visibility|JSON|
|features|StringArray|
|profile|JSON|