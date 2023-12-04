insert into aws_policy_results

with any_category as (
    select distinct TRUE as any_category
    from aws_rds_event_subscriptions
    where
        (source_type is null or source_type = 'db-cluster')
        and event_categories_list is null
),

any_source_id as (
    select COALESCE(ARRAY_AGG(category), '{}'::TEXT[]) as any_source_categories
    from
        aws_rds_event_subscriptions,
        UNNEST(event_categories_list) as category
    where
        source_type = 'db-cluster'
        and event_categories_list is not null
),

specific_categories as (
    select
        source_id,
        ARRAY_AGG(category) as specific_cats
    from
        aws_rds_event_subscriptions,
        UNNEST(source_ids_list) as source_id,
        UNNEST(event_categories_list) as category
    where source_type = 'db-cluster'
    group by source_id
)

select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'An RDS event notifications subscription should be configured for critical cluster events' as title,
    account_id,
    arn as resource_id,
    case when
    any_category is not TRUE
  and not any_source_categories @> '{"failure","maintenance"}'
  and (
        specific_cats is null or not specific_cats @> '{"failure","maintenance"}'
    )
    then 'fail' else 'pass' end as status
from
    aws_rds_clusters
left outer join any_category on TRUE
inner join any_source_id on TRUE
left outer join
    specific_categories on
        db_cluster_identifier = specific_categories.source_id
