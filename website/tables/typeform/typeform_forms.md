# Table: typeform_forms

This table shows data for Typeform Forms.

The primary key for this table is **id**.

## Relations

The following tables depend on typeform_forms:

  - [typeform_form_responses](typeform_form_responses)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`string`|
|created_at|`timestamp[s]`|
|last_updated_at|`timestamp[s]`|
|self|`extension<json<JSONType>>`|
|type|`string`|
|settings|`extension<json<JSONType>>`|
|theme|`extension<json<JSONType>>`|
|title|`string`|
|_links|`extension<json<JSONType>>`|