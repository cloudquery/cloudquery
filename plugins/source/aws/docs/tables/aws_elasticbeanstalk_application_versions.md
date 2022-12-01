# Table: aws_elasticbeanstalk_application_versions

https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationVersionDescription.html

The primary key for this table is **arn**.



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
|application_name|String|
|build_arn|String|
|date_created|Timestamp|
|date_updated|Timestamp|
|description|String|
|source_build_information|JSON|
|source_bundle|JSON|
|status|String|
|version_label|String|