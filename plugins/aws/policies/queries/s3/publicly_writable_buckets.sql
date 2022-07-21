insert into aws_policy_results
with policy_allow_public as (
    select
        arn,
        bucket_cq_id,
        count(*) as statement_count
    from
        (
            select
                aws_s3_buckets.arn,
                aws_s3_buckets.cq_id as bucket_cq_id,
                statements -> 'Principal' as principals
            from
                aws_s3_buckets,
                jsonb_array_elements(
                    case jsonb_typeof(policy -> 'Statement')
                        when
                            'string' then jsonb_build_array(
                                policy ->> 'Statement'
                            )
                        when 'array' then policy -> 'Statement'
                    end
                ) as statements
            where
                statements -> 'Effect' = '"Allow"'
        ) as foo
    where
        principals = '"*"'
        or (
            principals::JSONB ? 'AWS'
            and (
                principals -> 'AWS' = '"*"'
                or principals -> 'AWS' @> '"*"'
            )
        )
    group by
        arn, bucket_cq_id
)

select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'S3 buckets should prohibit public write access' as title,
    aws_s3_buckets.account_id,
    aws_s3_buckets.arn as resource_id,
    'fail' as status -- TODO FIXME
from
    -- Find and join all bucket ACLS that give a public write access
    aws_s3_buckets
left join
    aws_s3_bucket_grants on
        aws_s3_buckets.cq_id = aws_s3_bucket_grants.bucket_cq_id
-- Find all statements that could give public allow access 
-- Statements that give public access have 1) Effect == Allow 2) One of the following principal:
--       Principal = {"AWS": "*"}
--       Principal = {"AWS": ["arn:aws:iam::12345678910:root", "*"]}
--       Principal = "*"
left join policy_allow_public on
        aws_s3_buckets.cq_id = policy_allow_public.bucket_cq_id
where
    (
        aws_s3_buckets.block_public_acls != TRUE
        and (
            uri = 'http://acs.amazonaws.com/groups/global/AllUsers'
            and permission in ('WRITE_ACP', 'FULL_CONTROL')
        )
    )
    or (
        aws_s3_buckets.block_public_policy != TRUE
        and policy_allow_public.statement_count > 0
    )
