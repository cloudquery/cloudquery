WITH value_check AS (
    SELECT aps._cq_id, apsc.value
    FROM azure_postgresql_servers aps
        LEFT JOIN azure_postgresql_configurations apsc ON
            aps.id = apsc.postgresql_server_id
    WHERE apsc."name" = 'log_connections'
)
insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure server parameter "log_connections" is set to "ON" for PostgreSQL Database Server (Automated)' as title,
  s.subscription_id,
  s.id AS server_id,
  case
    when v.value IS NULL OR v.value != 'on'
      then 'fail' else 'pass'
  end
FROM azure_postgresql_servers s
    LEFT JOIN value_check v ON
        s._cq_id = v._cq_id
