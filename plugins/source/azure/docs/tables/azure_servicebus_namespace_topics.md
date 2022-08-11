
# Table: azure_servicebus_namespace_topics
Description of servicebus namespace topic resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|namespace_cq_id|uuid|Unique CloudQuery ID of azure_servicebus_namespaces table (FK)|
|size_in_bytes|bigint|Size of the topic, in bytes|
|created_at_time|timestamp without time zone||
|updated_at_time|timestamp without time zone||
|accessed_at_time|timestamp without time zone||
|subscription_count|integer|Number of subscriptions|
|count_details_active_message_count|bigint|Number of active messages in the queue, topic, or subscription|
|count_details_dead_letter_message_count|bigint|Number of messages that are dead lettered|
|count_details_scheduled_message_count|bigint|Number of scheduled messages|
|count_details_transfer_message_count|bigint|Number of messages transferred to another queue, topic, or subscription|
|count_details_transfer_dead_letter_message_count|bigint|Number of messages transferred into dead letters|
|default_message_time_to_live|text|ISO 8601 Default message timespan to live value|
|max_size_in_megabytes|integer|Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic|
|max_message_size_in_kilobytes|bigint|Maximum size (in KB) of the message payload that can be accepted by the topic|
|requires_duplicate_detection|boolean|Value indicating if this topic requires duplicate detection|
|duplicate_detection_history_time_window|text|ISO8601 timespan structure that defines the duration of the duplicate detection history|
|enable_batched_operations|boolean|Value that indicates whether server-side batched operations are enabled|
|status|text|Enumerates the possible values for the status of a messaging entity|
|support_ordering|boolean|Value that indicates whether the topic supports ordering|
|auto_delete_on_idle|text|ISO 8601 timespan idle interval after which the topic is automatically deleted|
|enable_partitioning|boolean|Value that indicates whether the topic to be partitioned across multiple message brokers is enabled|
|enable_express|boolean|Value that indicates whether Express Entities are enabled|
|system_data|jsonb|The system meta data relating to this resource|
|id|text|Resource Id|
|name|text|Resource name|
|type|text|Resource type|
