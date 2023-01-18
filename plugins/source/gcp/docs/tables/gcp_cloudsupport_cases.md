# Table: gcp_cloudsupport_cases

https://cloud.google.com/support/docs/reference/rest/v2beta/cases#Case

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
|classification|JSON|
|create_time|String|
|creator|JSON|
|description|String|
|display_name|String|
|escalated|Bool|
|language_code|String|
|priority|String|
|severity|String|
|state|String|
|subscriber_email_addresses|StringArray|
|test_case|Bool|
|time_zone|String|
|update_time|String|