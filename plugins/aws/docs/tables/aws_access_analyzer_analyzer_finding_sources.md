
# Table: aws_access_analyzer_analyzer_finding_sources
The source of the finding.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|analyzer_finding_id|uuid|AnalyzerFinding ID that belongs to aws_access_analyzer_analyzer_findings|
|type|text|Indicates the type of access that generated the finding.|
|detail_access_point_arn|text|The ARN of the access point that generated the finding.|
