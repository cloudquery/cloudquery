insert into aws_policy_results
with attachment as (select distinct ebs_volume_cq_id from aws_ec2_ebs_volume_attachments)
select :'execution_time'     as execution_time,
       :'framework'          as framework,
       :'check_id'           as check_id,
       'Detached EBS volume' as title,
       volume.account_id,
       volume.arn            as resource_id,
       'fail'                as status
from aws_ec2_ebs_volumes volume
         left join attachment on attachment.ebs_volume_cq_id = volume.cq_id
where attachment.ebs_volume_cq_id is null