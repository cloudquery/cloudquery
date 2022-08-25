insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Unused IAM user credentials should be removed' AS title,
    account_id,
    aws_iam_user_access_keys.access_key_id AS resource_id,
    case when date_part('day', now() - aws_iam_user_access_keys.last_used_date) > 90 then 'fail'
        else 'pass'
    end as status
from aws_iam_users
     left join
     aws_iam_user_access_keys on
     aws_iam_users.cq_id = aws_iam_user_access_keys.user_cq_id
where access_key_id is not null