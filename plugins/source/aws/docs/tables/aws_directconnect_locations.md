# Table: aws_directconnect_locations

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_Location.html

The composite primary key for this table is (**account_id**, **region**, **location_code**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|available_mac_sec_port_speeds|StringArray|
|available_port_speeds|StringArray|
|available_providers|StringArray|
|location_code (PK)|String|
|location_name|String|