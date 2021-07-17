
# Table: azure_monitor_diagnostic_settings
DiagnosticSettingsResource the diagnostic setting resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|storage_account_id|text|The resource ID of the storage account to which you would like to send Diagnostic Logs|
|service_bus_rule_id|text|The service bus rule Id of the diagnostic setting This is here to maintain backwards compatibility|
|event_hub_authorization_rule_id|text|The resource Id for the event hub authorization rule|
|event_hub_name|text|The name of the event hub If none is specified, the default event hub will be selected|
|workspace_id|text|The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/MicrosoftOperationalInsights/workspaces/viruela2|
|log_analytics_destination_type|text|A string indicating whether the export to Log Analytics should use the default destination type, ie AzureDiagnostics, or use a destination type constructed as follows: <normalized service identity>_<normalized category name> Possible values are: Dedicated and null (null is default)|
|id|text|Azure resource Id|
|name|text|Azure resource name|
|type|text|Azure resource type|
