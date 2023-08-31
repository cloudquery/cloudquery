# Table: facebookmarketing_adcloudplayables

This table shows data for Facebook Marketing Ad Cloud Playables.

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|owner|`json`|
|playable_ad_file_size|`int64`|
|playable_ad_orientation|`utf8`|
|playable_ad_package_name|`utf8`|
|playable_ad_reject_reason|`utf8`|
|playable_ad_status|`utf8`|
|playable_ad_upload_time|`timestamp[us, tz=UTC]`|