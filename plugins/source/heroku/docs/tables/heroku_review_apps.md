
# Table: heroku_review_apps
https://devcenter.heroku.com/articles/platform-api-reference#review-app-attributes
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|app|JSON||
|app_setup|JSON||
|branch|String||
|created_at|Timestamp||
|creator|JSON||
|error_status|String||
|fork_repo|JSON||
|id|String||
|message|String||
|pipeline|JSON||
|pr_number|Int||
|status|String||
|updated_at|Timestamp||
|wait_for_ci|Bool||
|_cq_id|UUID|Internal CQ ID of the row|
|_cq_fetch_time|Timestamp|Internal CQ row of when fetch was started (this will be the same for all rows in a single fetch)|
