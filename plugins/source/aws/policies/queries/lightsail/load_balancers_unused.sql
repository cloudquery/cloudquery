insert into aws_policy_results
with instance_health as (select distinct load_balancer_cq_id from aws_lightsail_load_balancer_instance_health_summary)
select :'execution_time'                 as execution_time,
       :'framework'                      as framework,
       :'check_id'                       as check_id,
       'Unused Lightsail load balancers' as title,
       lb.account_id,
       lb.arn                            as resource_id,
       'fail'                            as status
from aws_lightsail_load_balancers lb
         left join instance_health on instance_health.load_balancer_cq_id = lb.cq_id
where instance_health.load_balancer_cq_id is null