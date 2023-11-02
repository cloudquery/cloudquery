# Table: fastly_stats_services

This table shows data for Fastly Stats Services.

https://developer.fastly.com/reference/api/metrics-stats/historical-stats/

The composite primary key for this table is (**start_time**, **service_id**, **region**, **by**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|start_time (PK)|`timestamp[us, tz=UTC]`|
|service_id (PK)|`utf8`|
|region (PK)|`utf8`|
|by (PK)|`utf8`|
|attack_req_header_bytes|`int64`|
|attack_req_body_bytes|`int64`|
|attack_resp_synth_bytes|`int64`|
|bereq_body_bytes|`int64`|
|bereq_header_bytes|`int64`|
|bandwidth|`int64`|
|billed_body_bytes|`int64`|
|billed_header_bytes|`int64`|
|errors|`int64`|
|http2|`int64`|
|hit_ratio|`float64`|
|hits|`int64`|
|hits_time|`float64`|
|ipv6|`int64`|
|imgopto|`int64`|
|log|`int64`|
|miss|`int64`|
|miss_time|`float64`|
|otfp|`int64`|
|object_size_100k|`int64`|
|object_size_100m|`int64`|
|object_size_10k|`int64`|
|object_size_10m|`int64`|
|object_size_1g|`int64`|
|object_size_1k|`int64`|
|object_size_1m|`int64`|
|pci|`int64`|
|pass|`int64`|
|pass_time|`float64`|
|pipe|`int64`|
|req_body_bytes|`int64`|
|req_header_bytes|`int64`|
|requests|`int64`|
|resp_body_bytes|`int64`|
|resp_header_bytes|`int64`|
|restarts|`int64`|
|shield|`int64`|
|shield_resp_body_bytes|`int64`|
|shield_resp_header_bytes|`int64`|
|status_1xx|`int64`|
|status_200|`int64`|
|status_204|`int64`|
|status_206|`int64`|
|status_2xx|`int64`|
|status_301|`int64`|
|status_302|`int64`|
|status_304|`int64`|
|status_3xx|`int64`|
|status_400|`int64`|
|status_401|`int64`|
|status_403|`int64`|
|status_404|`int64`|
|status_416|`int64`|
|status_4xx|`int64`|
|status_500|`int64`|
|status_501|`int64`|
|status_502|`int64`|
|status_503|`int64`|
|status_504|`int64`|
|status_505|`int64`|
|status_5xx|`int64`|
|synth|`int64`|
|tls|`int64`|
|tls_v10|`int64`|
|tls_v11|`int64`|
|tls_v12|`int64`|
|tls_v13|`int64`|
|uncachable|`int64`|
|video|`int64`|
|waf_blocked|`int64`|
|waf_logged|`int64`|
|waf_passed|`int64`|
|compute_bereq_body_bytes|`int64`|
|compute_bereq_errors|`int64`|
|compute_bereq_header_bytes|`int64`|
|compute_bereqs|`int64`|
|compute_beresp_body_bytes|`int64`|
|compute_beresp_header_bytes|`int64`|
|compute_execution_time_ms|`int64`|
|compute_globals_limit_exceeded|`int64`|
|compute_guest_errors|`int64`|
|compute_heap_limit_exceeded|`int64`|
|compute_ram_used|`int64`|
|compute_req_body_bytes|`int64`|
|compute_req_header_bytes|`int64`|
|compute_request_time_ms|`int64`|
|compute_requests|`int64`|
|compute_resource_limit_exceeded|`int64`|
|compute_resp_body_bytes|`int64`|
|compute_resp_header_bytes|`int64`|
|compute_resp_status_1xx|`int64`|
|compute_resp_status_2xx|`int64`|
|compute_resp_status_3xx|`int64`|
|compute_resp_status_4xx|`int64`|
|compute_resp_status_5xx|`int64`|
|compute_runtime_errors|`int64`|
|compute_stack_limit_exceeded|`int64`|
|edge_hit_requests|`int64`|
|edge_hit_resp_body_bytes|`int64`|
|edge_hit_resp_header_bytes|`int64`|
|edge_miss_requests|`int64`|
|edge_miss_resp_body_bytes|`int64`|
|edge_miss_resp_header_bytes|`int64`|
|edge_requests|`int64`|
|edge_resp_body_bytes|`int64`|
|edge_resp_header_bytes|`int64`|
|origin_cache_fetch_resp_body_bytes|`int64`|
|origin_cache_fetch_resp_header_bytes|`int64`|
|origin_cache_fetches|`int64`|
|origin_fetch_body_bytes|`int64`|
|origin_fetch_header_bytes|`int64`|
|origin_fetch_resp_body_bytes|`int64`|
|origin_fetch_resp_header_bytes|`int64`|
|origin_fetches|`int64`|
|fanout_bereq_body_bytes|`int64`|
|fanout_bereq_header_bytes|`int64`|
|fanout_beresp_body_bytes|`int64`|
|fanout_beresp_header_bytes|`int64`|
|fanout_conn_time_ms|`int64`|
|fanout_recv_publishes|`int64`|
|fanout_req_body_bytes|`int64`|
|fanout_req_header_bytes|`int64`|
|fanout_resp_body_bytes|`int64`|
|fanout_resp_header_bytes|`int64`|
|fanout_send_publishes|`int64`|