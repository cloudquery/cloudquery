INSERT INTO aws_policy_results
WITH pass_resources AS (
    SELECT act.account_id, arn
    FROM aws_cloudtrail_trails act
        JOIN aws_cloudtrail_trail_event_selectors actes ON act._cq_id = actes._cq_parent_id, jsonb_array_elements(actes.data_resources) AS data_resource
    WHERE is_multi_region_trail
      AND data_resource->>'Type' = 'AWS::S3::Object'
      AND data_resource->'Values' ? 'arn:aws:s3'
      AND ( read_write_type = 'WriteOnly' OR read_write_type = 'All' )
    GROUP BY act.account_id, arn
)
SELECT
    :'execution_time' AS execution_time,
    :'framework' AS framework,
    :'check_id' AS check_id,
    'Ensure that Object-level logging for write events is enabled for S3 bucket' AS title,
    aws_cloudtrail_trails.account_id AS account_id,
    aws_cloudtrail_trails.arn AS resource_id,
    CASE
        WHEN bool_or(pass_resources.arn IS NULL)
        THEN 'fail'
        ELSE 'pass'
    END AS status
FROM aws_cloudtrail_trails
    LEFT JOIN pass_resources ON aws_cloudtrail_trails.arn = pass_resources.arn
GROUP BY aws_cloudtrail_trails.account_id, resource_id
