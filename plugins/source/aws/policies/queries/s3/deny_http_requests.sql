INSERT INTO aws_policy_results
WITH pass_buckets AS (
    SELECT
        arn
    FROM
        aws_s3_buckets s3,
        jsonb_array_elements(s3.policy->'Statement') statement,
        jsonb_array_elements_text(
            CASE jsonb_typeof(statement->'Resource')
                WHEN 'string' THEN jsonb_build_array(statement->>'Resource')
                WHEN 'array' THEN statement->'Resource'
            END
        ) resource
    WHERE
        statement->>'Effect' = 'Deny'
        AND statement->>'Action' = 's3:*'
        AND statement->'Condition'->'Bool'->>'aws:SecureTransport' IS NOT DISTINCT FROM 'false'
        AND (
            statement->>'Principal' = '*'
         OR statement->'Principal'->>'AWS' IS NOT DISTINCT FROM '*'
        )
        AND resource ~ CONCAT('^arn:aws:s3:::', name, '\/\*$')
    GROUP BY arn
)
SELECT
    :'execution_time'::timestamp                AS execution_time,
    :'framework'                                AS framework,
    :'check_id'                                 AS check_id,
    'S3 buckets should deny non-HTTPS requests' AS title,
    account_id,
    aws_s3_buckets.arn                          AS resource_id,
    CASE
        WHEN pass_buckets.arn IS NULL THEN 'fail'
        ELSE 'pass'
    END                                         AS status
FROM aws_s3_buckets
    LEFT JOIN pass_buckets ON aws_s3_buckets.arn = pass_buckets.arn
