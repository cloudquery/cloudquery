insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'IAM users'' access keys should be rotated every 90 days or less' AS title,
    account_id,
    access_key_id AS resource_id,
    case when date_part('day', now() - last_rotated) > 90 then 'fail'
         else 'pass'
    end as status
from aws_iam_user_access_keys
