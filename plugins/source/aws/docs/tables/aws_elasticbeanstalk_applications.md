# Table: aws_elasticbeanstalk_applications

https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationDescription.html

The composite primary key for this table is (**arn**, **date_created**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
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