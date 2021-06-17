
# Table: gcp_compute_interconnects
Represents an Interconnect resource  An Interconnect resource is a dedicated connection between the GCP network and your on-premises network
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|admin_enabled|boolean|Administrative status of the interconnect When this is set to true, the Interconnect is functional and can carry traffic When set to false, no packets can be carried over the interconnect and no BGP routes are exchanged over it By default, the status is set to true|
|creation_timestamp|text|Creation timestamp in RFC3339 text format|
|customer_name|text|Customer name, to put in the Letter of Authorization as the party authorized to request a crossconnect|
|description|text|An optional description of this resource Provide this property when you create the resource|
|google_ip_address|text|IP address configured on the Google side of the Interconnect link This can be used only for ping tests|
|google_reference_id|text|Google reference ID to be used when raising support tickets with Google or otherwise to debug backend connectivity issues|
|resource_id|text|The unique identifier for the resource This identifier is defined by the server|
|interconnect_attachments|text[]|A list of the URLs of all InterconnectAttachments configured to use this Interconnect|
|interconnect_type|text|Type of interconnect, which can take one of the following values: - PARTNER: A partner-managed interconnection shared between customers though a partner - DEDICATED: A dedicated physical interconnection with the customer Note that a value IT_PRIVATE has been deprecated in favor of DEDICATED|
|kind|text|Type of the resource Always compute#interconnect for interconnects|
|link_type|text|Type of link requested, which can take one of the following values: - LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics - LINK_TYPE_ETHERNET_100G_LR: A 100G Ethernet with LR optics Note that this field indicates the speed of each of the links in the bundle, not the speed of the entire bundle|
|location|text|URL of the InterconnectLocation object that represents where this connection is to be provisioned|
|name|text|Name of the resource Provided by the client when the resource is created|
|noc_contact_email|text|Email address to contact the customer NOC for operations and maintenance notifications regarding this Interconnect If specified, this will be used for notifications in addition to all other forms described, such as Stackdriver logs alerting and Cloud Notifications|
|operational_status|text|The current status of this Interconnect's functionality, which can take one of the following values: - OS_ACTIVE: A valid Interconnect, which is turned up and is ready to use Attachments may be provisioned on this Interconnect - OS_UNPROVISIONED: An Interconnect that has not completed turnup No attachments may be provisioned on this Interconnect - OS_UNDER_MAINTENANCE: An Interconnect that is undergoing internal maintenance No attachments may be provisioned or updated on this Interconnect|
|peer_ip_address|text|IP address configured on the customer side of the Interconnect link The customer should configure this IP address during turnup when prompted by Google NOC This can be used only for ping tests|
|provisioned_link_count|bigint|Number of links actually provisioned in this interconnect|
|requested_link_count|bigint|Target number of physical links in the link bundle, as requested by the customer|
|self_link|text|Server-defined URL for the resource|
|state|text|The current state of Interconnect functionality, which can take one of the following values: - ACTIVE: The Interconnect is valid, turned up and ready to use Attachments may be provisioned on this Interconnect - UNPROVISIONED: The Interconnect has not completed turnup No attachments may be provisioned on this Interconnect - UNDER_MAINTENANCE: The Interconnect is undergoing internal maintenance No attachments may be provisioned or updated on this Interconnect|
