
# Table: aws_elasticbeanstalk_environments
Describes the properties of an environment.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|abortable_operation_in_progress|boolean|Indicates if there is an in-progress environment configuration update or application version deployment that you can cancel.|
|application_name|text|The name of the application associated with this environment.|
|cname|text|The URL to the CNAME for this environment.|
|date_created|timestamp without time zone|The creation date for this environment.|
|date_updated|timestamp without time zone|The last modified date for this environment.|
|description|text|Describes this environment.|
|endpoint_url|text|For load-balanced, autoscaling environments, the URL to the LoadBalancer.|
|environment_arn|text|The environment's Amazon Resource Name (ARN), which can be used in other API requests that require an ARN.|
|environment_id|text|The ID of this environment.|
|environment_name|text|The name of this environment.|
|health|text|Describes the health status of the environment.|
|health_status|text|Returns the health status of the application running in your environment.|
|operations_role|text|The Amazon Resource Name (ARN) of the environment's operations role.|
|platform_arn|text|The ARN of the platform version.|
|load_balancer_domain|text|The domain name of the LoadBalancer.|
|listeners|jsonb|A list of Listeners used by the LoadBalancer.|
|load_balancer_name|text|The name of the LoadBalancer.|
|solution_stack_name|text|The name of the SolutionStack deployed with this environment.|
|status|text|The current operational status of the environment:  * Launching: Environment is in the process of initial deployment.|
|template_name|text|The name of the configuration template used to originally launch this environment.|
|tier_name|text|The name of this environment tier.|
|tier_type|text|The type of this environment tier.|
|tier_version|text|The version of this environment tier.|
|version_label|text|The application version deployed in this environment.|
