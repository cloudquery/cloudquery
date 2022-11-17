# Table: gitlab_users_users



The primary key for this table is **uid**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|uid (PK)|String|
|id|Int|
|username|String|
|email|String|
|name|String|
|state|String|
|web_url|String|
|created_at|Timestamp|
|bio|String|
|bot|Bool|
|location|String|
|public_email|String|
|skype|String|
|linkedin|String|
|twitter|String|
|website_url|String|
|organization|String|
|job_title|String|
|extern_uid|String|
|provider|String|
|theme_id|Int|
|last_activity_on|JSON|
|color_scheme_id|Int|
|is_admin|Bool|
|avatar_url|String|
|can_create_group|Bool|
|can_create_project|Bool|
|projects_limit|Int|
|current_sign_in_at|Timestamp|
|current_sign_in_ip|IntArray|
|last_sign_in_at|Timestamp|
|last_sign_in_ip|IntArray|
|confirmed_at|Timestamp|
|two_factor_enabled|Bool|
|note|String|
|identities|JSON|
|external|Bool|
|private_profile|Bool|
|shared_runners_minutes_limit|Int|
|extra_shared_runners_minutes_limit|Int|
|using_license_seat|Bool|
|custom_attributes|JSON|
|namespace_id|Int|