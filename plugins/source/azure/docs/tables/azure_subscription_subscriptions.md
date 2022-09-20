
# Table: azure_subscription_subscriptions
Azure subscription information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|authorization_source|text|The authorization source of the request|
|managed_by_tenants|text[]|An array containing the tenants managing the subscription|
|location_placement_id|text|The subscription location placement ID|
|quota_id|text|The subscription quota ID|
|spending_limit|text|The subscription spending limit|
|tags|jsonb|The tags attached to the subscription|
|display_name|text|The subscription display name|
|id|text|The fully qualified ID for the subscription|
|state|text|The subscription state|
|tenant_id|text|The subscription tenant ID|
