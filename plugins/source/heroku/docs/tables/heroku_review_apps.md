# Table: heroku_review_apps
https://devcenter.heroku.com/articles/platform-api-reference#review-app-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|app|JSON|
|app_setup|JSON|
|branch|String|
|created_at|Timestamp|
|creator|JSON|
|error_status|String|
|fork_repo|JSON|
|id (PK)|String|
|message|String|
|pipeline|JSON|
|pr_number|Int|
|status|String|
|updated_at|Timestamp|
|wait_for_ci|Bool|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|