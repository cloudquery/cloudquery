insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'S3 buckets should deny non-HTTPS requests' as title,
    account_id,
    arn as resource_id,
    'fail' as status
from
    aws_s3_buckets
where
    arn not in (
        -- Find all buckets that have a bucket policy that denies non-SSL requests
        select arn
        from (select aws_s3_buckets.arn,
                     statements,
                     statements -> 'Principal' as principals
              from aws_s3_buckets,
                   jsonb_array_elements(
                           case jsonb_typeof(policy -> 'Statement')
                               when
                                   'string' then jsonb_build_array(
                                       policy ->> 'Statement'
                                   )
                               when 'array' then policy -> 'Statement'
                               end
                       ) as statements
              where statements -> 'Effect' = '"Deny"') as foo,
             jsonb_array_elements_text(
                     statements -> 'Condition' -> 'Bool' -> 'aws:securetransport'
                 ) as ssl
        where principals = '"*"'
           or (
                          principals::JSONB ? 'AWS'
                      and (
                                          principals -> 'AWS' = '"*"'
                                  or principals -> 'AWS' @> '"*"'
                              )
                  )
            and ssl::BOOL = FALSE
        )
