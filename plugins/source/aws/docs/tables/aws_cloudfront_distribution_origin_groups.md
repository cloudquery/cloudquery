
# Table: aws_cloudfront_distribution_origin_groups
An origin group includes two origins (a primary origin and a second origin to failover to) and a failover criteria that you specify
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|distribution_cq_id|uuid|Unique CloudQuery ID of aws_cloudfront_distributions table (FK)|
|failover_criteria_status_codes|integer[]|The items (status codes) for an origin group.|
|id|text|The origin group's ID.|
|members_origin_ids|text[]|Items (origins) in an origin group.|
