insert into aws_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure IAM password policy requires at least one number',
  account_id,
  account_id,
  case when
    require_numbers = FALSE or policy_exists = FALSE
    then 'fail'
    else 'pass'
  end as status
from
    aws_iam_password_policies

