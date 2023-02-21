# Table: aws_iam_instance_profiles

https://docs.aws.amazon.com/IAM/latest/APIReference/API_InstanceProfile.html

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|id (PK)|String|
|tags|JSON|
|arn|String|
|create_date|Timestamp|
|instance_profile_id|String|
|instance_profile_name|String|
|path|String|
|roles|JSON|