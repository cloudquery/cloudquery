
# Table: aws_cognito_user_pool_schema_attributes
Contains information about the schema attribute.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_pool_cq_id|uuid|Unique CloudQuery ID of aws_cognito_user_pools table (FK)|
|user_pool_id|text|The ID of the user pool.|
|attribute_data_type|text|The attribute data type.|
|developer_only_attribute|boolean|We recommend that you use WriteAttributes (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolClientType.html#CognitoUserPools-Type-UserPoolClientType-WriteAttributes) in the user pool client to control how attributes can be mutated for new use cases instead of using DeveloperOnlyAttribute|
|mutable|boolean|Specifies whether the value of the attribute can be changed|
|name|text|A schema attribute of the name type.|
|number_attribute_constraints_max_value|text|The maximum value of an attribute that is of the number data type.|
|number_attribute_constraints_min_value|text|The minimum value of an attribute that is of the number data type.|
|required|boolean|Specifies whether a user pool attribute is required|
|string_attribute_constraints_max_length|text|The maximum length.|
|string_attribute_constraints_min_length|text|The minimum length.|
