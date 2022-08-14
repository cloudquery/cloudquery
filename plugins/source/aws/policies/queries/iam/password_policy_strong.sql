insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Password policies for IAM users should have strong configurations' as title,
    account_id,
    account_id AS resource_id,
    case when
        (
            require_uppercase_characters is not TRUE
            or require_lowercase_characters is not TRUE
            or require_numbers is not TRUE
            or minimum_password_length < 14
            or password_reuse_prevention is null
            or max_password_age is null
            or policy_exists is not TRUE
        )
    then 'fail' else 'pass' end as status
from aws_iam_password_policies
