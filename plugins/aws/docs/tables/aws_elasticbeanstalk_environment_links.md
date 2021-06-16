
# Table: aws_elasticbeanstalk_environment_links
A link to another environment, defined in the environment's manifest.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|environment_id|uuid|Unique ID of aws_elasticbeanstalk_environments table (FK)|
|environment_name|text|The name of the linked environment (the dependency).|
|link_name|text|The name of the link.|
