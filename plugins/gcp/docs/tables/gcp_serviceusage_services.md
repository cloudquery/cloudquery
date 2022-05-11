
# Table: gcp_serviceusage_services
A service that is available for use by the consumer
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|name|text||
|authentication|jsonb|Auth configuration|
|documentation|jsonb|Additional API documentation|
|title|text|The product title for this service|
|usage_producer_notification_channel|text|The full resource name of a channel used for sending notifications to the service producer|
|usage_requirements|text[]|Requirements that must be satisfied before a consumer project can use the service|
|parent|text|The resource name of the consumer|
|state|text|"STATE_UNSPECIFIED" - The default value, which indicates that the enabled state of the service is unspecified or not meaningful Currently, all consumers other than projects (such as folders and organizations) are always in this state   "DISABLED" - The service cannot be used by this consumer|
