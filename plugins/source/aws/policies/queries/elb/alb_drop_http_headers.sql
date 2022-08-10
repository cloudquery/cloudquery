insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Application load balancers should be configured to drop HTTP headers' as title,
  account_id,
  arn as resource_id,
  case when
    aws_elbv2_load_balancers.type = 'application' and aws_elbv2_load_balancer_attributes.routing_http_drop_invalid_header_fields is not true
    then 'fail'
    else 'pass'
  end as status
from aws_elbv2_load_balancers
inner join
    aws_elbv2_load_balancer_attributes on
        aws_elbv2_load_balancer_attributes.load_balancer_cq_id = aws_elbv2_load_balancers.cq_id
