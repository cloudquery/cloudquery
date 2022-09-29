# Table: heroku_vpn_connections
https://devcenter.heroku.com/articles/platform-api-reference#vpn-connection-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|String|
|ike_version|Int|
|name|String|
|public_ip|String|
|routable_cidrs|StringArray|
|space_cidr_block|String|
|status|String|
|status_message|String|
|tunnels|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|