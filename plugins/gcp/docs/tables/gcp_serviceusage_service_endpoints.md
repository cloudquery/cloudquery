
# Table: gcp_serviceusage_service_endpoints
`Endpoint` describes a network address of a service that serves a set of APIs
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of gcp_serviceusage_services table (FK)|
|allow_cors|boolean|Allowing CORS (https://enwikipediaorg/wiki/Cross-origin_resource_sharing), aka cross-domain traffic, would allow the backends served from this endpoint to receive and respond to HTTP OPTIONS requests|
|name|text|The canonical name of this endpoint|
|target|text|The specification of an Internet routable address of API frontend that will handle requests to this API Endpoint (https://cloudgooglecom/apis/design/glossary)|
