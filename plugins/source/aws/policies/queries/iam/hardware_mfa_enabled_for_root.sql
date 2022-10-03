insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure hardware MFA is enabled for the "root" account (Scored)' as title,
  aws_iam_users.account_id,
  arn as resource_id,
  case
    when aws_iam_users.user_name = '<root_account>' and (
        serial_number is null or aws_iam_users.mfa_active = FALSE) then 'fail'
    when aws_iam_users.user_name = '<root_account>' and 
        serial_number is not null and aws_iam_users.mfa_active = FALSE then 'pass'
  end as status
from aws_iam_users
left join
    aws_iam_virtual_mfa_devices on
        aws_iam_virtual_mfa_devices.user_arn = aws_iam_users.arn
where aws_iam_users.user_name = '<root_account>'
group by aws_iam_users.account_id, aws_iam_users.user_name, serial_number, mfa_active, arn
