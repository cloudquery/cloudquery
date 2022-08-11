
# Table: azure_cosmosdb_account_cors
CorsPolicy the CORS policy for the Cosmos DB database account.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_cq_id|uuid|Unique CloudQuery ID of azure_cosmosdb_accounts table (FK)|
|allowed_origins|text|The origin domains that are permitted to make a request against the service via CORS.|
|allowed_methods|text|The methods (HTTP request verbs) that the origin domain may use for a CORS request.|
|allowed_headers|text|The request headers that the origin domain may specify on the CORS request.|
|exposed_headers|text|The response headers that may be sent in the response to the CORS request and exposed by the browser to the request issuer.|
|max_age_in_seconds|bigint|The maximum amount time that a browser should cache the preflight OPTIONS request.|
