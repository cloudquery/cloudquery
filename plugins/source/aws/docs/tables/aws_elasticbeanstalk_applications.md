# Table: aws_elasticbeanstalk_applications

This table shows data for AWS Elastic Beanstalk Applications.

https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationDescription.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**arn**, **date_created**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|date_created|`timestamp[us, tz=UTC]`|
|tags|`json`|
|application_arn|`utf8`|
|application_name|`utf8`|
|configuration_templates|`list<item: utf8, nullable>`|
|date_updated|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|resource_lifecycle_config|`json`|
|versions|`list<item: utf8, nullable>`|