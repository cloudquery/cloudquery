insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'IAM principals should not have IAM inline policies that allow decryption and re-encryption actions on all KMS keys' AS title,
    account_id,
    arn AS resource_id,
    'fail' AS status -- TODO FIXME
from
    (
        -- select all user policies
        select
            statement,
            aws_iam_users.account_id,
            arn,
            policy_name
        from aws_iam_user_policies
        cross join lateral jsonb_array_elements(
                case jsonb_typeof(policy_document -> 'Statement')
                    when
                        'string' then jsonb_build_array(
                            policy_document ->> 'Statement'
                        )
                    when 'array' then policy_document -> 'Statement' end
        ) as statement
        inner join
            aws_iam_users on
                aws_iam_users.account_id = aws_iam_user_policies.account_id and
                aws_iam_users.arn = aws_iam_user_policies.user_arn
        union
        -- select all role policies
        select
            statement,
            aws_iam_roles.account_id,
            arn,
            policy_name
        from aws_iam_role_policies
             cross join lateral jsonb_array_elements(
                case jsonb_typeof(policy_document -> 'Statement')
                    when
                        'string' then jsonb_build_array(
                            policy_document ->> 'Statement'
                        )
                    when 'array' then policy_document -> 'Statement' end
        ) as statement
        inner join
            aws_iam_roles on
                aws_iam_roles.account_id = aws_iam_role_policies.account_id and
                aws_iam_roles.arn = aws_iam_role_policies.role_arn
        where lower(arn) not like 'arn:aws:iam::%:role/aws-service-role/%'
        union
        -- select all group policies
        select
            statement,
            aws_iam_groups.account_id,
            arn,
            policy_name
        from aws_iam_group_policies
        cross join lateral jsonb_array_elements(
                case jsonb_typeof(policy_document -> 'Statement')
                    when
                        'string' then jsonb_build_array(
                            policy_document ->> 'Statement'
                        )
                    when 'array' then policy_document -> 'Statement' end
        ) as statement
        inner join aws_iam_groups on
            aws_iam_groups.account_id = aws_iam_group_policies.account_id and
            aws_iam_groups.arn = aws_iam_group_policies.group_arn
        ) as t
where
    statement ->> 'Effect' = 'Allow'
    and lower(statement::TEXT)::JSONB -> 'resource' ?| array[
        '*',
        'arn:aws:kms:*:*:key/*',
        'arn:aws:kms:*:' || account_id || ':key/*'
        'arn:aws:kms:*:*:alias/*',
        'arn:aws:kms:*:' || account_id || ':alias/*'
    ]

    and lower(statement::TEXT)::JSONB -> 'action' ?| array[
        '*',
        'kms:*',
        'kms:decrypt',
        'kms:encrypt*',
        'kms:reencryptfrom'
    ]
