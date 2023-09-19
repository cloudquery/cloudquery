WITH dynamodb_tables AS (
    SELECT
        account_id,
        arn,
        table_name,
        billing_mode_summary->>'BillingMode' IS DISTINCT FROM 'PAY_PER_REQUEST' AS is_not_on_demand
    FROM aws_dynamodb_tables
),
replica_auto_scalings AS (
    SELECT
        s.table_arn,
        s.region,
        s.region_name,
        (s.replica_provisioned_read_capacity_auto_scaling_settings->>'AutoScalingDisabled')::BOOLEAN AS read_auto_scaling_disabled,
        (s.replica_provisioned_write_capacity_auto_scaling_settings->>'AutoScalingDisabled')::BOOLEAN AS write_auto_scaling_disabled
    FROM aws_dynamodb_table_replica_auto_scalings s
),
auto_scaling_disabled_count_in_replica_a_s AS (
    SELECT
        table_arn,
        region,
        region_name,
        CASE WHEN (read_auto_scaling_disabled IS NOT NULL AND read_auto_scaling_disabled IS TRUE) THEN 1 ELSE 0 END AS read_auto_scaling_disabled_count,
        CASE WHEN (write_auto_scaling_disabled IS NOT NULL AND write_auto_scaling_disabled IS TRUE) THEN 1 ELSE 0 END AS write_auto_scaling_disabled_count
    FROM replica_auto_scalings
),
sum_of_auto_scaling_disabled_count_in_replica_a_s AS (
    SELECT
        table_arn,
        SUM(read_auto_scaling_disabled_count + write_auto_scaling_disabled_count) AS auto_scaling_disabled_count
    FROM auto_scaling_disabled_count_in_replica_a_s
    GROUP BY table_arn
),
dynamodb_tables_with_replica_auto_scaling_disabled_count AS (
    SELECT
        account_id,
        arn,
        table_name,
        is_not_on_demand,
        CASE WHEN (auto_scaling_disabled_count IS NULL) THEN 0 ELSE auto_scaling_disabled_count END AS replica_auto_scaling_disabled_count
    FROM dynamodb_tables AS t1
    LEFT JOIN sum_of_auto_scaling_disabled_count_in_replica_a_s AS t2 ON t1.arn = t2.table_arn
),
policies_r AS (
    SELECT resource_id
    FROM aws_applicationautoscaling_policies
    WHERE (service_namespace = 'dynamodb'
        AND policy_type = 'TargetTrackingScaling'
        AND scalable_dimension = 'dynamodb:table:ReadCapacityUnits')
    GROUP BY resource_id
),
policies_w AS (
    SELECT resource_id
    FROM aws_applicationautoscaling_policies
    WHERE (service_namespace = 'dynamodb'
        AND policy_type = 'TargetTrackingScaling'
        AND scalable_dimension = 'dynamodb:table:WriteCapacityUnits')
    GROUP BY resource_id
)
INSERT INTO aws_policy_results
SELECT
    :'execution_time' AS execution_time,
    :'framework' AS framework,
    :'check_id' AS check_id,
    'DynamoDB tables should automatically scale capacity with demand' AS title,
    t.account_id,
    t.arn AS resource_id,
    CASE
        WHEN (t.is_not_on_demand IS TRUE
            AND t.replica_auto_scaling_disabled_count > 0) THEN 'fail'
        WHEN (t.is_not_on_demand IS TRUE
            AND t.replica_auto_scaling_disabled_count = 0
            AND (pr.resource_id IS NULL OR pw.resource_id IS NULL)) THEN 'fail'
        ELSE 'pass'
    END AS status
FROM dynamodb_tables_with_replica_auto_scaling_disabled_count AS t
LEFT JOIN policies_r pr ON (pr.resource_id = CONCAT('table/', t.table_name))
LEFT JOIN policies_w pw ON (pw.resource_id = CONCAT('table/', t.table_name))