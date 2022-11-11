insert into aws_policy_results

with iam_policies as (
    select
        (v->>'Document')::jsonb AS document,
        account_id,
        arn,
        id
    from aws_iam_policies, jsonb_array_elements(aws_iam_policies.policy_version_list) AS v
),

violations as (
    select distinct id
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
        violations.id is not null
    then 'fail' else 'pass' end as status
from aws_iam_policies
left join violations on violations.id = aws_iam_policies.id
