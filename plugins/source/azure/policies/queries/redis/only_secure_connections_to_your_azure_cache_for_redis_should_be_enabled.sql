insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Only secure connections to your Azure Cache for Redis should be enabled',
  subscription_id,
  id,
  case
    when (properties ->> 'enableNonSslPort')::boolean IS NOT FALSE
    then 'fail' else 'pass'
  end
FROM azure_redis_caches

