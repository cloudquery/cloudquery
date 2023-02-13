# Table: aws_directconnect_locations

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_Location.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|available_mac_sec_port_speeds|StringArray|
|available_port_speeds|StringArray|
|available_providers|StringArray|
|location_code|String|
|location_name|String|