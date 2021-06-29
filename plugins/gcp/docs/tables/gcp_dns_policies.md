
# Table: gcp_dns_policies
A policy is a collection of DNS rules applied to one or more Virtual Private Cloud resources
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|alternative_name_server_config_kind|text|alternative name server type|
|description|text|A mutable string of at most 1024 characters associated with this resource for the user's convenience Has no effect on the policy's function|
|enable_inbound_forwarding|boolean|Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy|
|enable_logging|boolean|Controls whether logging is enabled for the networks bound to this policy Defaults to no logging if not set|
|policy_id|text|Unique identifier for the resource; defined by the server (output only)|
|kind|text|The resource type|
|name|text|User-assigned name for this policy|
