
# Table: aws_elbv1_load_balancer_policies
Information about a policy.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid|Unique ID of aws_elbv1_load_balancers table (FK)|
|policy_attribute_descriptions|jsonb|The policy attributes.|
|policy_name|text|The name of the policy.|
|policy_type_name|text|The name of the policy type.|
