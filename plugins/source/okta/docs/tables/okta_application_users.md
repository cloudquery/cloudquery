# Table: okta_application_users



The composite primary key for this table is (**app_id**, **id**).

## Relations
This table depends on [okta_applications](okta_applications.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|app_id (PK)|String|
|created|Timestamp|
|credentials|JSON|
|external_id|String|
|id (PK)|String|
|last_sync|Timestamp|
|last_updated|Timestamp|
|password_changed|Timestamp|
|profile|JSON|
|scope|String|
|status|String|
|status_changed|Timestamp|
|sync_state|String|