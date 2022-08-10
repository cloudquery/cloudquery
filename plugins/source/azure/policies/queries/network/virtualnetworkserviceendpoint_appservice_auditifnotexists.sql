insert into azure_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'App Service should use a virtual network service endpoint',
  azure_subscription_subscriptions.id AS subscription_id,
  azure_web_apps.id,
  case
    when vnet_connection -> 'properties' -> 'vnetResourceId' is null
    then 'fail' else 'pass'
  end
from
    azure_web_apps,
    azure_subscription_subscriptions
where
  azure_subscription_subscriptions.subscription_id = azure_web_apps.subscription_id
