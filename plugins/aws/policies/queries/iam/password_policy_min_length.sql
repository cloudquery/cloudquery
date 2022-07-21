insert into aws_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure IAM password policy requires minimum length of 14 or greater',
  account_id,
  account_id,
  case when
    (minimum_password_length < 14) or policy_exists = FALSE
    then 'fail'
    else 'pass'
  end
from
    aws_iam_password_policies
