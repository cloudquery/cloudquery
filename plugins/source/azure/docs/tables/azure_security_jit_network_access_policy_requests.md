
# Table: azure_security_jit_network_access_policy_requests
JitNetworkAccessRequest ...
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|jit_network_access_policy_cq_id|uuid|Unique CloudQuery ID of azure_security_jit_network_access_policies table (FK)|
|virtual_machines|text[]||
|start_time_utc|timestamp without time zone||
|requestor|text|The identity of the person who made the request|
|justification|text|The justification for making the initiate request|
