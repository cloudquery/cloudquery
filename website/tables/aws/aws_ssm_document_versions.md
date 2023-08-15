# Table: aws_ssm_document_versions

This table shows data for AWS Systems Manager (SSM) Document Versions.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DocumentVersionInfo.html

The composite primary key for this table is (**document_arn**, **document_version**).

## Relations

This table depends on [aws_ssm_documents](aws_ssm_documents).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|document_arn (PK)|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|display_name|`utf8`|
|document_format|`utf8`|
|document_version (PK)|`utf8`|
|is_default_version|`bool`|
|name|`utf8`|
|review_status|`utf8`|
|status|`utf8`|
|status_information|`utf8`|
|version_name|`utf8`|