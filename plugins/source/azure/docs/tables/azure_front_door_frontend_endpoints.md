
# Table: azure_front_door_frontend_endpoints
Frontend endpoints available to routing rules
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|front_door_cq_id|uuid|Unique CloudQuery ID of azure_front_doors table (FK)|
|resource_state|text|Resource state|
|custom_https_provisioning_state|text|Provisioning status of custom https of the frontend endpoint|
|custom_https_provisioning_substate|text|Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step|
|certificate_source|text|Defines the source of the SSL certificate|
|protocol_type|text|Defines the TLS extension protocol that is used for secure delivery|
|minimum_tls_version|text|The minimum TLS version required from the clients to establish an SSL handshake with Front Door|
|vault_id|text|The Key Vault containing the SSL certificate|
|secret_name|text|The name of the Key Vault secret representing the full certificate PFX|
|secret_version|text|The version of the Key Vault secret representing the full certificate PFX|
|certificate_type|text|The type of the certificate used for secure connections to the frontend endpoint|
|host_name|text|The host name of the frontend endpoint|
|session_affinity_enabled_state|text|Whether session affinity is allowed on the host|
|session_affinity_ttl_seconds|integer|The TTL to use in seconds for session affinity, if applicable|
|web_application_firewall_policy_link_id|text|Defines the Web Application Firewall policy for each host (if applicable)|
|name|text|Resource name|
|type|text|Resource type|
|id|text|Resource ID|
