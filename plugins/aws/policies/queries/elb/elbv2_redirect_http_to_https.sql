insert into aws_policy_results
select
  :execution_time as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Application Load Balancer should be configured to redirect all HTTP requests to HTTPS' as title,
  aws_elbv2_listeners.account_id,
  aws_elbv2_listeners.arn as resource_id,
  case when
   protocol = 'HTTP' and (
        aws_elbv2_listener_default_actions.type != 'REDIRECT' or aws_elbv2_listener_default_actions.redirect_config_protocol != 'HTTPS')
    then 'fail'
    else 'pass'
  end as status
from aws_elbv2_listeners
inner join
    aws_elbv2_listener_default_actions on
        aws_elbv2_listeners.cq_id = aws_elbv2_listener_default_actions.listener_cq_id
