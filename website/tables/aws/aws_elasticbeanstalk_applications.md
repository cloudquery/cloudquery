# Table: aws_elasticbeanstalk_applications

This table shows data for AWS Elastic Beanstalk Applications.

https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationDescription.html

The composite primary key for this table is (**arn**, **date_created**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|date_created (PK)|`timestamp[us, tz=UTC]`|
|tags|`json`|
|application_arn|`utf8`|
|application_name|`utf8`|
|configuration_templates|`list<item: utf8, nullable>`|
|date_updated|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|resource_lifecycle_config|`json`|
|versions|`list<item: utf8, nullable>`|