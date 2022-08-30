insert into aws_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure credentials unused for 45 days or greater are disabled (Automated)',
  account_id,
  arn,
  case when
      (password_enabled and password_last_used < (now() - '45 days'::INTERVAL))
       or  (password_enabled and password_last_used is null and password_last_changed < (now() - '45 days'::INTERVAL))
      -- todo add access_key_1_last_used_date access_key_2_last_used_date to iam_users table
--        or  (access_key_1_active and access_key_1_last_used_date  < (now() - '45 days'::INTERVAL))
--        or  (access_key_1_active and access_key_1_last_used_date is null and access_key_1_last_rotated < (now() - '45 days'::INTERVAL))
--        or  (access_key_2_active and access_key_2_last_used_date  < (now() - '45 days'::INTERVAL))
--        or  (access_key_2_active and access_key_2_last_used_date is null and access_key_2_last_rotated < (now() - '45 days'::INTERVAL))
  then 'fail'
  else 'pass'
  end
from
    aws_iam_users
inner join aws_iam_user_access_keys
    on aws_iam_users.cq_id = aws_iam_user_access_keys.user_cq_id
