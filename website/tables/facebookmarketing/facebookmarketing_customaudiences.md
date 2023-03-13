# Table: facebookmarketing_customaudiences

https://developers.facebook.com/docs/marketing-api/reference/custom-audience#Reading

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|approximate_count_lower_bound|Int|
|approximate_count_upper_bound|Int|
|customer_file_source|String|
|data_source|JSON|
|data_source_types|String|
|datafile_custom_audience_uploading_status|String|
|delete_time|Int|
|delivery_status|JSON|
|description|String|
|excluded_custom_audiences|JSON|
|household_audience|Int|
|id (PK)|String|
|included_custom_audiences|JSON|
|is_household|Bool|
|is_snapshot|Bool|
|is_value_based|Bool|
|lookalike_audience_ids|StringArray|
|lookalike_spec|JSON|
|name|String|
|operation_status|JSON|
|opt_out_link|String|
|owner_business|JSON|
|page_deletion_marked_delete_time|Int|
|permission_for_actions|JSON|
|pixel_id|String|
|regulated_audience_spec|JSON|
|retention_days|Int|
|rev_share_policy_id|Int|
|rule|String|
|rule_aggregation|String|
|rule_v2|String|
|seed_audience|Int|
|sharing_status|JSON|
|subtype|String|
|time_content_updated|Int|
|time_created|Int|
|time_updated|Int|