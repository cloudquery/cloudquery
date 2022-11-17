# Table: azure_streamanalytics_streaming_jobs

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics#StreamingJob

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|identity|JSON|
|location|String|
|cluster|JSON|
|compatibility_level|String|
|content_storage_policy|String|
|data_locale|String|
|events_late_arrival_max_delay_in_seconds|Int|
|events_out_of_order_max_delay_in_seconds|Int|
|events_out_of_order_policy|String|
|functions|JSON|
|inputs|JSON|
|job_storage_account|JSON|
|job_type|String|
|output_error_policy|String|
|output_start_mode|String|
|output_start_time|Timestamp|
|outputs|JSON|
|sku|JSON|
|transformation|JSON|
|created_date|Timestamp|
|etag|String|
|job_id|String|
|job_state|String|
|last_output_event_time|Timestamp|
|provisioning_state|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|