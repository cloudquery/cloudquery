# Table: aws_appconfig_deployment_strategies

This table shows data for AWS AppConfig Deployment Strategies.

https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_DeploymentStrategy.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|deployment_duration_in_minutes|`int64`|
|description|`utf8`|
|final_bake_time_in_minutes|`int64`|
|growth_factor|`float64`|
|growth_type|`utf8`|
|id|`utf8`|
|name|`utf8`|
|replicate_to|`utf8`|