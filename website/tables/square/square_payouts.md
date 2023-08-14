# Table: square_payouts

This table shows data for Square Payouts.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`string`|
|status|`string`|
|location_id|`string`|
|created_at|`string`|
|updated_at|`string`|
|amount_money|`extension<json<JSONType>>`|
|destination|`extension<json<JSONType>>`|
|version|`int64`|
|type|`string`|
|payout_fee|`extension<json<JSONType>>`|
|arrival_date|`string`|
|end_to_end_id|`string`|