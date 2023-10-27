# Table: typeform_forms

This table shows data for Typeform Forms.

The primary key for this table is **id**.

## Relations

The following tables depend on typeform_forms:
  - [typeform_form_responses](typeform_form_responses)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`utf8`|
|created_at|`timestamp[s]`|
|last_updated_at|`timestamp[s]`|
|self|`json`|
|type|`utf8`|
|settings|`json`|
|theme|`json`|
|title|`utf8`|
|_links|`json`|