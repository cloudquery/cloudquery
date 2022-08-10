insert into aws_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure credentials unused for 90 days or greater are disabled (Scored)',
  account_id,
  arn,
  case when
      (password_enabled and password_last_used < (now() - '90 days'::INTERVAL)
        or (last_used < (now() - '90 days'::INTERVAL)))
      then 'fail'
      else 'pass'
  end
from
    aws_iam_users
inner join aws_iam_user_access_keys
    on aws_iam_users.cq_id = aws_iam_user_access_keys.user_cq_id
