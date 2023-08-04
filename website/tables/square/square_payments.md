# Table: square_payments

This table shows data for Square Payments.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`string`|
|created_at|`string`|
|updated_at|`string`|
|amount_money|`extension<json<JSONType>>`|
|tip_money|`extension<json<JSONType>>`|
|total_money|`extension<json<JSONType>>`|
|app_fee_money|`extension<json<JSONType>>`|
|approved_money|`extension<json<JSONType>>`|
|processing_fee|`extension<json<JSONType>>`|
|refunded_money|`extension<json<JSONType>>`|
|status|`string`|
|delay_duration|`string`|
|delay_action|`string`|
|delayed_until|`string`|
|source_type|`string`|
|card_details|`extension<json<JSONType>>`|
|cash_details|`extension<json<JSONType>>`|
|bank_account_details|`extension<json<JSONType>>`|
|external_details|`extension<json<JSONType>>`|
|wallet_details|`extension<json<JSONType>>`|
|buy_now_pay_later_details|`extension<json<JSONType>>`|
|location_id|`string`|
|order_id|`string`|
|reference_id|`string`|
|customer_id|`string`|
|employee_id|`string`|
|team_member_id|`string`|
|refund_ids|`extension<json<JSONType>>`|
|risk_evaluation|`extension<json<JSONType>>`|
|buyer_email_address|`string`|
|billing_address|`extension<json<JSONType>>`|
|shipping_address|`extension<json<JSONType>>`|
|note|`string`|
|statement_description_identifier|`string`|
|capabilities|`extension<json<JSONType>>`|
|receipt_number|`string`|
|receipt_url|`string`|
|device_details|`extension<json<JSONType>>`|
|application_details|`extension<json<JSONType>>`|
|version_token|`string`|