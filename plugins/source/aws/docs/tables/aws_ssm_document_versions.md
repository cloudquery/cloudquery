# Table: aws_ssm_document_versions

This table shows data for AWS Systems Manager (SSM) Document Versions.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DocumentVersionInfo.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**document_arn**, **document_version**).
## Relations

This table depends on [aws_ssm_documents](aws_ssm_documents.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|document_arn|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|display_name|`utf8`|
|document_format|`utf8`|
|document_version|`utf8`|
|is_default_version|`bool`|
|name|`utf8`|
|review_status|`utf8`|
|status|`utf8`|
|status_information|`utf8`|
|version_name|`utf8`|