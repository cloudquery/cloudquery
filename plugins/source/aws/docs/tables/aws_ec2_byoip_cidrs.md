# Table: aws_ec2_byoip_cidrs


The composite primary key for this table is (**account_id**, **region**, **cidr**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id (PK)|String|
|region (PK)|String|
|cidr (PK)|String|
|description|String|
|state|String|
|status_message|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|