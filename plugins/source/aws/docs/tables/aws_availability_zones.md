# Table: aws_availability_zones

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Region.html

The composite primary key for this table is (**account_id**, **region_name**, **zone_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|enabled|Bool|
|partition|String|
|region|String|
|group_name|String|
|messages|JSON|
|network_border_group|String|
|opt_in_status|String|
|parent_zone_id|String|
|parent_zone_name|String|
|region_name (PK)|String|
|state|String|
|zone_id (PK)|String|
|zone_name|String|
|zone_type|String|