insert into aws_policy_results

with iam_policies as (
    select
        id,
        (v->>'Document')::jsonb AS document
    from aws_iam_policies, jsonb_array_elements(aws_iam_policies.policy_version_list) AS v
    where aws_iam_policies.default_version_id = v->>'VersionId' and arn not like 'arn:aws:iam::aws:policy%'
),
policy_statements as (
    select
        id,
        JSONB_ARRAY_ELEMENTS(
            case JSONB_TYPEOF(document -> 'Statement')
                when
                    'string' then JSONB_BUILD_ARRAY(
                        document ->> 'Statement'
                    )
                when 'array' then document -> 'Statement' end
        ) as statement
    from
        iam_policies
),
allow_all_statements as (
    select
        id,
        COUNT(statement) as statements_count
    from policy_statements
    where (statement ->> 'Action' = '*'
        or statement ->> 'Action' like '%"*"%')
        and statement ->> 'Effect' = 'Allow'
        and (statement ->> 'Resource' = '*'
            or statement ->> 'Resource' like '%"*"%')
    group by id
)

select distinct
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'IAM policies should not allow full ''*'' administrative privileges' as title,
    aws_iam_policies.account_id,
    aws_iam_policies.arn AS resource_id,
    CASE WHEN statements_count > 0 THEN 'fail' ELSE 'pass' END AS status
from aws_iam_policies
left join
    allow_all_statements on aws_iam_policies.id = allow_all_statements.id
