
# Table: aws_elbv1_load_balancer_policies_app_cookie_stickiness
Information about a policy for application-controlled session stickiness.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_cq_id|uuid|Unique CloudQuery ID of aws_elbv1_load_balancers table (FK)|
|load_balance_name|text|The name of the load balancer.|
|cookie_name|text|The name of the application cookie used for stickiness.|
|policy_name|text|The mnemonic name for the policy being created.|
