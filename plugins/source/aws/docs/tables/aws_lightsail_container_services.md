
# Table: aws_lightsail_container_services
Describes an Amazon Lightsail container service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the container service|
|container_service_name|text|The name of the container service|
|created_at|timestamp without time zone|The timestamp when the container service was created|
|current_deployment_containers|jsonb|An object that describes the configuration for the containers of the deployment|
|current_deployment_created_at|timestamp without time zone|The timestamp when the deployment was created|
|current_deployment_public_endpoint_container_name|text|The name of the container entry of the deployment that the endpoint configuration applies to|
|current_deployment_public_endpoint_container_port|bigint|The port of the specified container to which traffic is forwarded to|
|current_deployment_public_endpoint_health_check|jsonb|An object that describes the health check configuration of the container|
|current_deployment_state|text|The state of the deployment|
|current_deployment_version|bigint|The version number of the deployment|
|is_disabled|boolean|A Boolean value indicating whether the container service is disabled|
|availability_zone|text|The Availability Zone|
|next_deployment_containers|jsonb|An object that describes the configuration for the containers of the deployment|
|next_deployment_created_at|timestamp without time zone|The timestamp when the deployment was created|
|next_deployment_public_endpoint_container_name|text|The name of the container entry of the deployment that the endpoint configuration applies to|
|next_deployment_public_endpoint_container_port|bigint|The port of the specified container to which traffic is forwarded to|
|next_deployment_public_endpoint_health_check|jsonb|An object that describes the health check configuration of the container|
|next_deployment_state|text|The state of the deployment|
|next_deployment_version|bigint|The version number of the deployment|
|power|text|The power specification of the container service|
|power_id|text|The ID of the power of the container service|
|principal_arn|text|The principal ARN of the container service|
|private_domain_name|text|The private domain name of the container service|
|private_registry_access_ecr_image_puller_role_is_active|boolean|A Boolean value that indicates whether the role is activated|
|private_registry_access_ecr_image_puller_role_principal_arn|text|The Amazon Resource Name (ARN) of the role, if it is activated|
|public_domain_names|jsonb|The public domain name of the container service, such as examplecom and wwwexamplecom|
|resource_type|text|The Lightsail resource type of the container service (ie, ContainerService)|
|scale|bigint|The scale specification of the container service|
|state|text|The current state of the container service|
|state_detail_code|text|The state code of the container service|
|state_detail_message|text|A message that provides more information for the state code|
|tags|jsonb|The tag keys and optional values for the resource|
|url|text|The publicly accessible URL of the container service|
