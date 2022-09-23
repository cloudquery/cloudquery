# Table: aws_elasticbeanstalk_applications


The composite primary key for this table is (**arn**, **date_created**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|date_created (PK)|Timestamp|
|application_name|String|
|configuration_templates|StringArray|
|date_updated|Timestamp|
|description|String|
|resource_lifecycle_config|JSON|
|versions|StringArray|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|