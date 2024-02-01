# Table: aws_networkmanager_links

This table shows data for Networkmanager Links.

https://docs.aws.amazon.com/networkmanager/latest/APIReference/API_Link.html
The  'request_region' column is added to show region of where the request was made from.

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_region**, **arn**, **global_network_id**).
## Relations

This table depends on [aws_networkmanager_global_networks](aws_networkmanager_global_networks.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|request_region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|bandwidth|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|global_network_id|`utf8`|
|link_arn|`utf8`|
|link_id|`utf8`|
|provider|`utf8`|
|site_id|`utf8`|
|state|`utf8`|
|type|`utf8`|