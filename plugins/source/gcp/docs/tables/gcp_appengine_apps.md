# Table: gcp_appengine_apps

This table shows data for GCP App Engine Apps.

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#Application

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|id|`utf8`|
|dispatch_rules|`json`|
|auth_domain|`utf8`|
|location_id|`utf8`|
|code_bucket|`utf8`|
|default_cookie_expiration|`int64`|
|serving_status|`utf8`|
|default_hostname|`utf8`|
|default_bucket|`utf8`|
|service_account|`utf8`|
|iap|`json`|
|gcr_domain|`utf8`|
|database_type|`utf8`|
|feature_settings|`json`|