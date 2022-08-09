
# Table: cloudflare_certificate_packs
CertificatePack is the overarching structure of a certificate pack response.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The Account ID of the resource.|
|zone_id|text|The Zone ID of the resource.|
|id|text|The unique identifier for a certificate_pack|
|type|text|Type of certificate pack|
|hosts|text[]|comma separated list of valid host names for the certificate packs. Must contain the zone apex, may not contain more than 50 hosts, and may not be empty.|
|primary_certificate|text|Identifier of the primary certificate in a pack|
