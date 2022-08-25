insert into aws_policy_results
WITH regions_with_enabled_accessanalyzer AS (
    SELECT ar.region AS analyzed_region
    FROM aws_regions ar
             LEFT JOIN aws_access_analyzer_analyzers aaaa ON
            ar.region = aaaa.region
    WHERE aaaa.status = 'ACTIVE' )
select
    :'execution_time' as execution_time,
        :'framework' as framework,
        :'check_id' as check_id,
        'certificate has less than 30 days to be renewed' as title,
    account_id,
    region AS resource_id,
    case when
             aregion.analyzed_region IS NULL
             AND ar.enabled = TRUE
         then 'fail'
         else 'pass'
        end as status
FROM aws_regions ar LEFT JOIN regions_with_enabled_accessanalyzer aregion ON
        ar.region = aregion.analyzed_region
