# Table: aws_accessanalyzer_analyzer_archive_rules

This table shows data for AWS Identity and Access Management (IAM) Access Analyzer Analyzer Archive Rules.

https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_ArchiveRuleSummary.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_accessanalyzer_analyzers](aws_accessanalyzer_analyzers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|analyzer_arn|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|filter|`json`|
|rule_name|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|