# Table: aws_accessanalyzer_analyzers



The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_accessanalyzer_analyzers`:
  - [`aws_accessanalyzer_analyzer_findings`](aws_accessanalyzer_analyzer_findings.md)
  - [`aws_accessanalyzer_analyzer_archive_rules`](aws_accessanalyzer_analyzer_archive_rules.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|arn (PK)|String|
|created_at|Timestamp|
|name|String|
|status|String|
|type|String|
|last_resource_analyzed|String|
|last_resource_analyzed_at|Timestamp|
|status_reason|JSON|
|tags|JSON|