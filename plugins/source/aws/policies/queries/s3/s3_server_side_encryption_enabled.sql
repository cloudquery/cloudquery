insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'S3 buckets should have server-side encryption enabled' as title,
    aws_s3_buckets.account_id,
    arn as resource_id,
    case when
        aws_s3_bucket_encryption_rules.bucket_arn is null
    then 'fail' else 'pass' end as status
from
    aws_s3_buckets
left join aws_s3_bucket_encryption_rules on aws_s3_bucket_encryption_rules.bucket_arn=aws_s3_buckets.arn

-- Note: This query doesn't validate if a bucket policy requires encryption for `put-object` requests
