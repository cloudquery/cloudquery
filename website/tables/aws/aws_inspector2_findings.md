# Table: aws_inspector2_findings

This table shows data for Inspector2 Findings.

https://docs.aws.amazon.com/inspector/v2/APIReference/API_Finding.html
The 'request_account_id' and 'request_region' columns are added to show from where the request was made.

The composite primary key for this table is (**request_account_id**, **request_region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|request_account_id (PK)|String|
|request_region (PK)|String|
|arn (PK)|String|
|aws_account_id|String|
|description|String|
|finding_arn|String|
|first_observed_at|Timestamp|
|last_observed_at|Timestamp|
|remediation|JSON|
|resources|JSON|
|severity|String|
|status|String|
|type|String|
|exploit_available|String|
|exploitability_details|JSON|
|fix_available|String|
|inspector_score|Float|
|inspector_score_details|JSON|
|network_reachability_details|JSON|
|package_vulnerability_details|JSON|
|title|String|
|updated_at|Timestamp|