
# Table: gcp_monitoring_alert_policies
A description of the conditions under which some aspect of your system is considered to be "unhealthy" and the ways to notify people or services about this state For an overview of alert policies, see Introduction to Alerting (https://cloudgooglecom/monitoring/alerts/)
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|combiner|text|How to combine the results of multiple conditions to determine if an incident should be opened If condition_time_series_query_language is present, this must be COMBINE_UNSPECIFIED  Possible values:   "COMBINE_UNSPECIFIED" - An unspecified combiner   "AND" - Combine conditions using the logical AND operator An incident is created only if all the conditions are met simultaneously This combiner is satisfied if all conditions are met, even if they are met on completely different resources   "OR" - Combine conditions using the logical OR operator An incident is created if any of the listed conditions is met   "AND_WITH_MATCHING_RESOURCE" - Combine conditions using logical AND operator, but unlike the regular AND option, an incident is created only if all conditions are met simultaneously on at least one resource|
|creation_record_mutate_time|text|When the change occurred|
|creation_record_mutated_by|text|The email address of the user making the change|
|display_name|text|A short name or phrase used to identify the policy in dashboards, notifications, and incidents To avoid confusion, don't use the same display name for multiple policies in the same project The name is limited to 512 Unicode characters|
|documentation_content|text|The text of the documentation, interpreted according to mime_type The content may not exceed 8,192 Unicode characters and may not exceed more than 10,240 bytes when encoded in UTF-8 format, whichever is smaller|
|documentation_mime_type|text|The format of the content field Presently, only the value "text/markdown" is supported See Markdown (https://enwikipediaorg/wiki/Markdown) for more information|
|enabled|boolean|Whether or not the policy is enabled On write, the default interpretation if unset is that the policy is enabled On read, clients should not make any assumption about the state if it has not been populated The field should always be populated on List and Get operations, unless a field projection has been specified that strips it out|
|mutate_time|text|When the change occurred|
|mutated_by|text|The email address of the user making the change|
|name|text|The resource name for this policy|
|notification_channels|text[]|Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when new violations occur on an already opened incident Each element of this array corresponds to the name field in each of the NotificationChannel objects that are returned from the ListNotificationChannels method|
|labels|jsonb|Labels for this resource|
|validity_code|bigint|The status code, which should be an enum value of googlerpcCode|
|validity_message|text|A developer-facing error message, which should be in English Any user-facing error message should be localized and sent in the googlerpcStatusdetails field, or localized by the client|
