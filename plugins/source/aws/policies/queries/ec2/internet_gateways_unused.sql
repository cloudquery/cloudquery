insert into aws_policy_results
with attachment as (select distinct internet_gateway_cq_id from aws_ec2_internet_gateway_attachments)
select :'execution_time'         as execution_time,
       :'framework'              as framework,
       :'check_id'               as check_id,
       'Unused internet gateway' as title,
       gateway.account_id,
       gateway.arn               as resource_id,
       'fail'                    as status
from aws_ec2_internet_gateways gateway
         left join attachment on attachment.internet_gateway_cq_id = gateway.cq_id
where attachment.internet_gateway_cq_id is null