INSERT INTO aws_policy_results
WITH failed_trails AS (
    SELECT act.account_id,
           act.arn
    FROM aws_cloudtrail_trails act
             JOIN aws_s3_buckets as3b on act.s3_bucket_name = as3b.name
             JOIN aws_s3_bucket_grants asbg on as3b._cq_id = asbg._cq_parent_id,
         jsonb_array_elements(as3b.policy -> 'Statement') as statement
    WHERE grantee ->> 'URI' = 'http://acs.amazonaws.com/groups/global/AllUsers'
       OR grantee ->> 'URI' = 'http://acs.amazonaws.com/groups/global/AuthenticatedUsers'
       OR (statement ->> 'Effect' = 'Allow' AND
           (statement ->> 'Principal' = '*' OR statement -> 'Principal' ->> 'AWS' = '*'))
    GROUP BY act.account_id, act.arn
)
SELECT
    :'execution_time' AS execution_time,
    :'framework' AS framework,
    :'check_id' AS check_id,
    'Ensure the S3 bucket used to store CloudTrail logs is not publicly accessible' AS title,
    aws_cloudtrail_trails.account_id AS account_id,
    aws_cloudtrail_trails.arn AS resource_id,
    CASE
        WHEN bool_and(failed_trails.arn IS NULL)
        THEN 'pass'
        ELSE 'fail'
    END AS status
FROM aws_cloudtrail_trails LEFT JOIN failed_trails ON aws_cloudtrail_trails.arn = failed_trails.arn
GROUP BY aws_cloudtrail_trails.account_id, aws_cloudtrail_trails.arn
