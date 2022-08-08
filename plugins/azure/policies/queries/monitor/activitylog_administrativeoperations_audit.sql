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
		AND cq_id IN (
		SELECT OPERATION
			.activity_log_alert_cq_id
		FROM
			azure_monitor_activity_log_alert_conditions
			OPERATION JOIN azure_monitor_activity_log_alert_conditions category ON category.activity_log_alert_cq_id = OPERATION.activity_log_alert_cq_id
		WHERE
			category.EQUALS = 'Administrative'
			AND category.field = 'category'
			AND OPERATION.field = 'operationName'
			AND OPERATION.EQUALS IN ( 'Microsoft.Sql/servers/firewallRules/write', 'Microsoft.Sql/servers/firewallRules/delete', 'Microsoft.Network/networkSecurityGroups/write', 'Microsoft.Network/networkSecurityGroups/delete', 'Microsoft.ClassicNetwork/networkSecurityGroups/write', 'Microsoft.ClassicNetwork/networkSecurityGroups/delete', 'Microsoft.Network/networkSecurityGroups/securityRules/write', 'Microsoft.Network/networkSecurityGroups/securityRules/delete', 'Microsoft.ClassicNetwork/networkSecurityGroups/securityRules/write', 'Microsoft.ClassicNetwork/networkSecurityGroups/securityRules/delete' )
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
	azure_subscription_subscriptions.subscription_id,
	azure_subscription_subscriptions.subscription_id
FROM
	azure_subscription_subscriptions
	LEFT JOIN alert_condition A ON azure_subscription_subscriptions.subscription_id = A.subscription_id
WHERE
	A.subscription_id IS NULL
GROUP BY
	azure_subscription_subscriptions.subscription_id,
	display_name;
