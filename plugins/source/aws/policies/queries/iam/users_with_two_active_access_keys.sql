insert into aws_policy_results
select
    :'execution_time',
    :'framework',
    :'check_id',
    'Ensure there is only one active access key available for any single IAM user (Automated)',
    account_id,
    user_name,
    case when
         access_key_1_active = TRUE
            and access_key_2_active = TRUE
         then 'fail'
         else 'pass'
    end
from
    aws_iam_users
