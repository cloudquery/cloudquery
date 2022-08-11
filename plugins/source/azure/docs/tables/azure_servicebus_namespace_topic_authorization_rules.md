
# Table: azure_servicebus_namespace_topic_authorization_rules
Description of servicebus namespace topic authorization rules
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|namespace_topic_cq_id|uuid|Unique CloudQuery ID of azure_servicebus_namespace_topics table (FK)|
|access_keys|jsonb||
|rights|text[]|The rights associated with the rule|
|system_data|jsonb|The system meta data relating to this resource|
|id|text|Resource Id|
|name|text|Resource name|
|type|text|Resource type|
