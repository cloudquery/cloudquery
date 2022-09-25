# Table: aws_backup_region_settings


The composite primary key for this table is (**account_id**, **region**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id (PK)|String|
|region (PK)|String|
|resource_type_management_preference|JSON|
|resource_type_opt_in_preference|JSON|
|result_metadata|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|