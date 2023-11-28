# Table: square_refunds

This table shows data for Square Refunds.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`utf8`|
|location_id|`utf8`|
|transaction_id|`utf8`|
|tender_id|`utf8`|
|created_at|`utf8`|
|reason|`utf8`|
|amount_money|`json`|
|status|`utf8`|
|processing_fee_money|`json`|
|additional_recipients|`json`|