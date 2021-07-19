
# Table: azure_network_public_ip_address_ip_tags
IPTag contains the IpTag associated with the object
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|public_ip_address_cq_id|uuid|Unique CloudQuery ID of azure_network_public_ip_addresses table (FK)|
|ip_tag_type|text|The IP tag type Example: FirstPartyUsage|
|tag|text|The value of the IP tag associated with the public IP Example: SQL|
