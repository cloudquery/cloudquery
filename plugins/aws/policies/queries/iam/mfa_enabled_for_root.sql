insert into aws_policy_results
select
  :execution_time as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure MFA is enabled for the "root" account' as title,
  account_id,
  arn as resource_id,
  case 
    when user_name = '<root_account>' and not mfa_active then 'fail'
    when user_name = '<root_account>' and mfa_active then 'pass'
  end as status
from aws_iam_users
