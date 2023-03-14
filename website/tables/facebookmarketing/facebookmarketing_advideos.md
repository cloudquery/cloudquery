# Table: facebookmarketing_advideos

This table shows data for Facebook Marketing Ad Videos.

https://developers.facebook.com/docs/graph-api/reference/video#Reading

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|ad_breaks|IntArray|
|backdated_time|Timestamp|
|backdated_time_granularity|String|
|content_category|String|
|content_tags|StringArray|
|created_time|Timestamp|
|custom_labels|StringArray|
|description|String|
|embed_html|String|
|embeddable|Bool|
|event|JSON|
|icon|String|
|id (PK)|String|
|is_crosspost_video|Bool|
|is_crossposting_eligible|Bool|
|is_episode|Bool|
|is_instagram_eligible|Bool|
|length|Float|
|live_audience_count|Int|
|live_status|String|
|permalink_url|String|
|picture|String|
|place|JSON|
|post_views|Int|
|premiere_living_room_status|String|
|privacy|JSON|
|published|Bool|
|scheduled_publish_time|Timestamp|
|source|String|
|spherical|Bool|
|status|JSON|
|title|String|
|universal_video_id|String|
|updated_time|Timestamp|
|views|Int|