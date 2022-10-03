insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure that "Notify about alerts with the following severity" is set to "High" (Automated)',
  subscription_id,
  id,
  case
    when email IS NOT NULL
      AND email != '' AND alert_notifications = 'On'
    then 'pass' else 'fail'
  end
FROM azure_security_contacts
