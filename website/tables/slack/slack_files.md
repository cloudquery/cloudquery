# Table: slack_files

This table shows data for Slack Files.

https://api.slack.com/methods/files.list

The composite primary key for this table is (**team_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|team_id (PK)|`utf8`|
|id (PK)|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|timestamp|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|title|`utf8`|
|mimetype|`utf8`|
|image_exif_rotation|`int64`|
|filetype|`utf8`|
|pretty_type|`utf8`|
|user|`utf8`|
|mode|`utf8`|
|editable|`bool`|
|is_external|`bool`|
|external_type|`utf8`|
|size|`int64`|
|url|`utf8`|
|url_download|`utf8`|
|url_private|`utf8`|
|url_private_download|`utf8`|
|original_h|`int64`|
|original_w|`int64`|
|thumb_64|`utf8`|
|thumb_80|`utf8`|
|thumb_160|`utf8`|
|thumb_360|`utf8`|
|thumb_360_gif|`utf8`|
|thumb_360_w|`int64`|
|thumb_360_h|`int64`|
|thumb_480|`utf8`|
|thumb_480_w|`int64`|
|thumb_480_h|`int64`|
|thumb_720|`utf8`|
|thumb_720_w|`int64`|
|thumb_720_h|`int64`|
|thumb_960|`utf8`|
|thumb_960_w|`int64`|
|thumb_960_h|`int64`|
|thumb_1024|`utf8`|
|thumb_1024_w|`int64`|
|thumb_1024_h|`int64`|
|permalink|`utf8`|
|permalink_public|`utf8`|
|edit_link|`utf8`|
|preview|`utf8`|
|preview_highlight|`utf8`|
|lines|`int64`|
|lines_more|`int64`|
|is_public|`bool`|
|public_url_shared|`bool`|
|channels|`list<item: utf8, nullable>`|
|groups|`list<item: utf8, nullable>`|
|ims|`list<item: utf8, nullable>`|
|initial_comment|`json`|
|comments_count|`int64`|
|num_stars|`int64`|
|is_starred|`bool`|
|shares|`json`|