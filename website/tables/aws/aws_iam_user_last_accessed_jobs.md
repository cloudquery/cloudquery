# Table: aws_iam_user_last_accessed_jobs

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GenerateServiceLastAccessedDetails.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_iam_users](aws_iam_users).

The following tables depend on aws_iam_user_last_accessed_jobs:
  - [aws_iam_user_last_accessed_details](aws_iam_user_last_accessed_details)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|arn (PK)|String|
|job_id|String|