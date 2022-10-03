WITH value_check AS (
    SELECT aps.cq_id, apsc.value
    FROM azure_postgresql_servers aps
        LEFT JOIN azure_postgresql_server_configurations apsc ON
                aps.cq_id = apsc.server_cq_id
    WHERE apsc."name" = 'log_retention_days'
)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure server parameter "log_retention_days" is greater than 3 days for PostgreSQL Database Server (Automated)',
  s.subscription_id,
  s.id AS server_id,
  case
    when v.value IS NULL OR v.value::INTEGER < 3
      then 'fail' else 'pass'
  end
FROM azure_postgresql_servers s
    LEFT JOIN value_check v ON
        s.cq_id = v.cq_id
