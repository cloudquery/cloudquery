
# Table: aws_access_analyzer_analyzer_finding_sources
The source of the finding.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|analyzer_finding_cq_id|uuid|Unique CloudQuery ID of aws_access_analyzer_analyzer_findings table (FK)|
|type|text|Indicates the type of access that generated the finding|
|detail_access_point_arn|text|The ARN of the access point that generated the finding|
