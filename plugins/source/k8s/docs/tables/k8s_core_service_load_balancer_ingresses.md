
# Table: k8s_core_service_load_balancer_ingresses
LoadBalancerIngress represents the status of a load-balancer ingress point: traffic intended for the service should be sent to an ingress point.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of k8s_core_services table (FK)|
|ip|text|IP is set for load-balancer ingress points that are IP based.|
|hostname|text|A set for load-balancer ingress points that are DNS based|
