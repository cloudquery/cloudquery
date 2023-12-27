insert into aws_policy_results
select :'execution_time'       as execution_time,
       :'framework'            as framework,
       :'check_id'             as check_id,
       'Unused EFS filesystem' as title,
       account_id,
       arn                     as resource_id,
       'fail'                  as status
from aws_efs_filesystems
where number_of_mount_targets = 0