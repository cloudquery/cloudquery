insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure that "Notify about alerts with the following severity" is set to "High" (Automated)' as title,
  subscription_id,
  id,
  case
    when email IS NOT NULL
      AND email != '' AND alert_notifications = 'On'
    then 'pass' else 'fail'
  end
FROM azure_security_contacts
