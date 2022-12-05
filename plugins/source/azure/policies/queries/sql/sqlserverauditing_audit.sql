insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Auditing on SQL server should be enabled' as title,
  azure_subscriptions.id,
	azure_subscriptions.display_name AS subscription_name,
	case
    when azure_sql_server_blob_auditing_policies.sql_server_id = azure_sql_servers.id
	    AND azure_subscriptions.id = azure_sql_servers.subscription_id
	    AND azure_sql_server_blob_auditing_policies.state = 'Disabled'
    then 'fail' else 'pass'
  end
FROM
    azure_sql_server_blob_auditing_policies,
	azure_sql_servers,
	azure_subscriptions