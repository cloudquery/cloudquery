with cachebeviors as (
	-- Handle all non defaults as well as when there is only a default route
	select distinct arn, account_id from (select arn,account_id, d as CacheBehavior from aws_cloudfront_distributions, jsonb_array_elements(distribution_config->'CacheBehaviors'->'Items') as d where distribution_config->'CacheBehaviors'->'Items' != 'null' 
	union 
	-- 	Handle default Cachebehaviors
	select arn,account_id, distribution_config->'DefaultCacheBehavior' as CacheBehavior from aws_cloudfront_distributions) as cachebeviors where CacheBehavior->>'ViewerProtocolPolicy' = 'allow-all'
)

insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CloudFront distributions should require encryption in transit' as title,
    account_id,
    arn as resource_id,
    'fail' as status
from cachebeviors