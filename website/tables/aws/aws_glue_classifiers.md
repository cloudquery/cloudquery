# Table: aws_glue_classifiers

https://docs.aws.amazon.com/glue/latest/webapi/API_Classifier.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|name (PK)|String|
|csv_classifier|JSON|
|grok_classifier|JSON|
|json_classifier|JSON|
|xml_classifier|JSON|