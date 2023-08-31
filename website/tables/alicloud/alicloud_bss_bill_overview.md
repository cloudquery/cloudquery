# Table: alicloud_bss_bill_overview

This table shows data for Alibaba Cloud BSS Bill Overview.

https://help.aliyun.com/document_detail/100400.html

The composite primary key for this table is (**billing_cycle**, **account_id**, **product_code**, **bill_account_id**, **product_type**, **pip_code**, **subscription_type**, **commodity_code**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|billing_cycle (PK)|`utf8`|
|account_id (PK)|`utf8`|
|account_name|`utf8`|
|deducted_by_coupons|`float64`|
|round_down_discount|`utf8`|
|product_name|`utf8`|
|product_detail|`utf8`|
|product_code (PK)|`utf8`|
|bill_account_id (PK)|`utf8`|
|product_type (PK)|`utf8`|
|deducted_by_cash_coupons|`float64`|
|outstanding_amount|`float64`|
|biz_type|`utf8`|
|payment_amount|`float64`|
|pip_code (PK)|`utf8`|
|deducted_by_prepaid_card|`float64`|
|invoice_discount|`float64`|
|item|`utf8`|
|subscription_type (PK)|`utf8`|
|pretax_gross_amount|`float64`|
|pretax_amount|`float64`|
|owner_id|`utf8`|
|currency|`utf8`|
|commodity_code (PK)|`utf8`|
|bill_account_name|`utf8`|
|adjust_amount|`float64`|
|cash_amount|`float64`|