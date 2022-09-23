# Table: aws_glue_security_configurations


The composite primary key for this table is (**account_id**, **region**, **name**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id (PK)|String|
|region (PK)|String|
|name (PK)|String|
|created_time_stamp|Timestamp|
|encryption_configuration|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|