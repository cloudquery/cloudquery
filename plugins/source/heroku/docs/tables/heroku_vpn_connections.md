# Table: heroku_vpn_connections

https://devcenter.heroku.com/articles/platform-api-reference#vpn-connection

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|ike_version|Int|
|name|String|
|public_ip|String|
|routable_cidrs|StringArray|
|space_cidr_block|String|
|status|String|
|status_message|String|
|tunnels|JSON|