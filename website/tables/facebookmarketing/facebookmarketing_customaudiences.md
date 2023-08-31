# Table: facebookmarketing_customaudiences

This table shows data for Facebook Marketing Custom Audiences.

https://developers.facebook.com/docs/marketing-api/reference/custom-audience#Reading

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|approximate_count_lower_bound|`int64`|
|approximate_count_upper_bound|`int64`|
|customer_file_source|`utf8`|
|data_source|`json`|
|data_source_types|`utf8`|
|datafile_custom_audience_uploading_status|`utf8`|
|delete_time|`int64`|
|delivery_status|`json`|
|description|`utf8`|
|excluded_custom_audiences|`json`|
|household_audience|`int64`|
|id (PK)|`utf8`|
|included_custom_audiences|`json`|
|is_household|`bool`|
|is_snapshot|`bool`|
|is_value_based|`bool`|
|lookalike_audience_ids|`list<item: utf8, nullable>`|
|lookalike_spec|`json`|
|name|`utf8`|
|operation_status|`json`|
|opt_out_link|`utf8`|
|owner_business|`json`|
|page_deletion_marked_delete_time|`int64`|
|permission_for_actions|`json`|
|pixel_id|`utf8`|
|regulated_audience_spec|`json`|
|retention_days|`int64`|
|rev_share_policy_id|`int64`|
|rule|`utf8`|
|rule_aggregation|`utf8`|
|rule_v2|`utf8`|
|seed_audience|`int64`|
|sharing_status|`json`|
|subtype|`utf8`|
|time_content_updated|`int64`|
|time_created|`int64`|
|time_updated|`int64`|