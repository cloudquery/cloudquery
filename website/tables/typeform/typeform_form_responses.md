# Table: typeform_form_responses

This table shows data for Typeform Form Responses.

The composite primary key for this table is (**form_id**, **response_id**).

## Relations

This table depends on [typeform_forms](typeform_forms).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|form_id (PK)|`utf8`|
|response_id (PK)|`utf8`|
|landing_id|`utf8`|
|landed_at|`timestamp[s]`|
|submitted_at|`timestamp[s]`|
|token|`utf8`|
|metadata|`json`|
|answers|`json`|
|hidden|`json`|
|calculated|`json`|
|variables|`json`|