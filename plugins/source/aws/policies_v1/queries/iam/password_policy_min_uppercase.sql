insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure IAM password policy requires at least one uppercase letter' as title,
  account_id,
  account_id,
  case when
    require_uppercase_characters is not true or policy_exists is not true
    then 'fail'
    else 'pass'
  end
from
    aws_iam_password_policies
