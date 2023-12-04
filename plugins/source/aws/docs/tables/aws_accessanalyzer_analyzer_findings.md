# Table: aws_accessanalyzer_analyzer_findings

This table shows data for AWS Identity and Access Management (IAM) Access Analyzer Analyzer Findings.

https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_FindingSummary.html

The composite primary key for this table is (**analyzer_arn**, **id**).

## Relations

This table depends on [aws_accessanalyzer_analyzers](aws_accessanalyzer_analyzers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|analyzer_arn (PK)|`utf8`|
|analyzed_at|`timestamp[us, tz=UTC]`|
|condition|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|id (PK)|`utf8`|
|resource_owner_account|`utf8`|
|resource_type|`utf8`|
|status|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|action|`list<item: utf8, nullable>`|
|error|`utf8`|
|is_public|`bool`|
|principal|`json`|
|resource|`utf8`|
|sources|`json`|