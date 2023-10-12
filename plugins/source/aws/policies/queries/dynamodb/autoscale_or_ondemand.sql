insert into aws_policy_results
SELECT
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'DynamoDB tables should automatically scale capacity with demand' AS title,
    t.account_id,
    t.arn AS resource_id,
    CASE
    WHEN t.billing_mode_summary->>'BillingMode' = 'PAY_PER_REQUEST' then 'pass'
    WHEN (t.billing_mode_summary->>'BillingMode' = 'PROVISIONED' or t.billing_mode_summary->>'BillingMode' is NULL )and (pr._cq_id IS not NULL and pw._cq_id IS not NULL) then 'pass'
    ELSE 'fail'
    END
        AS status
FROM
  aws_dynamodb_tables AS t
  LEFT JOIN aws_applicationautoscaling_policies AS pr ON
      pr.service_namespace = 'dynamodb'
      AND pr.resource_id = concat('table/', t.table_name)
	  AND pr.policy_type = 'TargetTrackingScaling'
      AND pr.scalable_dimension = 'dynamodb:table:ReadCapacityUnits'
  LEFT JOIN aws_applicationautoscaling_policies AS pw ON
      pw.service_namespace = 'dynamodb'
      AND pw.resource_id = concat('table/', t.table_name)
	  AND pw.policy_type = 'TargetTrackingScaling'
      AND pw.scalable_dimension = 'dynamodb:table:WriteCapacityUnits'
group by t.account_id, t.arn, status;