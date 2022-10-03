insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure "Additional email addresses" is configured with a security contact email (Automated)',
  subscription_id,
  id,
  case
    when email IS NOT NULL
      AND email != ''
    then 'pass' else 'fail'
  end
FROM azure_security_contacts
