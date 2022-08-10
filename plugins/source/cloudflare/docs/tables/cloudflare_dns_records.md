
# Table: cloudflare_dns_records
DNSRecord represents a DNS record in a zone.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The Account ID of the resource.|
|created_on|timestamp without time zone|When the record was created.|
|modified_on|timestamp without time zone|When the record was last modified.|
|type|text|Record type.|
|name|text|DNS record name.|
|content|text|A valid IPv4 address.|
|meta|jsonb|Extra Cloudflare-specific information about the record.|
|data|jsonb|Metadata about the record.|
|id|text|DNS record identifier tag.|
|zone_id|text|Zone identifier tag.|
|zone_name|text|The domain of the record.|
|priority|integer|The priority of the record.|
|ttl|bigint|Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1 for 'automatic'|
|proxied|boolean|Whether the record is receiving the performance and security benefits of Cloudflare.|
|proxiable|boolean|Whether the record can be proxied by Cloudflare or not.|
|locked|boolean|Whether this record can be modified/deleted (true means it's managed by Cloudflare).|
