# Table: aws_glue_classifiers

This table shows data for Glue Classifiers.

https://docs.aws.amazon.com/glue/latest/webapi/API_Classifier.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|name (PK)|`utf8`|
|csv_classifier|`json`|
|grok_classifier|`json`|
|json_classifier|`json`|
|xml_classifier|`json`|