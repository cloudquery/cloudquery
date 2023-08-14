# Table: square_bookings

This table shows data for Square Bookings.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`string`|
|version|`int64`|
|status|`string`|
|created_at|`string`|
|updated_at|`string`|
|start_at|`string`|
|location_id|`string`|
|customer_id|`string`|
|customer_note|`string`|
|seller_note|`string`|
|appointment_segments|`extension<json<JSONType>>`|
|transition_time_minutes|`int64`|
|all_day|`bool`|
|location_type|`string`|
|creator_details|`extension<json<JSONType>>`|
|source|`string`|