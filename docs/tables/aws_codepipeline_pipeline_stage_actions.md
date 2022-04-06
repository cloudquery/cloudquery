
# Table: aws_codepipeline_pipeline_stage_actions
Represents information about an action declaration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|pipeline_stage_cq_id|uuid|Unique CloudQuery ID of aws_codepipeline_pipeline_stages table (FK)|
|category|text|A category defines what kind of action can be taken in the stage, and constrains the provider type for the action|
|owner|text|The creator of the action being called|
|provider|text|The provider of the service being called by the action|
|version|text|A string that describes the action version.  This member is required.|
|name|text|The action declaration's name.  This member is required.|
|configuration|jsonb|The action's configuration|
|input_artifacts|text[]|The name or ID of the artifact consumed by the action, such as a test or build artifact.|
|namespace|text|The variable namespace associated with the action|
|output_artifacts|text[]|The name or ID of the result of the action declaration, such as a test or build artifact.|
|region|text|The action declaration's AWS Region, such as us-east-1.|
|role_arn|text|The ARN of the IAM service role that performs the declared action|
|run_order|integer|The order in which actions are run.|
