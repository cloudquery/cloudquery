# Table: stripe_payment_methods

https://stripe.com/docs/api/payment_methods

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|acss_debit|JSON|
|affirm|JSON|
|afterpay_clearpay|JSON|
|alipay|JSON|
|au_becs_debit|JSON|
|bacs_debit|JSON|
|bancontact|JSON|
|billing_details|JSON|
|blik|JSON|
|boleto|JSON|
|card|JSON|
|card_present|JSON|
|created|Timestamp|
|customer|JSON|
|customer_balance|JSON|
|eps|JSON|
|fpx|JSON|
|giropay|JSON|
|grabpay|JSON|
|ideal|JSON|
|interac_present|JSON|
|klarna|JSON|
|konbini|JSON|
|link|JSON|
|livemode|Bool|
|metadata|JSON|
|object|String|
|oxxo|JSON|
|p24|JSON|
|paynow|JSON|
|pix|JSON|
|promptpay|JSON|
|radar_options|JSON|
|sepa_debit|JSON|
|sofort|JSON|
|type|String|
|us_bank_account|JSON|
|wechat_pay|JSON|