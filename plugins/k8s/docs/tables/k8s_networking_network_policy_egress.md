
# Table: k8s_networking_network_policy_egress
NetworkPolicyEgressRule describes a particular set of traffic that is allowed out of pods matched by a NetworkPolicySpec's podSelector
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|network_policy_cq_id|uuid|Unique CloudQuery ID of k8s_networking_network_policies table (FK)|
|network_policy_uid|text|Unique internal ID of Network Policy resource|
