
# Table: k8s_core_endpoint_subset_ports
EndpointPort is a tuple that describes a single port.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|endpoint_subset_cq_id|uuid|Unique CloudQuery ID of k8s_core_endpoint_subsets table (FK)|
|name|text|The name of this port|
|port|integer|The port number of the endpoint.|
|protocol|text|The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.|
|app_protocol|text|The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.|
