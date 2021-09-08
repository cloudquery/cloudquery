
# Table: digitalocean_cdn

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|text|A unique ID that can be used to identify and reference a CDN endpoint.|
|origin|text|The fully qualified domain name (FQDN) for the origin server which provides the content for the CDN. This is currently restricted to a Space.|
|endpoint|text|The fully qualified domain name (FQDN) from which the CDN-backed content is served.|
|created_at|timestamp without time zone|A time value given in ISO8601 combined date and time format that represents when the CDN endpoint was created.|
|ttl|bigint|The amount of time the content is cached by the CDN's edge servers in seconds. TTL must be one of 60, 600, 3600, 86400, or 604800. Defaults to 3600 (one hour) when excluded.|
|certificate_id|text|The ID of a DigitalOcean managed TLS certificate used for SSL when a custom subdomain is provided.|
|custom_domain|text|The fully qualified domain name (FQDN) of the custom subdomain used with the CDN endpoint.|
