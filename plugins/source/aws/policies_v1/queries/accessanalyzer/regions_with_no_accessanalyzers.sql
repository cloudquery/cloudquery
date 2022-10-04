insert into aws_policy_results
WITH regions_with_enabled_accessanalyzer
         AS (SELECT ar.region AS analyzed_region
             FROM aws_regions ar
                      LEFT JOIN aws_accessanalyzer_analyzers aaaa ON
                 ar.region = aaaa.region
             WHERE aaaa.status = 'ACTIVE')
SELECT :'execution_time'                                                        AS execution_time,
       :'framework'                                                             AS framework,
       :'check_id'                                                              AS check_id,
       'Ensure that IAM Access analyzer is enabled for all regions (Automated)' AS title,
       account_id,
       region                                                                   AS resource_id,
       CASE
           WHEN
                   aregion.analyzed_region IS NULL
                   AND ar.enabled = TRUE
               THEN 'fail'
           ELSE 'pass'
           END                                                                  AS status
FROM aws_regions ar
         LEFT JOIN regions_with_enabled_accessanalyzer aregion ON
    ar.region = aregion.analyzed_region;