insert into azure_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'App Service should use a virtual network service endpoint',
  awa.subscription_id AS subscription_id,
  awa.id,
  case
    when vnet_resource_id is null
    then 'fail' else 'pass'
  end
from
    azure_web_apps awa
left join
    azure_web_vnet_connections as vnet
on vnet.web_app_id = awa.id
