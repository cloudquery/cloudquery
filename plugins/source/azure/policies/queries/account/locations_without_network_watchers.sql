insert into azure_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Network Watcher should be enable' as title,
  l.subscription_id,
  l.id,
  case
    when anw._cq_id is null then 'fail' else 'pass'
  end
from azure_subscription_subscription_locations l
  left join azure_network_watchers anw on l.name = anw.location
