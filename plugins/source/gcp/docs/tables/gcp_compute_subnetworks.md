
# Table: gcp_compute_subnetworks
Represents a Subnetwork resource  A subnetwork (also known as a subnet) is a logical partition of a Virtual Private Cloud network with one primary IP range and zero or more secondary IP ranges
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|creation_timestamp|text|Creation timestamp in RFC3339 text format|
|description|text|An optional description of this resource Provide this property when you create the resource This field can be set only at resource creation time|
|enable_flow_logs|boolean|Whether to enable flow logging for this subnetwork If this field is not explicitly set, it will not appear in get listings If not set the default behavior is to disable flow logging This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER|
|fingerprint|text|Fingerprint of this resource A hash of the contents stored in this object This field is used in optimistic locking This field will be ignored when inserting a Subnetwork An up-to-date fingerprint must be provided in order to update the Subnetwork, otherwise the request will fail with error 412 conditionNotMet  To see the latest fingerprint, make a get() request to retrieve a Subnetwork|
|gateway_address|text|The gateway address for default routes to reach destination addresses outside this subnetwork|
|id|text|The unique identifier for the resource This identifier is defined by the server|
|ip_cidr_range|text|The range of internal addresses that are owned by this subnetwork Provide this property when you create the subnetwork For example, 10000/8 or 1006400/10 Ranges must be unique and non-overlapping within a network Only IPv4 is supported This field is set at resource creation time The range can be any range listed in the Valid ranges list The range can be expanded after creation using expandIpCidrRange|
|ipv6_cidr_range|text|The range of internal IPv6 addresses that are owned by this subnetwork|
|kind|text|Type of the resource Always compute#subnetwork for Subnetwork resources|
|log_config_aggregation_interval|text|Can only be specified if VPC flow logging for this subnetwork is enabled Toggles the aggregation interval for collecting flow logs Increasing the interval time will reduce the amount of generated flow logs for long lasting connections Default is an interval of 5 seconds per connection|
|log_config_enable|boolean|Whether to enable flow logging for this subnetwork If this field is not explicitly set, it will not appear in get listings If not set the default behavior is to disable flow logging|
|log_config_filter_expr|text|Can only be specified if VPC flow logs for this subnetwork is enabled Export filter used to define which VPC flow logs should be logged|
|log_config_flow_sampling|float|Can only be specified if VPC flow logging for this subnetwork is enabled The value of the field must be in [0, 1] Set the sampling rate of VPC flow logs within the subnetwork where 10 means all collected logs are reported and 00 means no logs are reported Default is 05, which means half of all collected logs are reported|
|log_config_metadata|text|Can only be specified if VPC flow logs for this subnetwork is enabled Configures whether all, none or a subset of metadata fields should be added to the reported VPC flow logs Default is EXCLUDE_ALL_METADATA|
|log_config_metadata_fields|text[]|Can only be specified if VPC flow logs for this subnetwork is enabled and "metadata" was set to CUSTOM_METADATA|
|name|text|Name of the resource Provided by the client when the resource is created|
|network|text|The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork Only networks that are in the distributed mode can have subnetworks This field can be set only at resource creation time|
|private_ip_google_access|boolean|Whether the VMs in this subnet can access Google services without assigned external IP addresses This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess|
|private_ipv6_google_access|text|The private IPv6 google access type for the VMs in this subnet This is an expanded field of enablePrivateV6Access If both fields are set, privateIpv6GoogleAccess will take priority  This field can be both set at resource creation time and updated using patch|
|purpose|text|The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing If unspecified, the purpose defaults to PRIVATE_RFC_1918 The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCE|
|region|text|URL of the region where the Subnetwork resides This field can be set only at resource creation time|
|role|text|The role of subnetwork Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER The value can be set to ACTIVE or BACKUP An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining This field can be updated with a patch request|
|self_link|text|Server-defined URL for the resource|
|state|text|The state of the subnetwork, which can be one of the following values: READY: Subnetwork is created and ready to use DRAINING: only applicable to subnetworks that have the purpose set to INTERNAL_HTTPS_LOAD_BALANCER and indicates that connections to the load balancer are being drained|
