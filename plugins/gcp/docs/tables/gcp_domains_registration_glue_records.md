
# Table: gcp_domains_registration_glue_records
Defines a host on your domain that is a DNS name server for your domain and/or other domains Glue records are a way of making the IP address of a name server known, even when it serves DNS queries for its parent domain For example, when `nsexamplecom` is a name server for `examplecom`, the host `nsexamplecom` must have a glue record to break the circular DNS reference
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|registration_id|uuid|Unique ID of gcp_domains_registrations table (FK)|
|host_name|text|Required Domain name of the host in Punycode format|
|ipv4_addresses|text[]|List of IPv4 addresses corresponding to this host in the standard decimal format (eg `198511001`) At least one of `ipv4_address` and `ipv6_address` must be set|
|ipv6_addresses|text[]|List of IPv6 addresses corresponding to this host in the standard hexadecimal format (eg `2001:db8::`) At least one of `ipv4_address` and `ipv6_address` must be set|
