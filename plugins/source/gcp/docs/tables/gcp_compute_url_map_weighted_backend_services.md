
# Table: gcp_compute_url_map_weighted_backend_services
In contrast to a single BackendService in HttpRouteAction to which all matching traffic is directed to, WeightedBackendService allows traffic to be split across multiple BackendServices
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|url_map_cq_id|uuid|Unique CloudQuery ID of gcp_compute_url_maps table (FK)|
|backend_service|text|The full or partial URL to the default BackendService resource Before forwarding the request to backendService, the loadbalancer applies any relevant headerActions specified as part of this backendServiceWeight|
|header_action|jsonb|Specifies changes to request and response headers that need to take effect for the selected backendService headerAction specified here take effect before headerAction in the enclosing HttpRouteRule, PathMatcher and UrlMap Note that headerAction is not supported for Loadbalancers that have their loadBalancingScheme set to EXTERNAL Not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true|
|weight|bigint|Specifies the fraction of traffic sent to backendService, computed as weight / (sum of all weightedBackendService weights in routeAction)  The selection of a backend service is determined only for new traffic Once a user's request has been directed to a backendService, subsequent requests will be sent to the same backendService as determined by the BackendService's session affinity policy|
