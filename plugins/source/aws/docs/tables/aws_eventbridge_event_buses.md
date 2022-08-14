
# Table: aws_eventbridge_event_buses
An event bus receives events from a source and routes them to rules associated with that event bus
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb||
|arn|text|The ARN of the event bus|
|name|text|The name of the event bus|
|policy|text|The permissions policy of the event bus, describing which other Amazon Web Services accounts can write events to this event bus|
