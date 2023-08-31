# Table: stripe_payment_methods

This table shows data for Stripe Payment Methods.

https://stripe.com/docs/api/payment_methods

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|acss_debit|`json`|
|affirm|`json`|
|afterpay_clearpay|`json`|
|alipay|`json`|
|au_becs_debit|`json`|
|bacs_debit|`json`|
|bancontact|`json`|
|billing_details|`json`|
|blik|`json`|
|boleto|`json`|
|card|`json`|
|card_present|`json`|
|cashapp|`json`|
|created|`timestamp[us, tz=UTC]`|
|customer|`json`|
|customer_balance|`json`|
|eps|`json`|
|fpx|`json`|
|giropay|`json`|
|grabpay|`json`|
|ideal|`json`|
|interac_present|`json`|
|klarna|`json`|
|konbini|`json`|
|link|`json`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|oxxo|`json`|
|p24|`json`|
|paynow|`json`|
|pix|`json`|
|promptpay|`json`|
|radar_options|`json`|
|sepa_debit|`json`|
|sofort|`json`|
|type|`utf8`|
|us_bank_account|`json`|
|wechat_pay|`json`|