# Table: heroku_teams

This table shows data for Heroku Teams.

https://devcenter.heroku.com/articles/platform-api-reference#team

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|created_at|timestamp[us, tz=UTC]|
|credit_card_collections|bool|
|default|bool|
|enterprise_account|json|
|identity_provider|json|
|membership_limit|float64|
|name|utf8|
|provisioned_licenses|bool|
|role|utf8|
|type|utf8|
|updated_at|timestamp[us, tz=UTC]|