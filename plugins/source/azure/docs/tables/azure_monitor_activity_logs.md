
# Table: azure_monitor_activity_logs
Azure network watcher
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|authorization_action|text|the permissible actions For instance: microsoftsupport/supporttickets/write|
|authorization_role|text|the role of the user For instance: Subscription Admin|
|authorization_scope|text|the scope|
|claims|jsonb|key value pairs to identify ARM permissions|
|caller|text|the email address of the user who has performed the operation, the UPN claim or SPN claim based on availability|
|description|text|the description of the event|
|id|text|the Id of this event as required by ARM for RBAC It contains the EventDataID and a timestamp information|
|event_data_id|text|the event data Id This is a unique identifier for an event|
|correlation_id|text|the correlation Id, usually a GUID in the string format The correlation Id is shared among the events that belong to the same uber operation|
|event_name_value|text|the invariant value|
|event_name_localized_value|text|the locale specific value|
|category_value|text|the invariant value|
|category_localized_value|text|the locale specific value|
|http_request_client_request_id|text|the client request id|
|http_request_client_ip_address|text|the client Ip Address|
|http_request_method|text|the Http request method|
|http_request_uri|text|the Uri|
|level|text|the event level Possible values include: 'EventLevelCritical', 'EventLevelError', 'EventLevelWarning', 'EventLevelInformational', 'EventLevelVerbose'|
|resource_group_name|text|the resource group name of the impacted resource|
|resource_provider_name_value|text|the invariant value|
|resource_provider_name_localized_value|text|the locale specific value|
|resource_id|text|the resource uri that uniquely identifies the resource that caused this event|
|resource_type_value|text|the invariant value|
|resource_type_localized_value|text|the locale specific value|
|operation_id|text|It is usually a GUID shared among the events corresponding to single operation This value should not be confused with EventName|
|operation_name_value|text|the invariant value|
|operation_name_localized_value|text|the locale specific value|
|properties|jsonb|the set of <Key, Value> pairs (usually a Dictionary<String, String>) that includes details about the event|
|status_value|text|the invariant value|
|status_localized_value|text|the locale specific value|
|sub_status_value|text|the invariant value|
|sub_status_localized_value|text|the locale specific value|
|event_timestamp_time|timestamp without time zone||
|submission_timestamp_time|timestamp without time zone||
|subscription_id|text|the Azure subscription Id usually a GUID|
|tenant_id|text|the Azure tenant Id|
