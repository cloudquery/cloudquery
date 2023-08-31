# Table: facebookmarketing_advideos

This table shows data for Facebook Marketing Ad Videos.

https://developers.facebook.com/docs/graph-api/reference/video#Reading

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|ad_breaks|`list<item: int64, nullable>`|
|backdated_time|`timestamp[us, tz=UTC]`|
|backdated_time_granularity|`utf8`|
|content_category|`utf8`|
|content_tags|`list<item: utf8, nullable>`|
|created_time|`timestamp[us, tz=UTC]`|
|custom_labels|`list<item: utf8, nullable>`|
|description|`utf8`|
|embed_html|`utf8`|
|embeddable|`bool`|
|event|`json`|
|icon|`utf8`|
|id (PK)|`utf8`|
|is_crosspost_video|`bool`|
|is_crossposting_eligible|`bool`|
|is_episode|`bool`|
|is_instagram_eligible|`bool`|
|length|`float64`|
|live_audience_count|`int64`|
|live_status|`utf8`|
|permalink_url|`utf8`|
|picture|`utf8`|
|place|`json`|
|post_views|`int64`|
|premiere_living_room_status|`utf8`|
|privacy|`json`|
|published|`bool`|
|scheduled_publish_time|`timestamp[us, tz=UTC]`|
|source|`utf8`|
|spherical|`bool`|
|status|`json`|
|title|`utf8`|
|universal_video_id|`utf8`|
|updated_time|`timestamp[us, tz=UTC]`|
|views|`int64`|