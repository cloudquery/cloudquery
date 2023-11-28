insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'S3 Block Public Access setting should be enabled' as title,
    aws_iam_accounts.account_id,
    aws_iam_accounts.account_id AS resource_id,
    case when
        config_exists is not TRUE
        or block_public_acls is not TRUE
        or block_public_policy is not TRUE
        or ignore_public_acls is not TRUE
        or restrict_public_buckets is not TRUE
    then 'fail' else 'pass' end as status
from
    aws_iam_accounts
left join
    aws_s3_accounts on
        aws_iam_accounts.account_id = aws_s3_accounts.account_id
