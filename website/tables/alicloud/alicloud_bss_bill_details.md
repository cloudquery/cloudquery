# Table: alicloud_bss_bill_details

This table shows data for Alibaba Cloud BSS Bill Details.

https://help.aliyun.com/document_detail/100392.html

The composite primary key for this table is (**billing_cycle**, **billing_date**, **account_id**, **subscription_type**, **commodity_code**, **product_code**, **product_type**, **pip_code**, **record_id**, **instance_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|billing_cycle (PK)|`utf8`|
|billing_date (PK)|`utf8`|
|account_id (PK)|`utf8`|
|account_name|`utf8`|
|product_name|`utf8`|
|sub_order_id|`utf8`|
|deducted_by_cash_coupons|`float64`|
|payment_time|`utf8`|
|payment_amount|`float64`|
|deducted_by_prepaid_card|`float64`|
|invoice_discount|`float64`|
|usage_end_time|`utf8`|
|item|`utf8`|
|subscription_type (PK)|`utf8`|
|pretax_gross_amount|`float64`|
|currency|`utf8`|
|commodity_code (PK)|`utf8`|
|usage_start_time|`utf8`|
|adjust_amount|`float64`|
|status|`utf8`|
|deducted_by_coupons|`float64`|
|round_down_discount|`utf8`|
|product_detail|`utf8`|
|product_code (PK)|`utf8`|
|product_type (PK)|`utf8`|
|outstanding_amount|`float64`|
|pip_code (PK)|`utf8`|
|pretax_amount|`float64`|
|owner_id|`utf8`|
|record_id (PK)|`utf8`|
|resource_group|`utf8`|
|instance_id (PK)|`utf8`|
|cash_amount|`float64`|