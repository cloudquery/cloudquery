insert into aws_policy_results
select
  :execution_time as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Amazon ECS services should not have public IP addresses assigned to them automatically' as title,
  account_id,
  s.arn as resource_id,
  case when
    network_configuration_awsvpc_configuration_assign_public_ip is distinct from 'DISABLED'
    then 'fail'
    else 'pass'
  end as status
from aws_ecs_clusters c
     left join aws_ecs_cluster_services s ON c.cq_id = s.cluster_cq_id
