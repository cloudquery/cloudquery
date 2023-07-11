# Table: aws_networkmanager_sites

This table shows data for Networkmanager Sites.

https://docs.aws.amazon.com/networkmanager/latest/APIReference/API_Site.html
The  'request_region' column is added to show region of where the request was made from.

The composite primary key for this table is (**request_region**, **arn**, **global_network_id**).

## Relations

This table depends on [aws_networkmanager_global_networks](aws_networkmanager_global_networks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|request_region (PK)|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|global_network_id (PK)|`utf8`|
|location|`json`|
|site_arn|`utf8`|
|site_id|`utf8`|
|state|`utf8`|