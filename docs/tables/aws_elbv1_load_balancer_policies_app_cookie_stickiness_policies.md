
# Table: aws_elbv1_load_balancer_policies_app_cookie_stickiness_policies
Information about a policy for application-controlled session stickiness.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid|Unique ID of aws_elbv1_load_balancers table (FK)|
|cookie_name|text|The name of the application cookie used for stickiness.|
|policy_name|text|The mnemonic name for the policy being created.|
