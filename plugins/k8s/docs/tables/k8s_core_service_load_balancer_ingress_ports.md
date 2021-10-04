
# Table: k8s_core_service_load_balancer_ingress_ports

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_load_balancer_ingress_cq_id|uuid|Unique CloudQuery ID of k8s_core_service_load_balancer_ingresses table (FK)|
|port|integer|Port is the port number of the service port of which status is recorded here|
|protocol|text|Protocol is the protocol of the service port of which status is recorded here.|
|error|text|Error is to record the problem with the service port.|
