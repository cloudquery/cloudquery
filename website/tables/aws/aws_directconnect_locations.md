# Table: aws_directconnect_locations

This table shows data for AWS Direct Connect Locations.

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_Location.html

The composite primary key for this table is (**account_id**, **region**, **location_code**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|available_mac_sec_port_speeds|`list<item: utf8, nullable>`|
|available_port_speeds|`list<item: utf8, nullable>`|
|available_providers|`list<item: utf8, nullable>`|
|location_code (PK)|`utf8`|
|location_name|`utf8`|