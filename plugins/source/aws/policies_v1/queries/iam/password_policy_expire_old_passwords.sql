insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure IAM password policy expires passwords within 90 days or less' as title,
  account_id,
  account_id,
  case when
    (max_password_age is null or max_password_age > 90) or policy_exists = false
    then 'fail'
    else 'pass'
  end
from
    aws_iam_password_policies
