# Table: aws_availability_zones

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Region.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|enabled|Bool|
|partition|String|
|region|String|
|group_name|String|
|messages|JSON|
|network_border_group|String|
|opt_in_status|String|
|parent_zone_id|String|
|parent_zone_name|String|
|region_name|String|
|state|String|
|zone_id|String|
|zone_name|String|
|zone_type|String|