insert into aws_policy_results
with instance as (select distinct host_cq_id from aws_ec2_host_instances)
select :'execution_time'       as execution_time,
       :'framework'            as framework,
       :'check_id'             as check_id,
       'Unused dedicated host' as title,
       host.account_id,
       host.arn                as resource_id,
       'fail'                  as status
from aws_ec2_hosts host
         left join instance on instance.host_cq_id = host.cq_id
where instance.host_cq_id is null