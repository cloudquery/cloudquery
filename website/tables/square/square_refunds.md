# Table: square_refunds

This table shows data for Square Refunds.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`string`|
|location_id|`string`|
|transaction_id|`string`|
|tender_id|`string`|
|created_at|`string`|
|reason|`string`|
|amount_money|`extension<json<JSONType>>`|
|status|`string`|
|processing_fee_money|`extension<json<JSONType>>`|
|additional_recipients|`extension<json<JSONType>>`|