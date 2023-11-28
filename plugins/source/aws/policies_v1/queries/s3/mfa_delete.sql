insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Ensure MFA Delete is enabled on S3 buckets (Automated)' as title,
    account_id,
    arn AS resource_id,
    case when
        versioning_status is distinct from 'Enabled'
        or versioning_mfa_delete is distinct from 'Enabled'
    then 'fail' else 'pass' end as status
from
    aws_s3_buckets
