insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Subnets should be associated with a Network Security Group',
  sub.id,
	sg.id,
  case
    when subnet.id IS NULL then 'fail' else 'pass'
  end
FROM
	azure_network_security_groups AS sg
	JOIN azure_subscription_subscriptions AS sub ON sub.subscription_id = sg.subscription_id
	LEFT JOIN azure_network_virtual_network_subnets AS subnet ON subnet.network_security_group_id = sg.id
