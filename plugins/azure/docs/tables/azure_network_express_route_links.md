
# Table: azure_network_express_route_links
ExpressRouteLink resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|express_route_port_cq_id|uuid|Unique CloudQuery ID of azure_network_express_route_ports table (FK)|
|id|text|Resource ID.|
|admin_state|text|Administrative state of the physical port.|
|connector_type|text|Physical fiber port type.|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|interface_name|text|Name of Azure router interface.|
|mac_sec_config_cak_secret_identifier|text|Keyvault Secret Identifier URL containing Mac security CAK key.|
|mac_sec_config_cipher|text|Mac security cipher.|
|mac_sec_config_ckn_secret_identifier|text|Keyvault Secret Identifier URL containing Mac security CKN key.|
|mac_sec_config_sci_state|text|Sci mode enabled/disabled.|
|name|text|Resource name.|
|patch_panel_id|text|Mapping between physical port to patch panel port.|
|provisioning_state|text|The provisioning state of the express route link resource.|
|rack_id|text|Mapping of physical patch panel to rack.|
|router_name|text|Name of Azure router associated with physical port.|
