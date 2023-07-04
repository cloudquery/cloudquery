# Table: gcp_cloudsupport_cases

This table shows data for GCP Cloudsupport Cases.

https://cloud.google.com/support/docs/reference/rest/v2beta/cases#Case

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|classification|`json`|
|create_time|`utf8`|
|creator|`json`|
|description|`utf8`|
|display_name|`utf8`|
|escalated|`bool`|
|language_code|`utf8`|
|name (PK)|`utf8`|
|priority|`utf8`|
|severity|`utf8`|
|state|`utf8`|
|subscriber_email_addresses|`list<item: utf8, nullable>`|
|test_case|`bool`|
|time_zone|`utf8`|
|update_time|`utf8`|