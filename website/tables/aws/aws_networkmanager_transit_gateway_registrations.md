# Table: aws_networkmanager_transit_gateway_registrations

This table shows data for Networkmanager Transit Gateway Registrations.

https://docs.aws.amazon.com/networkmanager/latest/APIReference/API_TransitGatewayRegistration.html
The  'request_region' column is added to show region of where the request was made from.

The composite primary key for this table is (**request_region**, **global_network_id**, **transit_gateway_arn**).

## Relations

This table depends on [aws_networkmanager_global_networks](aws_networkmanager_global_networks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|request_region (PK)|`utf8`|
|global_network_id (PK)|`utf8`|
|state|`json`|
|transit_gateway_arn (PK)|`utf8`|