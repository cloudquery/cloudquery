insert into aws_policy_results
select
    :execution_time as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    '' as title, -- TODO FIXME
    aws_s3_buckets.account_id,
    aws_s3_buckets.arn as resource_id,
    case when
        aws_s3_bucket_replication_rules.status is distinct from 'Enabled'
    then 'fail' else 'pass' end as status
from
    aws_s3_buckets
left join aws_s3_bucket_replication_rules on aws_s3_bucket_replication_rules.bucket_cq_id=aws_s3_buckets.cq_id

-- Note: This query doesn't validate that the destination bucket is actually in a different region
