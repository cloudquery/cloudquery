WITH basic_table as (
select 
	arn,
	region,
	jsonb_array_elements(replication_rules) -> 'Destination' ->> 'Bucket' as dest,
	jsonb_array_elements(replication_rules) ->> 'Status' as stat
	from
	aws_s3_buckets
),
sec as (
	select bt.arn,
		bt.region,
		bt.dest,
		asb.region as other,
		bt.stat
		from basic_table as bt
		left join aws_s3_buckets as asb on dest = asb.arn
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
            FROM sec
            WHERE 
			sec.arn = aws_s3_buckets.arn
			and
			stat = 'Enabled' and region != other
        )
        THEN 'pass'
        ELSE 'fail'
    END AS status
FROM
    aws_s3_buckets;