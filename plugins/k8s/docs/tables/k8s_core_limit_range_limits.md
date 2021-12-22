
# Table: k8s_core_limit_range_limits
LimitRangeItem defines a min/max usage limit for any resource that matches on kind.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|limit_range_cq_id|uuid|Unique CloudQuery ID of k8s_core_limit_ranges table (FK)|
|type|text|Type of resource that this limit applies to.|
|max|jsonb|Max usage constraints on this kind by resource name.|
|min|jsonb|Min usage constraints on this kind by resource name.|
|default|jsonb|Default resource requirement limit value by resource name if resource limit is omitted.|
|default_request|jsonb|DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.|
|max_limit_request_ratio|jsonb|MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.|
