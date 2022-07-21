insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Classic Load Balancers with HTTPS/SSL listeners should use a predefined security policy that has strong configuration' as title,
  aws_elbv1_load_balancers.account_id,
  aws_elbv1_load_balancers.arn as resource_id,
  case when
    aws_elbv1_load_balancer_listeners.listener_protocol in ('HTTPS', 'SSL')
    and 'ELBSecurityPolicy-TLS-1-2-2017-01' != any(
        aws_elbv1_load_balancers.other_policies)
    then 'fail'
    else 'pass'
  end as status
from aws_elbv1_load_balancers
inner join
    aws_elbv1_load_balancer_listeners on
        aws_elbv1_load_balancer_listeners.load_balancer_cq_id = aws_elbv1_load_balancers.cq_id
