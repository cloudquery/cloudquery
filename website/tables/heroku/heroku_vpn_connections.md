# Table: heroku_vpn_connections

This table shows data for Heroku VPN Connections.

https://devcenter.heroku.com/articles/platform-api-reference#vpn-connection

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|ike_version|int64|
|name|utf8|
|public_ip|utf8|
|routable_cidrs|list<item: utf8, nullable>|
|space_cidr_block|utf8|
|status|utf8|
|status_message|utf8|
|tunnels|json|