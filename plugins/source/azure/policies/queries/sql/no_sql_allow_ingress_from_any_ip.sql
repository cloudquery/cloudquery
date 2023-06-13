insert into azure_policy_results
SELECT
    :'execution_time'                                          AS execution_time,
    :'framework'                                               AS framework,
    :'check_id'                                                AS check_id,
    'Ensure no SQL Databases allow ingress 0.0.0.0/0 (ANY IP)' AS title,
    subscription_id                                            AS subscription_id,
    id                                                         AS resource_id,
    CASE
        WHEN properties->>'startIpAddress' = '0.0.0.0'
                 AND (properties->>'endIpAddress' = '0.0.0.0' OR properties->>'endIpAddress' = '255.255.255.255')
        THEN 'fail'
        ELSE 'pass'
    END                                                        AS status
FROM azure_sql_server_firewall_rules
