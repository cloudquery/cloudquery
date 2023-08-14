# Table: typeform_form_responses

This table shows data for Typeform Form Responses.

The composite primary key for this table is (**form_id**, **response_id**).

## Relations

This table depends on [typeform_forms](typeform_forms).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|form_id (PK)|`string`|
|response_id (PK)|`string`|
|landing_id|`string`|
|landed_at|`timestamp[s]`|
|submitted_at|`timestamp[s]`|
|token|`string`|
|metadata|`extension<json<JSONType>>`|
|answers|`extension<json<JSONType>>`|
|hidden|`extension<json<JSONType>>`|
|calculated|`extension<json<JSONType>>`|
|variables|`extension<json<JSONType>>`|