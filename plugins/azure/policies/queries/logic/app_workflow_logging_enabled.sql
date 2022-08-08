WITH details AS (
	SELECT DISTINCT id AS
		workflow_id
	FROM
		azure_logic_app_workflows,
		jsonb_array_elements ( diagnostic_settings ) setting,
		jsonb_array_elements ( setting -> 'properties' -> 'logs' ) log
	WHERE
		diagnostic_settings IS NOT NULL
		AND (
			( ( log ->> 'enabled' ) :: bool AND ( log -> 'retentionPolicy' ->> 'enabled' ) :: bool AND ( log -> 'retentionPolicy' ) :: jsonb ? 'days' )
			OR ( ( log ->> 'enabled' ) :: bool AND ( log -> 'retentionPolicy' ->> 'enabled' <> 'true' OR setting -> 'properties' ->> 'storageAccountId' = '' ) )
		)
	)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Resource logs in Logic Apps should be enabled',
	sub.id AS subscription_id,
	workflows.id AS logic_app_workflow_id,
  case
    when workflows.diagnostic_settings IS NULL
      then 'fail' else 'pass'
  end
FROM
	azure_logic_app_workflows
	AS workflows LEFT JOIN details AS l ON workflows.id = l.workflow_id,
	azure_subscription_subscriptions AS sub
WHERE
	sub.subscription_id = workflows.subscription_id
