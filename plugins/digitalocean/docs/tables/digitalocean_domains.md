
# Table: digitalocean_domains
Domain represents a DigitalOcean domain
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|name|text|The name of the domain itself. This should follow the standard domain format of domain.TLD. For instance, `example.com` is a valid domain name.|
|ttl|bigint|This value is the time to live for the records on this domain, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.|
|zone_file|text|This attribute contains the complete contents of the zone file for the selected domain. Individual domain record resources should be used to get more granular control over records. However, this attribute can also be used to get information about the SOA record, which is created automatically and is not accessible as an individual record resource.|
