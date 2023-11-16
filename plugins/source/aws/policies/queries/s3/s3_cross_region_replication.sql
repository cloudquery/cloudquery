WITH s3_replication_info AS (
    SELECT
        arn,
        region,
        jsonb_array_elements(replication_rules) -> 'Destination' ->> 'Bucket' as destination_bucket,
        jsonb_array_elements(replication_rules) ->> 'Status' as replication_status
    FROM aws_s3_buckets
),
cross_region_replication AS (
    SELECT
        sri.arn,
        sri.region,
        sri.destination_bucket,
        asb.region as destination_region,
        sri.replication_status
    FROM s3_replication_info sri
    LEFT JOIN aws_s3_buckets asb ON sri.destination_bucket = asb.arn
)
insert into aws_policy_results
SELECT
	:'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'S3 buckets should have cross-Region replication enabled' as title,
    aws_s3_buckets.account_id,
    aws_s3_buckets.arn AS resource_id,
    CASE
        WHEN EXISTS (
            SELECT 1
            FROM cross_region_replication
            WHERE 
                cross_region_replication.arn = aws_s3_buckets.arn
                AND replication_status = 'Enabled'
                AND region != destination_region
        )
        THEN 'pass'
        ELSE 'fail'
    END AS status
FROM
    aws_s3_buckets;