insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Application load balancers should be configured to drop HTTP headers' as title,
  lb.account_id,
  lb.arn as resource_id,
  case when
       lb.type = 'application' and (a.value)::boolean is not true -- TODO check
    then 'fail'
    else 'pass'
  end as status
from aws_elbv2_load_balancers lb
inner join
    aws_elbv2_load_balancer_attributes a on
        a.load_balancer_arn = lb.arn and a.key='routing.http.drop_invalid_header_fields.enabled'
