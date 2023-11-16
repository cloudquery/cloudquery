# Table: aws_inspector2_findings

This table shows data for Inspector2 Findings.

https://docs.aws.amazon.com/inspector/v2/APIReference/API_Finding.html
The 'request_account_id' and 'request_region' columns are added to show from where the request was made.

The composite primary key for this table is (**request_account_id**, **request_region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|arn (PK)|`utf8`|
|aws_account_id|`utf8`|
|description|`utf8`|
|finding_arn|`utf8`|
|first_observed_at|`timestamp[us, tz=UTC]`|
|last_observed_at|`timestamp[us, tz=UTC]`|
|remediation|`json`|
|resources|`json`|
|severity|`utf8`|
|status|`utf8`|
|type|`utf8`|
|code_vulnerability_details|`json`|
|epss|`json`|
|exploit_available|`utf8`|
|exploitability_details|`json`|
|fix_available|`utf8`|
|inspector_score|`float64`|
|inspector_score_details|`json`|
|network_reachability_details|`json`|
|package_vulnerability_details|`json`|
|title|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|