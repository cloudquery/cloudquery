insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Ensure CloudTrail is enabled in all regions' as title,
    aws_cloudtrail_trails.account_id,
    arn as resource_id,
    case
        when aws_cloudtrail_trails.is_multi_region_trail = FALSE then 'fail'
        when exists(select *
                    from jsonb_array_elements(aws_cloudtrail_trail_event_selectors.event_selectors) as es
                    where (es ->>'ReadWriteType')::text != 'All' or (es->>'IncludeManagementEvents')::boolean = FALSE)
            then 'fail'
        when exists(select *
                    from jsonb_array_elements(aws_cloudtrail_trail_event_selectors.advanced_event_selectors) as aes
                    where exists(select *
                                 from jsonb_array_elements((aes ->>'FieldSelectors')::jsonb) as aes_fs
                                 where (aes_fs ->>'Field')::text = 'readOnly'))
            then 'fail'
        else 'pass'
    end as status
from aws_cloudtrail_trails
inner join
    aws_cloudtrail_trail_event_selectors on
        aws_cloudtrail_trails.arn = aws_cloudtrail_trail_event_selectors.trail_arn
        and aws_cloudtrail_trails.region = aws_cloudtrail_trail_event_selectors.region
        and aws_cloudtrail_trails.account_id = aws_cloudtrail_trail_event_selectors.account_id
