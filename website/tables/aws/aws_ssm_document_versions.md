# Table: aws_ssm_document_versions

This table shows data for Ssm Document Versions.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DocumentVersionInfo.html

The composite primary key for this table is (**document_arn**, **document_version**).

## Relations

This table depends on [aws_ssm_documents](aws_ssm_documents).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|document_arn (PK)|String|
|created_date|Timestamp|
|display_name|String|
|document_format|String|
|document_version (PK)|String|
|is_default_version|Bool|
|name|String|
|review_status|String|
|status|String|
|status_information|String|
|version_name|String|