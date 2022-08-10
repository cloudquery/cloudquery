
# Table: aws_lightsail_container_service_deployments
Describes a container deployment configuration of an Amazon Lightsail container service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|container_service_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_container_services table (FK)|
|containers|jsonb|An object that describes the configuration for the containers of the deployment|
|created_at|timestamp without time zone|The timestamp when the deployment was created|
|public_endpoint_container_name|text|The name of the container entry of the deployment that the endpoint configuration applies to|
|public_endpoint_container_port|integer|The port of the specified container to which traffic is forwarded to|
|public_endpoint_health_check|jsonb|An object that describes the health check configuration of the container|
|state|text|The state of the deployment|
|version|integer|The version number of the deployment|
