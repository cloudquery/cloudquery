
# Table: gcp_compute_url_map_host_rules
UrlMaps A host-matching rule for a URL If matched, will use the named PathMatcher to select the BackendService
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|url_map_cq_id|uuid|Unique CloudQuery ID of gcp_compute_url_maps table (FK)|
|description|text|An optional description of this resource Provide this property when you create the resource|
|hosts|text[]|The list of host patterns to match They must be valid hostnames with optional port numbers in the format host:port * matches any string of ([a-z0-9-]*) In that case, * must be the first character and must be followed in the pattern by either - or  * based matching is not supported when the URL map is bound to target gRPC proxy that has validateForProxyless field set to true|
|path_matcher|text|The name of the PathMatcher to use to match the path portion of the URL if the hostRule matches the URL's host portion|
