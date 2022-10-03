insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure that Windows Defender ATP (WDATP) integration with Security Center is selected (Automatic)',
  subscription_id,
  id,
  case
    when pricing_tier = 'Standard'
    then 'pass' else 'fail'
  end
FROM azure_security_pricings asp
WHERE "name" = 'KeyVaults'
