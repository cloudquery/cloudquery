insert into aws_policy_results
with association as (select distinct network_acl_cq_id from aws_ec2_network_acl_associations)
select :'execution_time'                    as execution_time,
       :'framework'                         as framework,
       :'check_id'                          as check_id,
       'Unused network access control list' as title,
       acl.account_id,
       acl.arn                              as resource_id,
       'fail'                               as status
from aws_ec2_network_acls acl
         left join association on association.network_acl_cq_id = acl.cq_id
where association.network_acl_cq_id is null