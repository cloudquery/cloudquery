insert into aws_policy_results

with iam_policies as (
    select
        document,
        account_id,
        arn,
        aws_iam_policies.cq_id
    from aws_iam_policy_versions
    inner join
        aws_iam_policies on
            aws_iam_policies.cq_id = aws_iam_policy_versions.policy_cq_id
),

violations as (
    select distinct cq_id
    from iam_policies,
        jsonb_array_elements(
            case jsonb_typeof(document -> 'Statement')
                when 'string' then jsonb_build_array(document ->> 'Statement')
                when 'array' then document -> 'Statement'
            end
        ) as statement
    where
        not(
            arn like 'arn:aws:iam::aws:policy%' or arn like 'arn:aws-us-gov:iam::aws:policy%'
        )
        and statement ->> 'Effect' = 'Allow'
        AND statement -> 'Resource'?| array['*', 'arn:aws:kms:*:' || account_id || ':key/*', 'arn:aws:kms:*:' || account_id || ':alias/*'] -- noqa
        AND statement -> 'Action' ?| array['*', 'kms:*', 'kms:decrypt', 'kms:reencryptfrom', 'kms:reencrypt*'] -- noqa
)

select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'IAM customer managed policies should not allow decryption and re-encryption actions on all KMS keys' AS title,
    account_id,
    arn AS resource_id,
    case when
        violations.cq_id is not null
    then 'fail' else 'pass' end as status
from aws_iam_policies
left join violations on violations.cq_id = aws_iam_policies.cq_id
