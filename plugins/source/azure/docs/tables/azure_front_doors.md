
# Table: azure_front_doors
Front Door represents a collection of backend endpoints to route traffic to along with rules that specify how traffic is sent there.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription ID|
|resource_state|text|Resource state of the Front Door|
|provisioning_state|text|Provisioning state of the Front Door|
|cname|text|The host that each frontend endpoint must CNAME to|
|frontdoor_id|text|The ID of the Front Door|
|friendly_name|text|A friendly name for the Front Door|
|enforce_certificate_name_check|text|Whether to enforce certificate name check on HTTPS requests to all backend pools|
|send_recv_timeout_seconds|integer|Send and receive timeout on forwarding request to the backend|
|enabled_state|text|Operational status of the Front Door load balancer|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
|location|text|Resource location|
|tags|jsonb|Resource tags|
