# Table: slack_files

https://api.slack.com/methods/files.list

The composite primary key for this table is (**team_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|team_id (PK)|String|
|id (PK)|String|
|created|Timestamp|
|timestamp|Timestamp|
|name|String|
|title|String|
|mimetype|String|
|image_exif_rotation|Int|
|filetype|String|
|pretty_type|String|
|user|String|
|mode|String|
|editable|Bool|
|is_external|Bool|
|external_type|String|
|size|Int|
|url|String|
|url_download|String|
|url_private|String|
|url_private_download|String|
|original_h|Int|
|original_w|Int|
|thumb_64|String|
|thumb_80|String|
|thumb_160|String|
|thumb_360|String|
|thumb_360_gif|String|
|thumb_360_w|Int|
|thumb_360_h|Int|
|thumb_480|String|
|thumb_480_w|Int|
|thumb_480_h|Int|
|thumb_720|String|
|thumb_720_w|Int|
|thumb_720_h|Int|
|thumb_960|String|
|thumb_960_w|Int|
|thumb_960_h|Int|
|thumb_1024|String|
|thumb_1024_w|Int|
|thumb_1024_h|Int|
|permalink|String|
|permalink_public|String|
|edit_link|String|
|preview|String|
|preview_highlight|String|
|lines|Int|
|lines_more|Int|
|is_public|Bool|
|public_url_shared|Bool|
|channels|StringArray|
|groups|StringArray|
|ims|StringArray|
|initial_comment|JSON|
|comments_count|Int|
|num_stars|Int|
|is_starred|Bool|
|shares|JSON|