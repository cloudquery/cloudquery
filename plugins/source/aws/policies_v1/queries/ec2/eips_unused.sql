insert into aws_policy_results
select :'execution_time' as execution_time,
       :'framework'      as framework,
       :'check_id'       as check_id,
       'Unused EC2 EIP'  as title,
       account_id,
       allocation_id     as resource_id,
       'fail'            as status
from aws_ec2_eips
where association_id is null
