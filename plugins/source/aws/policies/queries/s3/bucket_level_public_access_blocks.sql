insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'S3 Block Public Access setting should be enabled at the bucket-level' as title,
    s3.account_id,
    s3.arn AS resource_id,
    case 
        when (s3.block_public_acls or as3a.block_public_acls)
         and (s3.block_public_policy or as3a.block_public_policy)
         and (s3.ignore_public_acls or as3a.ignore_public_acls)
         and (s3.restrict_public_buckets or as3a.restrict_public_buckets)
        then 'pass' 
        else 'fail' 
    end as status
from
    aws_s3_buckets s3
    join aws_s3_accounts as3a on s3.account_id = as3a.account_id
