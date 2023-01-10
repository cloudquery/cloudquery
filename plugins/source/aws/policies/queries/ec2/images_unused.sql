insert into aws_policy_results
select :'execution_time'      as execution_time,
       :'framework'           as framework,
       :'check_id'            as check_id,
       'Unused own EC2 image' as title,
       account_id,
       arn                    as resource_id,
       'fail'                 as status
from aws_ec2_images
where coalesce(jsonb_array_length(block_device_mappings), 0) = 0;
