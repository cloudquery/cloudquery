
# Table: aws_access_analyzer_analyzer_findings
Contains information about a finding.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|analyzer_id|uuid|Analyzer ID that belongs to aws_access_analyzer_analyzers|
|analyzed_at|timestamp without time zone|The time at which the resource-based policy that generated the finding was analyzed.|
|condition|jsonb|The condition in the analyzed policy statement that resulted in a finding.|
|created_at|timestamp without time zone|The time at which the finding was created.|
|finding_id|text|The ID of the finding.|
|resource_owner_account|text|The AWS account ID that owns the resource.|
|resource_type|text|The type of the resource that the external principal has access to.|
|status|text|The status of the finding.|
|updated_at|timestamp without time zone|The time at which the finding was most recently updated.|
|action|text[]|The action in the analyzed policy statement that an external principal has permission to use.|
|error|text|The error that resulted in an Error finding.|
|is_public|boolean|Indicates whether the finding reports a resource that has a policy that allows public access.|
|principal|jsonb|The external principal that has access to a resource within the zone of trust.|
|resource|text|The resource that the external principal has access to.|
