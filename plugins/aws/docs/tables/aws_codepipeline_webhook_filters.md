
# Table: aws_codepipeline_webhook_filters
The event criteria that specify when a webhook notification is sent to your URL.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|webhook_cq_id|uuid|Unique CloudQuery ID of aws_codepipeline_webhooks table (FK)|
|json_path|text|A JsonPath expression that is applied to the body/payload of the webhook|
|match_equals|text|The value selected by the JsonPath expression must match what is supplied in the MatchEquals field|
