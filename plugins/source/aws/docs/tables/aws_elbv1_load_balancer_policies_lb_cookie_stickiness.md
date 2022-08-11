
# Table: aws_elbv1_load_balancer_policies_lb_cookie_stickiness
Information about a policy for duration-based session stickiness.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_cq_id|uuid|Unique CloudQuery ID of aws_elbv1_load_balancers table (FK)|
|load_balance_name|text|The name of the load balancer.|
|cookie_expiration_period|bigint|The time period, in seconds, after which the cookie should be considered stale.|
|policy_name|text|The name of the policy.|
