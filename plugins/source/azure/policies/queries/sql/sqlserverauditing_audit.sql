insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Auditing on SQL server should be enabled',
  azure_subscription_subscriptions.id,
	azure_subscription_subscriptions.display_name AS subscription_name,
	case
    when azure_sql_server_db_blob_auditing_policies.server_cq_id = azure_sql_servers.cq_id
	    AND azure_subscription_subscriptions.subscription_id = azure_sql_servers.subscription_id
	    AND azure_sql_server_db_blob_auditing_policies.state = 'Disabled'
    then 'fail' else 'pass'
  end
FROM
	azure_sql_server_db_blob_auditing_policies,
	azure_sql_servers,
	azure_subscription_subscriptions