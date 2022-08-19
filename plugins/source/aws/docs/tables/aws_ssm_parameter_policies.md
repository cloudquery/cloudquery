
# Table: aws_ssm_parameter_policies
One or more policies assigned to a parameter
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|parameter_cq_id|uuid|Unique CloudQuery ID of aws_ssm_parameters table (FK)|
|policy_status|text|The status of the policy|
|policy_text|text|The JSON text of the policy|
|policy_type|text|The type of policy|
