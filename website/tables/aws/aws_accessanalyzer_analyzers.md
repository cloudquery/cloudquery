# Table: aws_accessanalyzer_analyzers

This table shows data for AWS Identity and Access Management (IAM) Access Analyzer Analyzers.

https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_AnalyzerSummary.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_accessanalyzer_analyzers:
  - [aws_accessanalyzer_analyzer_archive_rules](aws_accessanalyzer_analyzer_archive_rules)
  - [aws_accessanalyzer_analyzer_findings](aws_accessanalyzer_analyzer_findings)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|created_at|timestamp[us, tz=UTC]|
|name|utf8|
|status|utf8|
|type|utf8|
|last_resource_analyzed|utf8|
|last_resource_analyzed_at|timestamp[us, tz=UTC]|
|status_reason|json|
|tags|json|