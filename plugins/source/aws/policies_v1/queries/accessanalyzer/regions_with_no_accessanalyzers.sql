insert into aws_policy_results
SELECT 
    :'execution_time'                                                        AS execution_time,
    :'framework'                                                             AS framework,
    :'check_id'                                                              AS check_id,
    'Ensure that IAM Access analyzer is enabled for all regions (Automated)' AS title,
    ar.account_id,
    ar.region                                                                AS resource_id,
    CASE
        WHEN
            ar.enabled
            AND aregion.region IS NULL
            AND aregion.status IS DISTINCT FROM 'ACTIVE'
        THEN 'fail'
        ELSE 'pass'
    END                                                                      AS status
FROM
    aws_regions ar
    LEFT JOIN aws_accessanalyzer_analyzers aregion ON
        ar.region = aregion.region;
