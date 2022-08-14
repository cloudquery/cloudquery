
# Table: aws_access_analyzer_analyzer_archive_rules
Contains information about an archive rule
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|analyzer_cq_id|uuid|Unique CloudQuery ID of aws_access_analyzer_analyzers table (FK)|
|created_at|timestamp without time zone|The time at which the archive rule was created|
|filter|jsonb|A filter used to define the archive rule|
|rule_name|text|The name of the archive rule|
|updated_at|timestamp without time zone|The time at which the archive rule was last updated|
