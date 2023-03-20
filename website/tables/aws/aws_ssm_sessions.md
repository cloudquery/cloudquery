# Table: aws_ssm_sessions

This table shows data for AWS Systems Manager (SSM) Sessions.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_Session.html. 
Only Active sessions are fetched.

The composite primary key for this table is (**account_id**, **region**, **session_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|details|String|
|document_name|String|
|end_date|Timestamp|
|max_session_duration|String|
|output_url|JSON|
|owner|String|
|reason|String|
|session_id (PK)|String|
|start_date|Timestamp|
|status|String|
|target|String|