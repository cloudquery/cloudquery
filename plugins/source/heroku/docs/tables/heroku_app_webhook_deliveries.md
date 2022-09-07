
# Table: heroku_app_webhook_deliveries
https://devcenter.heroku.com/articles/platform-api-reference#app-webhook-delivery-attributes
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|created_at|Timestamp||
|event|JSON||
|id|String||
|last_attempt|JSON||
|next_attempt_at|Timestamp||
|num_attempts|Int||
|status|String||
|updated_at|Timestamp||
|webhook|JSON||
|_cq_id|UUID|Internal CQ ID of the row|
|_cq_fetch_time|Timestamp|Internal CQ row of when fetch was started (this will be the same for all rows in a single fetch)|
