insert into azure_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'App Service should use a virtual network service endpoint',
  awa.subscription_id AS subscription_id,
  awa.id,
  case
    when vnet.properties -> 'vnetResourceId' is null
    then 'fail' else 'pass'
  end
from
    azure_appservice_web_apps awa
left join
    azure_appservice_web_app_vnet_connections as vnet
on vnet._cq_parent_id = awa._cq_id
