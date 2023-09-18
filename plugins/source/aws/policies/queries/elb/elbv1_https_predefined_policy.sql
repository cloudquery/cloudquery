INSERT INTO aws_policy_results
WITH flat_listeners AS (
    SELECT
        arn,
        account_id,
        li->'Listener'->>'Protocol' as protocol,
        li->'PolicyNames' as policies_arr
    FROM aws_elbv1_load_balancers lb, jsonb_array_elements(lb.listener_descriptions) AS li
),
violations AS (
SELECT 
    fl.arn,
    fl.account_id,
    CASE 
        WHEN fl.protocol IN ('HTTPS', 'SSL')
          AND NOT EXISTS (
            SELECT 1
            FROM aws_elbv1_load_balancer_policies pol
            WHERE fl.policies_arr @> ('["' || pol.policy_name || '"]')::jsonb
            AND pol.policy_attribute_descriptions->>'Reference-Security-Policy' = 'ELBSecurityPolicy-TLS-1-2-2017-01'
        ) THEN 1
        ELSE 0
    END AS flag
FROM flat_listeners fl
)
SELECT
  DISTINCT
    :'execution_time' AS execution_time,
    :'framework' AS framework,
    :'check_id' AS check_id,
    'Classic Load Balancers with HTTPS/SSL listeners should use a predefined security policy that has strong configuration' AS title,
    v.account_id,
    v.arn AS resource_id,
    CASE
      WHEN MAX(flag) OVER(PARTITION BY arn) = 1 THEN 'fail'
      ELSE 'pass'
    END as status
FROM violations v