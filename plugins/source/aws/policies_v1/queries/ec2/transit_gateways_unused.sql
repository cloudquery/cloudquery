insert into aws_policy_results
with attachment as (select distinct transit_gateway_arn from aws_ec2_transit_gateway_attachments)
select :'execution_time'        as execution_time,
       :'framework'             as framework,
       :'check_id'              as check_id,
       'Unused transit gateway' as title,
       gateway.account_id,
       gateway.arn              as resource_id,
       'fail'                   as status
from aws_ec2_transit_gateways gateway
         left join attachment on attachment.transit_gateway_arn = gateway.arn
where attachment.transit_gateway_arn is null;