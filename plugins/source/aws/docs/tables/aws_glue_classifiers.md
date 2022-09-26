# Table: aws_glue_classifiers


The composite primary key for this table is (**account_id**, **region**, **name**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id (PK)|String|
|region (PK)|String|
|name (PK)|String|
|csv_classifier|JSON|
|grok_classifier|JSON|
|json_classifier|JSON|
|xml_classifier|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|