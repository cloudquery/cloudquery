
# Table: gcp_dns_managed_zones
A zone is a subtree of the DNS namespace under one administrative responsibility A ManagedZone is a resource that represents a DNS zone hosted by the Cloud DNS service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|creation_time|text|The time that this resource was created on the server This is in RFC3339 text format Output only|
|description|text|A mutable string of at most 1024 characters associated with this resource for the user's convenience Has no effect on the managed zone's function|
|dns_name|text|The DNS name of this managed zone, for instance "examplecom"|
|dnssec_config_kind|text||
|dnssec_config_non_existence|text|Specifies the mechanism for authenticated denial-of-existence responses Can only be changed while the state is OFF|
|dnssec_config_state|text|Specifies whether DNSSEC is enabled, and what mode it is in  Possible values:   "off" - DNSSEC is disabled; the zone is not signed   "on" - DNSSEC is enabled; the zone is signed and fully managed   "transfer" - DNSSEC is enabled, but in a "transfer" mode|
|forwarding_config_kind|text||
|managed_zone_id|text||
|kind|text|The resource type|
|labels|jsonb|User assigned labels for this resource|
|name|text|User assigned name for this resource|
|name_server_set|text|specifies the NameServerSet for this ManagedZone|
|name_servers|text[]||
|peering_config_kind|text||
|peering_config_target_network_deactivate_time|text|The time at which the zone was deactivated, in RFC 3339 date-time format An empty string indicates that the peering connection is active The producer network can deactivate a zone The zone is automatically deactivated if the producer network that the zone targeted is deleted Output only|
|peering_config_target_network_kind|text||
|peering_config_target_network_network_url|text|The fully qualified URL of the VPC network to forward queries to This should be formatted like https://wwwgoogleapis|
|private_visibility_config_kind|text||
|reverse_lookup_config_kind|text||
|service_directory_config_kind|text||
|service_directory_config_namespace_deletion_time|text|The time that the namespace backing this zone was deleted; an empty string if it still exists This is in RFC3339 text format Output only|
|service_directory_config_namespace_kind|text||
|service_directory_config_namespace_namespace_url|text|The fully qualified URL of the namespace associated with the zone Format must be https://servicedirectorygoogleapis|
|visibility|text|The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources|
