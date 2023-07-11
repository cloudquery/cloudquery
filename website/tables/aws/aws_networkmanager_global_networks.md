# Table: aws_networkmanager_global_networks

This table shows data for Networkmanager Global Networks.

https://docs.aws.amazon.com/networkmanager/latest/APIReference/API_GlobalNetwork.html
The  'request_region' column is added to show region of where the request was made from.

The composite primary key for this table is (**request_region**, **arn**).

## Relations

The following tables depend on aws_networkmanager_global_networks:
  - [aws_networkmanager_links](aws_networkmanager_links)
  - [aws_networkmanager_sites](aws_networkmanager_sites)
  - [aws_networkmanager_transit_gateway_registrations](aws_networkmanager_transit_gateway_registrations)

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
|global_network_arn|`utf8`|
|global_network_id|`utf8`|
|state|`utf8`|