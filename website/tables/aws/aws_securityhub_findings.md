# Table: aws_securityhub_findings

This table shows data for AWS Security Hub Findings.

https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_GetFindings.html
The `request_account_id` and `request_region` columns are added to show the account and region of where the request was made from.
This is useful when multi region and account aggregation is enabled.

The composite primary key for this table is (**request_account_id**, **request_region**, **aws_account_id**, **created_at**, **description**, **generator_id**, **id**, **product_arn**, **schema_version**, **title**, **updated_at**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|aws_account_id (PK)|`utf8`|
|created_at (PK)|`timestamp[us, tz=UTC]`|
|description (PK)|`utf8`|
|generator_id (PK)|`utf8`|
|id (PK)|`utf8`|
|product_arn (PK)|`utf8`|
|resources|`json`|
|schema_version (PK)|`utf8`|
|title (PK)|`utf8`|
|updated_at (PK)|`timestamp[us, tz=UTC]`|
|action|`json`|
|company_name|`utf8`|
|compliance|`json`|
|confidence|`int64`|
|criticality|`int64`|
|finding_provider_fields|`json`|
|first_observed_at|`timestamp[us, tz=UTC]`|
|last_observed_at|`timestamp[us, tz=UTC]`|
|malware|`json`|
|network|`json`|
|network_path|`json`|
|note|`json`|
|patch_summary|`json`|
|process|`json`|
|product_fields|`json`|
|product_name|`utf8`|
|record_state|`utf8`|
|region (PK)|`utf8`|
|related_findings|`json`|
|remediation|`json`|
|sample|`bool`|
|severity|`json`|
|source_url|`utf8`|
|threat_intel_indicators|`json`|
|threats|`json`|
|types|`list<item: utf8, nullable>`|
|user_defined_fields|`json`|
|verification_state|`utf8`|
|vulnerabilities|`json`|
|workflow|`json`|
|workflow_state|`utf8`|