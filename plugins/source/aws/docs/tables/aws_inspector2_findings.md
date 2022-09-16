
# Table: aws_inspector2_findings
Details about an Amazon Inspector finding
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|arn|text|The Amazon Resource Name (ARN) of the finding|
|region|text|The AWS Region of the resource|
|account_id|text|The Amazon Web Services account ID associated with the finding|
|description|text|The description of the finding|
|finding_arn|text|The Amazon Resource Number (ARN) of the finding|
|first_observed_at|timestamp without time zone|The date and time that the finding was first observed|
|last_observed_at|timestamp without time zone|The date and time that the finding was last observed|
|remediation_recommendation_text|text|The recommended course of action to remediate the finding|
|remediation_recommendation_url|text|The URL address to the CVE remediation recommendations|
|severity|text|The severity of the finding|
|status|text|The status of the finding|
|type|text|The type of the finding|
|inspector_score|float|The Amazon Inspector score given to the finding|
|inspector_score_details|jsonb|An object that contains details of the Amazon Inspector score|
|network_reachability_details|jsonb|An object that contains the details of a network reachability finding|
|package_vulnerability_details|jsonb|An object that contains the details of a package vulnerability finding|
|title|text|The title of the finding|
|updated_at|timestamp without time zone|The date and time the finding was last updated at|
