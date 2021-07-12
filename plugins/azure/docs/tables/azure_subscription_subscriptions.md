
# Table: azure_subscription_subscriptions
Model subscription information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|text|The fully qualified ID for the subscription For example, /subscriptions/00000000-0000-0000-0000-000000000000|
|subscription_id|text|The subscription ID|
|display_name|text|The subscription display name|
|state|text|The subscription state Possible values are Enabled, Warned, PastDue, Disabled, and Deleted Possible values include: 'Enabled', 'Warned', 'PastDue', 'Disabled', 'Deleted'|
|location_placement_id|text|The subscription location placement ID The ID indicates which regions are visible for a subscription For example, a subscription with a location placement Id of Public_2014-09-01 has access to Azure public regions|
|quota_id|text|The subscription quota ID|
|spending_limit|text|The subscription spending limit Possible values include: 'On', 'Off', 'CurrentPeriodOff'|
|authorization_source|text|The authorization source of the request Valid values are one or more combinations of Legacy, RoleBased, Bypassed, Direct and Management For example, 'Legacy, RoleBased'|
