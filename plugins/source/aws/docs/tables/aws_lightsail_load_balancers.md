# Table: aws_lightsail_load_balancers

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_LoadBalancer.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_lightsail_load_balancers:
  - [aws_lightsail_load_balancer_tls_certificates](aws_lightsail_load_balancer_tls_certificates.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|configuration_options|JSON|
|created_at|Timestamp|
|dns_name|String|
|health_check_path|String|
|https_redirection_enabled|Bool|
|instance_health_summary|JSON|
|instance_port|Int|
|ip_address_type|String|
|location|JSON|
|name|String|
|protocol|String|
|public_ports|IntArray|
|resource_type|String|
|state|String|
|support_code|String|
|tags|JSON|
|tls_certificate_summaries|JSON|
|tls_policy_name|String|