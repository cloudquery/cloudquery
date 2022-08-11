
# Table: aws_redshift_cluster_endpoint_vpc_endpoints
The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_redshift_clusters table (FK)|
|vpc_endpoint_id|text|The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.|
|vpc_id|text|The VPC identifier that the endpoint is associated.|
