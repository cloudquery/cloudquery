
# Table: aws_codebuild_project_environment_variables
Information about an environment variable for a build project or a build.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_cq_id|uuid|Unique CloudQuery ID of aws_codebuild_projects table (FK)|
|name|text|The name or key of the environment variable.|
|value|text|The value of the environment variable|
|type|text|The type of environment variable|
