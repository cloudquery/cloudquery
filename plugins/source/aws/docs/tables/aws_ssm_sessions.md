# Table: aws_ssm_sessions

This table shows data for AWS Systems Manager (SSM) Sessions.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_Session.html. 
Only Active sessions are fetched.

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **session_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|details|`utf8`|
|document_name|`utf8`|
|end_date|`timestamp[us, tz=UTC]`|
|max_session_duration|`utf8`|
|output_url|`json`|
|owner|`utf8`|
|reason|`utf8`|
|session_id|`utf8`|
|start_date|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|target|`utf8`|