insert into aws_policy_results

with violations as (
    select
        account_id, arn, attachment_count
    from view_aws_iam_policy_statements
    where
        not(
            arn like 'arn:aws:iam::aws:policy%' or arn like 'arn:aws-us-gov:iam::aws:policy%'
        )
        and effect = 'Allow'
        AND resources ?| array['*', 'arn:aws:kms:*:' || account_id || ':key/*', 'arn:aws:kms:*:' || account_id || ':alias/*'] -- noqa
        AND actions ?| array['*', 'kms:*', 'kms:decrypt', 'kms:reencryptfrom', 'kms:reencrypt*'] -- noqa
)

select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'IAM customer managed policies should not allow decryption and re-encryption actions on all KMS keys' AS title,
    account_id,
    arn AS resource_id,
    case sum(attachment_count) 
        when 0 then 'pass' 
        else 'fail'
    end as status
from violations
group by account_id, arn
