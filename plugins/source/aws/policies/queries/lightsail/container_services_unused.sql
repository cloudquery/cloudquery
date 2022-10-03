insert into aws_policy_results
with deployment as (select distinct container_service_arn from aws_lightsail_container_service_deployments)
select :'execution_time'                     as execution_time,
       :'framework'                          as framework,
       :'check_id'                           as check_id,
       'Unused Lightsail container services' as title,
       cs.account_id,
       cs.arn                                as resource_id,
       'fail'                                as status
from aws_lightsail_container_services cs
         left join deployment on deployment.container_service_arn = cs.arn
where deployment.container_service_arn is null;
