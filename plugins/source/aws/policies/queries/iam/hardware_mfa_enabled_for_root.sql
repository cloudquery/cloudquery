insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure hardware MFA is enabled for the "root" account (Scored)' as title,
  split_part(cr.arn, ':', 5) as account_id,
  cr.arn as resource_id,
  case
    when mfa.serial_number is null or cr.mfa_active = FALSE then 'fail'
    when mfa.serial_number is not null and cr.mfa_active = TRUE then 'pass'
  end as status
from aws_iam_credential_reports cr
left join
    aws_iam_virtual_mfa_devices mfa on
        mfa.user->>'Arn' = cr.arn
where cr.user = '<root_account>'
group by mfa.serial_number, cr.mfa_active, cr.arn;
