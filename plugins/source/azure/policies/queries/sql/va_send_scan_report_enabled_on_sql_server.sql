WITH vulnerability_emails AS (
    SELECT
        id,
        UNNEST((v.properties->'recurringScans'->>'emails')::text[]) AS emails
    FROM azure_sql_server_vulnerability_assessments v
),
emails_count AS (
    SELECT
        id,
        count(emails) AS emails_number
    FROM vulnerability_emails
    GROUP BY id
)
insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure that VA setting Send scan reports to is configured for a SQL server (Automated)' as title,
  s.subscription_id,
  s.id AS server_id,
  case
    when c.emails_number = 0 OR c.emails_number IS NULL
      then 'fail' else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_server_vulnerability_assessments sv ON
        s._cq_id = sv._cq_parent_id
    LEFT JOIN emails_count c ON
        sv.id = c.id
