# Table: aws_xray_sampling_rules


The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|tags|JSON|
|created_at|Timestamp|
|modified_at|Timestamp|
|sampling_rule|JSON|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|