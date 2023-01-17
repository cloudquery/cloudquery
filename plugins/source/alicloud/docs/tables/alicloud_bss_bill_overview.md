# Table: alicloud_bss_bill_overview

https://help.aliyun.com/document_detail/100400.html

The composite primary key for this table is (**billing_cycle**, **account_id**, **product_code**, **bill_account_id**, **product_type**, **pip_code**, **subscription_type**, **commodity_code**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|billing_cycle (PK)|String|
|account_id (PK)|String|
|account_name|String|
|deducted_by_coupons|Float|
|round_down_discount|String|
|product_name|String|
|product_detail|String|
|product_code (PK)|String|
|bill_account_id (PK)|String|
|product_type (PK)|String|
|deducted_by_cash_coupons|Float|
|outstanding_amount|Float|
|biz_type|String|
|payment_amount|Float|
|pip_code (PK)|String|
|deducted_by_prepaid_card|Float|
|invoice_discount|Float|
|item|String|
|subscription_type (PK)|String|
|pretax_gross_amount|Float|
|pretax_amount|Float|
|owner_id|String|
|currency|String|
|commodity_code (PK)|String|
|bill_account_name|String|
|adjust_amount|Float|
|cash_amount|Float|