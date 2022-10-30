WITH alert_condition AS (
	SELECT
		subscription_id
	FROM
		azure_monitor_activity_log_alerts,
		jsonb_array_elements_text ( to_jsonb ( scopes ) ) SCOPE
	WHERE
		LOCATION = 'Global'
		AND enabled
		AND SCOPE = '/subscriptions/' || subscription_id
		AND _cq_id IN (
		SELECT op._cq_id
		FROM
			azure_monitor_activity_log_alerts op, jsonb_array_elements(op.condition) AS opcond,
			azure_monitor_activity_log_alerts cat, jsonb_array_elements(cat.condition) AS catcond
		WHERE
			    -- TODO check
			catcond->>'equals' = 'Administrative'
			AND catcond->>'field' = 'category'
			AND opcond->>'field' = 'operationName'
			AND opcond->>'equals' IN ( 'Microsoft.Sql/servers/firewallRules/write', 'Microsoft.Sql/servers/firewallRules/delete', 'Microsoft.Network/networkSecurityGroups/write', 'Microsoft.Network/networkSecurityGroups/delete', 'Microsoft.ClassicNetwork/networkSecurityGroups/write', 'Microsoft.ClassicNetwork/networkSecurityGroups/delete', 'Microsoft.Network/networkSecurityGroups/securityRules/write', 'Microsoft.Network/networkSecurityGroups/securityRules/delete', 'Microsoft.ClassicNetwork/networkSecurityGroups/securityRules/write', 'Microsoft.ClassicNetwork/networkSecurityGroups/securityRules/delete' )
		)
	GROUP BY
		subscription_id
	)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'An activity log alert should exist for specific Administrative operations',
	azure_subscriptions.id,
	azure_subscriptions.id
FROM
	azure_subscriptions
	LEFT JOIN alert_condition A ON azure_subscriptions.id = A.subscription_id
WHERE
	A.subscription_id IS NULL
GROUP BY
	azure_subscriptions.id,
	display_name;
