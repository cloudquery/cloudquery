
# Table: azure_security_jit_network_access_policy_virtual_machines
JitNetworkAccessPolicyVirtualMachine ...
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|jit_network_access_policy_cq_id|uuid|Unique CloudQuery ID of azure_security_jit_network_access_policies table (FK)|
|id|text|Resource ID of the virtual machine that is linked to this policy|
|ports|jsonb|Port configurations for the virtual machine|
|public_ip_address|inet|Public IP address of the Azure Firewall that is linked to this policy, if applicable|
