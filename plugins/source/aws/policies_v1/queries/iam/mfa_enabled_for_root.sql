insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure MFA is enabled for the "root" account' as title,
  split_part(arn, ':', 5) as account_id,
  arn as resource_id,
  case
    when user = '<root_account>' and not mfa_active then 'fail' -- TODO check
    when user = '<root_account>' and mfa_active then 'pass'
  end as status
from aws_iam_credential_reports
