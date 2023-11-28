insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure IAM password policy requires at least one lowercase letter' as title,
  account_id,
  account_id,
  case when
    require_lowercase_characters = false or policy_exists = false
    then 'fail'
    else 'pass'
  end as status
from
    aws_iam_password_policies
    
