insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure "Additional email addresses" is configured with a security contact email (Automated)' as title,
  subscription_id,
  id,
  case
    when email IS NOT NULL
      AND email != ''
    then 'pass' else 'fail'
  end
FROM azure_security_contacts
