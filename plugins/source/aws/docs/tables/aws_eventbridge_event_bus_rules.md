
# Table: aws_eventbridge_event_bus_rules
Contains information about a rule in Amazon EventBridge
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|event_bus_cq_id|uuid|Unique CloudQuery ID of aws_eventbridge_event_buses table (FK)|
|tags|jsonb||
|arn|text|The Amazon Resource Name (ARN) of the rule|
|description|text|The description of the rule|
|event_bus_name|text|The name or ARN of the event bus associated with the rule|
|event_pattern|text|The event pattern of the rule|
|managed_by|text|If the rule was created on behalf of your account by an Amazon Web Services service, this field displays the principal name of the service that created the rule|
|name|text|The name of the rule|
|role_arn|text|The Amazon Resource Name (ARN) of the role that is used for target invocation If you're setting an event bus in another account as the target and that account granted permission to your account through an organization instead of directly by the account ID, you must specify a RoleArn with proper permissions in the Target structure, instead of here in this parameter|
|schedule_expression|text|The scheduling expression|
|state|text|The state of the rule|
