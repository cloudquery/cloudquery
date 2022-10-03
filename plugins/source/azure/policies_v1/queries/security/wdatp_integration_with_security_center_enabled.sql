SELECT
  :'execution_time'
  :'framework',
  :'check_id',
  '',
  subscription_id,
  id,
  case
    when "name" = 'WDATP'
      AND enabled = TRUE
    then 'fail' else 'pass'
  end
FROM azure_security_settings
