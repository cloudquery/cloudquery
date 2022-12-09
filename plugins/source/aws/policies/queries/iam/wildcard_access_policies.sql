insert into aws_policy_results

with policy_statements as (
    select
        aws_iam_policies.id,
        JSONB_ARRAY_ELEMENTS(
            case JSONB_TYPEOF(((v->>'Document')::jsonb) -> 'Statement')
                when
                    'string' then JSONB_BUILD_ARRAY(
                        ((v->>'Document')::jsonb) ->> 'Statement'
                    )
                when
                    'array' then ((v->>'Document')::jsonb) -> 'Statement'
            end
        ) as statement
    from
        aws_iam_policies, jsonb_array_elements(aws_iam_policies.policy_version_list) AS v
    where aws_iam_policies.arn not like 'arn:aws:iam::aws:policy%'
),

allow_all_statements as (
    select
        id,
        COUNT(statement) as statements_count
    from
        policy_statements
    where
        statement ->> 'Effect' = 'Allow'
        and (
            statement ->> 'Action' like '%*%'
            or statement ->> 'NotAction' like '%*%')
    group by
        id
)

select distinct
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'IAM customer managed policies that you create should not allow wildcard actions for services' as title,
    aws_iam_policies.account_id,
    aws_iam_policies.arn AS resource_id,
    CASE WHEN statements_count > 0 THEN 'fail' ELSE 'pass' END AS status
from aws_iam_policies
     left join
     allow_all_statements on aws_iam_policies.id = allow_all_statements.id
