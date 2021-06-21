
# Table: aws_directconnect_gateway_associations
Information about the association between an Direct Connect Gateway and either a Virtual Private Gateway, or Transit Gateway
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|directconnect_gateway_id|uuid|Unique ID of aws_directconnect_gateways table (FK)|
|allowed_prefixes_to_direct_connect_gateway|text[]|The Amazon VPC prefixes to advertise to the Direct Connect gateway.|
|associated_gateway_id|text|The ID of the associated gateway.|
|associated_gateway_owner_account|text|The ID of the AWS account that owns the associated virtual private gateway or transit gateway.|
|associated_gateway_region|text|The Region where the associated gateway is located.|
|associated_gateway_type|text|The type of associated gateway.|
|association_id|text|The ID of the Direct Connect gateway association|
|association_state|text|The state of the association.|
|direct_connect_gateway_owner_account|text|The ID of the AWS account that owns the associated gateway.|
|state_change_error|text|The error message if the state of an object failed to advance.|
|virtual_gateway_id|text|The ID of the virtual private gateway. Applies only to private virtual interfaces.|
|virtual_gateway_owner_account|text|The ID of the AWS account that owns the virtual private gateway.|
|resource_id|text|The ID of the Direct Connect gateway association|
