insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'DynamoDB tables should automatically scale capacity with demand' as title,
    t.account_id,
    pr.arn as resource_id,
  case when
    t.billing_mode_summary->>'BillingMode' IS DISTINCT FROM 'PAY_PER_REQUEST'
    AND (
        (s.replica_provisioned_read_capacity_auto_scaling_settings->>'AutoScalingDisabled')::boolean IS DISTINCT FROM FALSE
        OR (s.replica_provisioned_write_capacity_auto_scaling_settings->>'AutoScalingDisabled')::boolean IS DISTINCT FROM FALSE
    )
    AND (pr._cq_id IS NULL OR pw._cq_id IS NULL)
    then 'fail'
    else 'pass'
  end as status
FROM aws_dynamodb_tables t
    LEFT JOIN aws_dynamodb_table_replica_auto_scalings s ON s.table_arn = t.arn
    LEFT JOIN aws_applicationautoscaling_policies pr ON (pr.service_namespace = 'dynamodb'
        AND pr.resource_id = CONCAT('table/', t.table_name)
        AND pr.policy_type = 'TargetTrackingScaling'
        AND pr.scalable_dimension = 'dynamodb:table:ReadCapacityUnits')
    LEFT JOIN aws_applicationautoscaling_policies pw ON (pw.service_namespace = 'dynamodb'
        AND pw.resource_id = CONCAT('table/', t.table_name)
        AND pw.policy_type = 'TargetTrackingScaling'
        AND pw.scalable_dimension = 'dynamodb:table:WriteCapacityUnits')
