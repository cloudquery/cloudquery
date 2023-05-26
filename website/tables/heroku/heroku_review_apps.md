# Table: heroku_review_apps

This table shows data for Heroku Review Apps.

https://devcenter.heroku.com/articles/platform-api-reference#review-app

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|app|json|
|app_setup|json|
|branch|utf8|
|created_at|timestamp[us, tz=UTC]|
|creator|json|
|error_status|utf8|
|fork_repo|json|
|message|utf8|
|pipeline|json|
|pr_number|int64|
|status|utf8|
|updated_at|timestamp[us, tz=UTC]|
|wait_for_ci|bool|