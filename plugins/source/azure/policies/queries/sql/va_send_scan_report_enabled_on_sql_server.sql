WITH vulnerability_emails AS (
    SELECT
        id,
        UNNEST(recurring_scans_emails) AS emails
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
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure that VA setting Send scan reports to is configured for a SQL server (Automated)',
  s.subscription_id,
  s.id AS server_id,
  case
    when c.emails_number = 0 OR c.emails_number IS NULL
      then 'fail' else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_server_vulnerability_assessments sv ON
        s.cq_id = sv.server_cq_id
    LEFT JOIN emails_count c ON
        sv.id = c.id
