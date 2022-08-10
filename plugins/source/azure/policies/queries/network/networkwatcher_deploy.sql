insert into azure_policy_results
SELECT
  :'execution_time'
  :'framework',
  :'check_id',
  '',
  sub.id AS subscription_id
	azure_network_virtual_networks.id,
	case
    when azure_network_watchers.location IS NULL
	    OR LOWER ( split_part( azure_network_watchers.id, '/', 5 ) ) != 'networkwatcherrg'
    then 'fail' else 'pass'
  end
FROM
	azure_network_virtual_networks
	LEFT JOIN azure_network_watchers ON azure_network_virtual_networks.location = azure_network_watchers.location
	JOIN azure_subscription_subscriptions sub ON sub.subscription_id = azure_network_virtual_networks.subscription_id
