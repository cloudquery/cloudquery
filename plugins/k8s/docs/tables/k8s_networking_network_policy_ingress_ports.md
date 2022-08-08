
# Table: k8s_networking_network_policy_ingress_ports
NetworkPolicyPort describes a port to allow traffic on
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|network_policy_ingress_cq_id|uuid|Unique CloudQuery ID of k8s_networking_network_policy_ingress table (FK)|
|protocol|text|The protocol (TCP, UDP, or SCTP) which traffic must match|
|port_type|bigint||
|port_int_val|integer||
|port_str_val|text||
|end_port|integer|If set, indicates that the range of ports from port to endPort, inclusive, should be allowed by the policy|
