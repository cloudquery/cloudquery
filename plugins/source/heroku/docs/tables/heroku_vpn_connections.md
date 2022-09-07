
# Table: heroku_vpn_connections
https://devcenter.heroku.com/articles/platform-api-reference#vpn-connection-attributes
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|String||
|ike_version|Int||
|name|String||
|public_ip|String||
|routable_cidrs|StringArray||
|space_cidr_block|String||
|status|String||
|status_message|String||
|tunnels|JSON||
|_cq_id|UUID|Internal CQ ID of the row|
|_cq_fetch_time|Timestamp|Internal CQ row of when fetch was started (this will be the same for all rows in a single fetch)|
