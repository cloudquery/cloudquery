insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure IAM password policy prevents password reuse' as title,
  account_id,
  account_id,
  case when
    password_reuse_prevention is distinct from 24
    then 'fail'
    else 'pass'
  end
from
    aws_iam_password_policies
