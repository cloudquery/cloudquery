# Table: aws_glue_classifiers

This table shows data for Glue Classifiers.

https://docs.aws.amazon.com/glue/latest/webapi/API_Classifier.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|name|`utf8`|
|csv_classifier|`json`|
|grok_classifier|`json`|
|json_classifier|`json`|
|xml_classifier|`json`|