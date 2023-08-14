# Table: aws_ssm_documents

This table shows data for AWS Systems Manager (SSM) Documents.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DocumentDescription.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ssm_documents:
  - [aws_ssm_document_versions](aws_ssm_document_versions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|permissions|`json`|
|tags|`json`|
|approved_version|`utf8`|
|attachments_information|`json`|
|author|`utf8`|
|category|`list<item: utf8, nullable>`|
|category_enum|`list<item: utf8, nullable>`|
|created_date|`timestamp[us, tz=UTC]`|
|default_version|`utf8`|
|description|`utf8`|
|display_name|`utf8`|
|document_format|`utf8`|
|document_type|`utf8`|
|document_version|`utf8`|
|hash|`utf8`|
|hash_type|`utf8`|
|latest_version|`utf8`|
|name|`utf8`|
|owner|`utf8`|
|parameters|`json`|
|pending_review_version|`utf8`|
|platform_types|`list<item: utf8, nullable>`|
|requires|`json`|
|review_information|`json`|
|review_status|`utf8`|
|schema_version|`utf8`|
|sha1|`utf8`|
|status|`utf8`|
|status_information|`utf8`|
|target_type|`utf8`|
|version_name|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### SSM documents should not be public

```sql
SELECT
  'SSM documents should not be public' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN 'all' = ANY (ARRAY (SELECT jsonb_array_elements_text(p->'AccountIds')))
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_ssm_documents, jsonb_array_elements(aws_ssm_documents.permissions) AS p
WHERE
  owner IN (SELECT account_id FROM aws_iam_accounts);
```


