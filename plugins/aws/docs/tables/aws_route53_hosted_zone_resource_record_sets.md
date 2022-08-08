
# Table: aws_route53_hosted_zone_resource_record_sets
Information about the resource record set to create or delete.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hosted_zone_cq_id|uuid|Unique CloudQuery ID of aws_route53_hosted_zones table (FK)|
|resource_records|text[]||
|name|text|For ChangeResourceRecordSets requests, the name of the record that you want to create, update, or delete.|
|type|text|The DNS record type.|
|dns_name|text|Alias resource record sets only: The value that you specify depends on where you want to route queries: Amazon API Gateway custom regional APIs and edge-optimized APIs Specify the applicable domain name for your API.|
|evaluate_target_health|boolean|Applies only to alias, failover alias, geolocation alias, latency alias, and weighted alias resource record sets: When EvaluateTargetHealth is true, an alias resource record set inherits the health of the referenced AWS resource, such as an ELB load balancer or another resource record set in the hosted zone.|
|failover|text|Failover resource record sets only: To configure failover, you add the Failover element to two resource record sets.|
|geo_location_continent_code|text|The two-letter code for the continent.|
|geo_location_country_code|text|For geolocation resource record sets, the two-letter code for a country.|
|geo_location_subdivision_code|text|For geolocation resource record sets, the two-letter code for a state of the United States.|
|health_check_id|text|If you want Amazon Route 53 to return this resource record set in response to a DNS query only when the status of a health check is healthy, include the HealthCheckId element and specify the ID of the applicable health check.|
|multi_value_answer|boolean|Multivalue answer resource record sets only: To route traffic approximately randomly to multiple resources, such as web servers, create one multivalue answer record for each resource and specify true for MultiValueAnswer.|
|region|text|Latency-based resource record sets only: The Amazon EC2 Region where you created the resource that this resource record set refers to.|
|set_identifier|text|Resource record sets that have a routing policy other than simple: An identifier that differentiates among multiple resource record sets that have the same combination of name and type, such as multiple weighted resource record sets named acme.|
|ttl|bigint|The resource record cache time to live (TTL), in seconds.|
|traffic_policy_instance_id|text|When you create a traffic policy instance, Amazon Route 53 automatically creates a resource record set.|
|weight|bigint|Weighted resource record sets only: Among resource record sets that have the same combination of DNS name and type, a value that determines the proportion of DNS queries that Amazon Route 53 responds to using the current resource record set.|
