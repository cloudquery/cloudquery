WITH
    ds AS (
        SELECT resource_id as logic_workflow_id, jsonb_array_elements(properties -> 'logs') AS logs FROM azure_monitor_diagnostic_settings
    ),
    details AS (
	SELECT DISTINCT id AS
		workflow_id
	FROM
        azure_logic_workflows w
    LEFT JOIN ds ON ds.logic_workflow_id = w.id
	WHERE
        (ds.logs->>'enabled')::boolean IS TRUE AND
        (ds.logs->'retentionPolicy'->>'enabled')::boolean IS TRUE)
-- TODO check
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Resource logs in Logic Apps should be enabled',
	subscription_id,
	workflows.id AS logic_app_workflow_id,
  case
    when l.workflow_id IS NULL
      then 'fail' else 'pass'
  end
FROM
    azure_logic_workflows
	AS workflows LEFT JOIN details AS l ON workflows.id = l.workflow_id
