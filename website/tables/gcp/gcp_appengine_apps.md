# Table: gcp_appengine_apps

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#Application

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|id|String|
|dispatch_rules|JSON|
|auth_domain|String|
|location_id|String|
|code_bucket|String|
|default_cookie_expiration|Int|
|serving_status|String|
|default_hostname|String|
|default_bucket|String|
|service_account|String|
|iap|JSON|
|gcr_domain|String|
|database_type|String|
|feature_settings|JSON|