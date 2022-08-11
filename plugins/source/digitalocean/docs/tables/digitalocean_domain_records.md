
# Table: digitalocean_domain_records
DomainRecord represents a DigitalOcean DomainRecord
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|domain_cq_id|uuid|Unique CloudQuery ID of digitalocean_domains table (FK)|
|id|bigint|A unique identifier for each domain record.|
|type|text|The type of the DNS record. For example: A, CNAME, TXT, ...|
|name|text|The host name, alias, or service being defined by the record.|
|data|text|Variable data depending on record type. For example, the "data" value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.|
|priority|bigint|The priority for SRV and MX records.|
|port|bigint|The port for SRV records.|
|ttl|bigint|This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.|
|weight|bigint|The weight for SRV records.|
|flags|bigint|An unsigned integer between 0-255 used for CAA records.|
|tag|text|The parameter tag for CAA records. Valid values are "issue", "issuewild", or "iodef"|
