WITH ad_admins_count AS ( SELECT ass._cq_id, count(*) AS admins_count
    FROM azure_sql_servers ass
        LEFT JOIN azure_sql_server_admins assa ON
            ass.id = assa.sql_server_id WHERE assa.administrator_type = 'ActiveDirectory' GROUP BY ass._cq_id,
        assa.administrator_type
)
insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure that Azure Active Directory Admin is configured (Automated)' as title,
  s.subscription_id,
  s.id,
  case
    when a.admins_count IS NULL
      OR a.admins_count = 0
    then 'fail' else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN ad_admins_count a ON
        s._cq_id = a._cq_id
