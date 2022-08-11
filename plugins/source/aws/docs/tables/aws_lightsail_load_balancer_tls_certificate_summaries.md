
# Table: aws_lightsail_load_balancer_tls_certificate_summaries
Provides a summary of SSL/TLS certificate metadata
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_load_balancers table (FK)|
|is_attached|boolean|When true, the SSL/TLS certificate is attached to the Lightsail load balancer|
|name|text|The name of the SSL/TLS certificate|
