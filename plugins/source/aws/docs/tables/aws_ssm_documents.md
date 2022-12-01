# Table: aws_ssm_documents

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DocumentDescription.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|permissions|JSON|
|approved_version|String|
|attachments_information|JSON|
|author|String|
|category|StringArray|
|category_enum|StringArray|
|created_date|Timestamp|
|default_version|String|
|description|String|
|display_name|String|
|document_format|String|
|document_type|String|
|document_version|String|
|hash|String|
|hash_type|String|
|latest_version|String|
|name|String|
|owner|String|
|parameters|JSON|
|pending_review_version|String|
|platform_types|StringArray|
|requires|JSON|
|review_information|JSON|
|review_status|String|
|schema_version|String|
|sha1|String|
|status|String|
|status_information|String|
|tags|JSON|
|target_type|String|
|version_name|String|