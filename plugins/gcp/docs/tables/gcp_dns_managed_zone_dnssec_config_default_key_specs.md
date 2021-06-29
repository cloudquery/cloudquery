
# Table: gcp_dns_managed_zone_dnssec_config_default_key_specs
Parameters for DnsKey key generation Used for generating initial keys for a new ManagedZone and as default when adding a new DnsKey
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|managed_zone_id|uuid|Unique ID of gcp_dns_managed_zones table (FK)|
|algorithm|text|String mnemonic specifying the DNSSEC algorithm of this key|
|key_length|bigint|Length of the keys in bits|
|key_type|text|Specifies whether this is a key signing key (KSK) or a zone signing key (ZSK) Key signing keys have the Secure Entry Point flag set and, when active, are only used to sign resource record sets of type DNSKEY Zone signing keys do not have the Secure Entry Point flag set and are used to sign all other types of resource record sets|
|kind|text||
