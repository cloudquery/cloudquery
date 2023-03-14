# Table: aws_accessanalyzer_analyzer_archive_rules

This table shows data for AWS Identity and Access Management (IAM) Access Analyzer Analyzer Archive Rules.

https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_ArchiveRuleSummary.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_accessanalyzer_analyzers](aws_accessanalyzer_analyzers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|analyzer_arn|String|
|created_at|Timestamp|
|filter|JSON|
|rule_name|String|
|updated_at|Timestamp|