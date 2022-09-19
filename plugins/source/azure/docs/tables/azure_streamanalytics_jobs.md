
# Table: azure_streamanalytics_jobs
StreamingJob a streaming job object, containing all information associated with the named streaming job.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id.|
|sku_name|text|The name of the SKU.|
|job_id|text|A GUID uniquely identifying the streaming job.|
|provisioning_state|text|Describes the provisioning status of the streaming job.|
|job_state|text|Describes the state of the streaming job.|
|job_type|text|Describes the type of the job.|
|output_start_mode|text|This property should only be utilized when it is desired that the job be started immediately upon creation.|
|output_start_time|timestamp without time zone|Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started.|
|last_output_event_time|timestamp without time zone|Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced.|
|events_out_of_order_policy|text|Indicates the policy to apply to events that arrive out of order in the input event stream.|
|output_error_policy|text|Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed.|
|events_out_of_order_max_delay|integer|The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.|
|events_late_arrival_max_delay|integer|The maximum tolerable delay in seconds where events arriving late could be included.|
|data_locale|text|The data locale of the stream analytics job.|
|compatibility_level|text|Controls certain runtime behaviors of the streaming job.|
|created_date|timestamp without time zone|Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.|
|transformation_properties_streaming_units|integer|Specifies the number of streaming units that the streaming job uses.|
|transformation_properties_valid_streaming_units|integer[]|Specifies the valid streaming units a streaming job can scale to.|
|transformation_properties_query|text|Specifies the query that will be run in the streaming job.|
|transformation_properties_etag|text|The current entity tag for the transformation|
|transformation_id|text|Transformation resource Id.|
|transformation_name|text|Transformation resource name.|
|transformation_type|text|Transformation resource type.|
|etag|text|The current entity tag for the streaming job.|
|job_storage_account_authentication_mode|text|Authentication Mode.|
|job_storage_account_name|text|The name of the Azure Storage account.|
|job_storage_account_key|text|The account key for the Azure Storage account.|
|content_storage_policy|text|Valid values are JobStorageAccount and SystemAccount|
|cluster_id|text|The resource id of cluster.|
|identity_tenant_id|text|The identity tenantId.|
|identity_principal_id|text|The identity principal ID.|
|identity_type|text|The identity type|
|tags|jsonb|Resource tags.|
|location|text|The geo-location where the resource lives.|
|id|text|Fully qualified resource Id for the resource.|
|name|text|The name of the resource.|
|type|text|The type of the resource|
