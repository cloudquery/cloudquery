
# Table: gcp_compute_forwarding_rules
Represents a Forwarding Rule resource  Forwarding rule resources in GCP can be either regional or global.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|ip_address|text|IP address that this forwarding rule serves When a client sends traffic to this IP address, the forwarding rule directs the traffic to the target that you specify in the forwarding rule|
|ip_protocol|text|The IP protocol to which this rule applies  For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP and ICMP  The valid IP protocols are different for different load balancing products: - Internal TCP/UDP Load Balancing: The load balancing scheme is INTERNAL, and one of TCP, UDP or ALL is valid - Traffic Director: The load balancing scheme is INTERNAL_SELF_MANAGED, and only TCP is valid - Internal HTTP(S) Load Balancing: The load balancing scheme is INTERNAL_MANAGED, and only TCP is valid - HTTP(S), SSL Proxy, and TCP Proxy Load Balancing: The load balancing scheme is EXTERNAL and only TCP is valid - Network Load Balancing: The load balancing scheme is EXTERNAL, and one of TCP or UDP is valid|
|all_ports|boolean|This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance This field cannot be used with port or portRange fields  When the load balancing scheme is INTERNAL and protocol is TCP/UDP, specify this field to allow packets addressed to any ports will be forwarded to the backends configured with this forwarding rule|
|allow_global_access|boolean|This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance If the field is set to TRUE, clients can access ILB from all regions Otherwise only allows access from clients in the same region as the internal load balancer|
|backend_service|text|Identifies the backend service to which the forwarding rule sends traffic Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types|
|creation_timestamp|text|Creation timestamp in RFC3339 text format|
|description|text|An optional description of this resource Provide this property when you create the resource|
|fingerprint|text|Fingerprint of this resource A hash of the contents stored in this object|
|id|text|The unique identifier for the resource This identifier is defined by the server|
|ip_version|text|The IP Version that will be used by this forwarding rule Valid options are IPV4 or IPV6 This can only be specified for an external global forwarding rule|
|is_mirroring_collector|boolean|Indicates whether or not this load balancer can be used as a collector for packet mirroring To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL|
|kind|text|Type of the resource Always compute#forwardingRule for Forwarding Rule resources|
|label_fingerprint|text|A fingerprint for the labels being applied to this resource|
|labels|jsonb|Labels for this resource|
|load_balancing_scheme|text|Specifies the forwarding rule type  - EXTERNAL is used for: - Classic Cloud VPN gateways - Protocol forwarding to VMs from an external IP address - HTTP(S), SSL Proxy, TCP Proxy, and Network Load Balancing - INTERNAL is used for: - Protocol forwarding to VMs from an internal IP address - Internal TCP/UDP Load Balancing - INTERNAL_MANAGED is used for: - Internal HTTP(S) Load Balancing - INTERNAL_SELF_MANAGED is used for: - Traffic Director  For more information about forwarding rules, refer to Forwarding rule concepts|
|name|text|Name of the resource|
|network|text|This field is not used for external load balancing  For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule If this field is not specified, the default network will be used  For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided|
|network_tier|text|This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD  For regional ForwardingRule, the valid values are PREMIUM and STANDARD For GlobalForwardingRule, the valid value is PREMIUM  If this field is not specified, it is assumed to be PREMIUM If IPAddress is specified, this value must be equal to the networkTier of the Address|
|port_range|text|This field can be used only if: * Load balancing scheme is one of EXTERNAL,  INTERNAL_SELF_MANAGED or INTERNAL_MANAGED, and * IPProtocol is one of TCP, UDP, or SCTP  Packets addressed to ports in the specified range will be forwarded to target or  backend_service You can only use one of ports, port_range, or allPorts The three are mutually exclusive Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges|
|ports|text[]|The ports field is only supported when the forwarding rule references a backend_service directly Supported load balancing products are Internal TCP/UDP Load Balancing and Network Load Balancing Only packets addressed to the specified list of ports are forwarded to backends  You can only use one of ports and port_range, or allPorts The three are mutually exclusive  You can specify a list of up to five ports, which can be non-contiguous  For Internal TCP/UDP Load Balancing, if you specify allPorts, you should not specify ports  For more information, see Port specifications (/load-balancing/docs/forwarding-rule-concepts#port_specifications)|
|psc_connection_id|bigint|The PSC connection id of the PSC Forwarding Rule|
|region|text|URL of the region where the regional forwarding rule resides|
|self_link|text|Server-defined URL for the resource|
|service_label|text|An optional prefix to the service name for this Forwarding Rule|
|service_name|text|The internal fully qualified service name for this Forwarding Rule  This field is only used for internal load balancing|
|subnetwork|text|This field is only used for internal load balancing|
|target|text||
