
# Table: aws_directconnect_gateway_attachments
Information about the attachment between a Direct Connect gateway and virtual interfaces.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|directconnect_gateway_id|uuid|Unique ID of aws_directconnect_gateways table (FK)|
|attachment_state|text|The state of the attachment.|
|attachment_type|text|The type of attachment.|
|state_change_error|text|The error message if the state of an object failed to advance.|
|virtual_interface_id|text|The ID of the virtual interface.|
|virtual_interface_owner_account|text|The ID of the AWS account that owns the virtual interface.|
|virtual_interface_region|text|The AWS Region where the virtual interface is located.|
