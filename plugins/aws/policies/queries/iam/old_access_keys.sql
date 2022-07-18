insert into aws_policy_results
select
  :execution_time,
  :'framework',
  :'check_id',
  'Ensure access keys are rotated every 90 days or less',
  account_id,
  arn,
  case when
    last_rotated < (now() - '90 days'::INTERVAL)
    then 'fail'
    else 'pass'
  end
from aws_iam_users
inner join
    aws_iam_user_access_keys on
        aws_iam_users.cq_id = aws_iam_user_access_keys.user_cq_id
