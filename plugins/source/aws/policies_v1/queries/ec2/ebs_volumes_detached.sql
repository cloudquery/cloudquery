insert into aws_policy_results
select :'execution_time'     as execution_time,
       :'framework'          as framework,
       :'check_id'           as check_id,
       'Detached EBS volume' as title,
       account_id,
       arn            as resource_id,
       'fail'                as status
from aws_ec2_ebs_volumes
where coalesce(jsonb_array_length(attachments), 0) = 0;
