# Table: gitlab_users

This table shows data for Gitlab Users.

The composite primary key for this table is (**base_url**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|base_url (PK)|`utf8`|
|id (PK)|`int64`|
|username|`utf8`|
|email|`utf8`|
|name|`utf8`|
|state|`utf8`|
|web_url|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|bio|`utf8`|
|bot|`bool`|
|location|`utf8`|
|public_email|`utf8`|
|skype|`utf8`|
|linkedin|`utf8`|
|twitter|`utf8`|
|website_url|`utf8`|
|organization|`utf8`|
|job_title|`utf8`|
|extern_uid|`utf8`|
|provider|`utf8`|
|theme_id|`int64`|
|last_activity_on|`timestamp[us, tz=UTC]`|
|color_scheme_id|`int64`|
|is_admin|`bool`|
|avatar_url|`utf8`|
|can_create_group|`bool`|
|can_create_project|`bool`|
|projects_limit|`int64`|
|current_sign_in_at|`timestamp[us, tz=UTC]`|
|current_sign_in_ip|`inet`|
|last_sign_in_at|`timestamp[us, tz=UTC]`|
|last_sign_in_ip|`inet`|
|confirmed_at|`timestamp[us, tz=UTC]`|
|two_factor_enabled|`bool`|
|note|`utf8`|
|identities|`json`|
|external|`bool`|
|private_profile|`bool`|
|shared_runners_minutes_limit|`int64`|
|extra_shared_runners_minutes_limit|`int64`|
|using_license_seat|`bool`|
|custom_attributes|`json`|
|namespace_id|`int64`|