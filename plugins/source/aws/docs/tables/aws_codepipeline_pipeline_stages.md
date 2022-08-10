
# Table: aws_codepipeline_pipeline_stages
Represents information about a stage and its definition
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|pipeline_cq_id|uuid|Unique CloudQuery ID of aws_codepipeline_pipelines table (FK)|
|stage_order|bigint|The stage order in the pipeline.|
|name|text|The name of the stage|
|blockers|jsonb|Reserved for future use|
