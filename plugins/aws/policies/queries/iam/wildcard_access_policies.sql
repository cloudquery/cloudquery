insert into aws_policy_results

with policy_statements as (
    select
        aws_iam_policies.cq_id as cq_id,
        JSONB_ARRAY_ELEMENTS(
            case JSONB_TYPEOF(aws_iam_policy_versions.document -> 'Statement')
                when
                    'string' then JSONB_BUILD_ARRAY(
                        aws_iam_policy_versions.document ->> 'Statement'
                    )
                when
                    'array' then aws_iam_policy_versions.document -> 'Statement'
            end
        ) as statement
    from
        aws_iam_policies
    left join aws_iam_policy_versions
        on
            aws_iam_policies.cq_id = aws_iam_policy_versions.policy_cq_id and aws_iam_policies.default_version_id = aws_iam_policy_versions.version_id
    where aws_iam_policies.arn not like 'arn:aws:iam::aws:policy%'
),

allow_all_statements as (
    select
        cq_id,
        COUNT(statement) as statements_count
    from
        policy_statements
    where
        statement ->> 'Effect' = 'Allow'
        and (
            statement ->> 'Action' like '%*%'
            or statement ->> 'NotAction' like '%*%')
    group by
        cq_id
)

select distinct
    :'execution_time'::timestamp,
    :'framework',
    :'check_id',
    'IAM customer managed policies that you create should not allow wildcard actions for services' AS title,
    aws_iam_policies.account_id,
    aws_iam_policies.arn AS resource_id,
    CASE WHEN statements_count > 0 THEN 'fail' ELSE 'pass' END AS status
from aws_iam_policies
     left join
     allow_all_statements on aws_iam_policies.cq_id = allow_all_statements.cq_id
