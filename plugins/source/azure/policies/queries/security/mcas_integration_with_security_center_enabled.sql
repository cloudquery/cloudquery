insert into azure_policy_results
SELECT 
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure that Microsoft Cloud App Security (MCAS) integration with Security Center is selected (Automatic)' as title,
  subscription_id,
  id,
  case
    when enabled = TRUE
    then 'pass' else 'fail'
  end
FROM azure_security_settings ass
WHERE "name" = 'MCAS'
