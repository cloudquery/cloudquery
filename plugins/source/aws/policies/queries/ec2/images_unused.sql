insert into aws_policy_results
with mapping as (select distinct image_cq_id from aws_ec2_image_block_device_mappings)
select :'execution_time'      as execution_time,
       :'framework'           as framework,
       :'check_id'            as check_id,
       'Unused own EC2 image' as title,
       image.account_id,
       image.arn              as resource_id,
       'fail'                 as status
from aws_ec2_images image
         left join mapping on mapping.image_cq_id = image.cq_id
where image.account_id = image.owner_id
  and mapping.image_cq_id is null