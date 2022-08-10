
# Table: github_hook_deliveries
HookDelivery represents the data that is received from GitHub's Webhook Delivery API  GitHub API docs: - https://docs.github.com/en/rest/webhooks/repo-deliveries#list-deliveries-for-a-repository-webhook - https://docs.github.com/en/rest/webhooks/repo-deliveries#get-a-delivery-for-a-repository-webhook
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hook_cq_id|uuid|Unique CloudQuery ID of github_hooks table (FK)|
|id|bigint||
|guid|text||
|delivered_at_time|timestamp without time zone||
|redelivery|boolean||
|duration|float||
|status|text||
|status_code|bigint||
|event|text||
|action|text||
|installation_id|bigint||
|repository_id|bigint||
|request_headers|jsonb||
|request_raw_payload|bytea||
|response_headers|jsonb||
|response_raw_payload|bytea||
