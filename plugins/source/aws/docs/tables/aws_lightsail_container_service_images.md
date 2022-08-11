
# Table: aws_lightsail_container_service_images
Describes a container image that is registered to an Amazon Lightsail container service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|container_service_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_container_services table (FK)|
|created_at|timestamp without time zone|The timestamp when the container image was created|
|digest|text|The digest of the container image|
|image|text|The name of the container image|
