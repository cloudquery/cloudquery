INSERT INTO aws_policy_results
WITH global_trail_status AS (
    SELECT
        trails.account_id,
        trails.arn,
        bool_or(
            is_multi_region_trail
            AND (status->>'IsLogging')::bool
            AND include_management_events
            AND read_write_type = 'All'
        ) AS is_global_trail
    FROM
        aws_cloudtrail_trails trails
        JOIN aws_cloudtrail_trail_event_selectors actes
             ON trails._cq_id = actes._cq_parent_id
    GROUP BY trails.account_id, trails.arn
)
SELECT
    :'execution_time'                             AS execution_time,
    :'framework'                                  AS framework,
    :'check_id'                                   AS check_id,
    'Ensure CloudTrail is enabled in all regions' AS title,
    account_id,
    arn                                           AS resource_id,
    CASE
        WHEN is_global_trail
        THEN 'pass'
        ELSE 'fail'
    END                                           AS status
FROM global_trail_status
