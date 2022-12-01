# Table: azure_streamanalytics_streaming_jobs

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2020-03-01/streamanalytics#StreamingJob

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|sku|JSON|
|job_id|String|
|provisioning_state|String|
|job_state|String|
|job_type|String|
|output_start_mode|String|
|output_start_time|Timestamp|
|last_output_event_time|Timestamp|
|events_out_of_order_policy|String|
|output_error_policy|String|
|events_out_of_order_max_delay_in_seconds|Int|
|events_late_arrival_max_delay_in_seconds|Int|
|data_locale|String|
|compatibility_level|String|
|created_date|Timestamp|
|inputs|JSON|
|transformation|JSON|
|outputs|JSON|
|functions|JSON|
|etag|String|
|job_storage_account|JSON|
|content_storage_policy|String|
|cluster|JSON|
|identity|JSON|
|tags|JSON|
|location|String|
|id (PK)|String|
|name|String|
|type|String|