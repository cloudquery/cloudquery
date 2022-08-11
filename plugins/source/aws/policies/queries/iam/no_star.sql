insert into aws_policy_results

with violations as (
    select
        policy_cq_id,
        COUNT(*) as violations
    from aws_iam_policy_versions,
        JSONB_ARRAY_ELEMENTS(
            case JSONB_TYPEOF(document -> 'Statement')
                when 'string' then JSONB_BUILD_ARRAY(document ->> 'Statement')
                when 'array' then document -> 'Statement'
            end
        ) as statement,
        JSONB_ARRAY_ELEMENTS_TEXT(
            case JSONB_TYPEOF(statement -> 'Resource')
                when 'string' then JSONB_BUILD_ARRAY(statement ->> 'Resource')
                when 'array' then statement -> 'Resource' end
        ) as resource,
        JSONB_ARRAY_ELEMENTS_TEXT( case JSONB_TYPEOF(statement -> 'Action')
                when 'string' then JSONB_BUILD_ARRAY(statement ->> 'Action')
                when 'array' then statement -> 'Action' end
        ) as action
    where statement ->> 'Effect' = 'Allow'
          and resource = '*'
          and ( action = '*' or action = '*:*' )
    group by policy_cq_id
)

select distinct
    :'execution_time'::timestamp as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'IAM policies should not allow full ''*'' administrative privileges' as title,
    account_id,
    arn AS resource_id,
    case when
        violations.policy_cq_id is not null AND violations.violations > 0
    then 'fail' else 'pass' end as status
from aws_iam_policies
left join violations on violations.policy_cq_id = aws_iam_policies.cq_id
