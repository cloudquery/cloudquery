
# Table: heroku_peerings
https://devcenter.heroku.com/articles/platform-api-reference#peering-attributes
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|aws_account_id|String||
|aws_region|String||
|aws_vpc_id|String||
|cidr_blocks|StringArray||
|expires|Timestamp||
|pcx_id|String||
|status|String||
|type|String||
|_cq_id|UUID|Internal CQ ID of the row|
|_cq_fetch_time|Timestamp|Internal CQ row of when fetch was started (this will be the same for all rows in a single fetch)|
