insert into aws_policy_results
with association as (select distinct route_table_cq_id from aws_ec2_route_table_associations)
select :'execution_time'    as execution_time,
       :'framework'         as framework,
       :'check_id'          as check_id,
       'Unused route table' as title,
       route_table.account_id,
       route_table.arn      as resource_id,
       'fail'               as status
from aws_ec2_route_tables route_table
         left join association on association.route_table_cq_id = route_table.cq_id
where association.route_table_cq_id is null