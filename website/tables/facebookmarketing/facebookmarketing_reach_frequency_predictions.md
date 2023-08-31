# Table: facebookmarketing_reach_frequency_predictions

This table shows data for Facebook Marketing Reach Frequency Predictions.

https://developers.facebook.com/docs/marketing-api/reference/reach-frequency-prediction/#Reading

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`int64`|
|activity_status|`json`|
|ad_formats|`json`|
|auction_entry_option_index|`int64`|
|audience_size_lower_bound|`int64`|
|audience_size_upper_bound|`int64`|
|business_id|`int64`|
|buying_type|`utf8`|
|campaign_group_id|`int64`|
|campaign_id|`utf8`|
|campaign_time_start|`timestamp[us, tz=UTC]`|
|campaign_time_stop|`timestamp[us, tz=UTC]`|
|currency|`utf8`|
|curve_budget_reach|`json`|
|curve_reach|`list<item: int64, nullable>`|
|daily_grp_curve|`json`|
|daily_impression_curve|`json`|
|day_parting_schedule|`json`|
|destination_id|`utf8`|
|end_time|`timestamp[us, tz=UTC]`|
|expiration_time|`timestamp[us, tz=UTC]`|
|external_budget|`int64`|
|external_impression|`int64`|
|external_maximum_budget|`int64`|
|external_maximum_impression|`utf8`|
|external_maximum_reach|`int64`|
|external_minimum_budget|`int64`|
|external_minimum_impression|`int64`|
|external_minimum_reach|`int64`|
|external_reach|`int64`|
|feed_ratio_0000|`int64`|
|frequency_cap|`int64`|
|grp_audience_size|`float64`|
|grp_avg_probability_map|`utf8`|
|grp_country_audience_size|`float64`|
|grp_curve|`json`|
|grp_dmas_audience_size|`float64`|
|grp_filtering_threshold_00|`int64`|
|grp_points|`float64`|
|grp_ratio|`float64`|
|grp_reach_ratio|`float64`|
|grp_status|`utf8`|
|holdout_percentage|`int64`|
|id (PK)|`utf8`|
|impression_curve|`list<item: int64, nullable>`|
|instagram_destination_id|`utf8`|
|instream_packages|`list<item: utf8, nullable>`|
|interval_frequency_cap|`int64`|
|interval_frequency_cap_reset_period|`int64`|
|is_bonus_media|`int64`|
|is_conversion_goal|`int64`|
|is_higher_average_frequency|`bool`|
|is_io|`bool`|
|is_reserved_buying|`int64`|
|is_trp|`bool`|
|name|`utf8`|
|objective|`int64`|
|objective_name|`utf8`|
|odax_objective|`int64`|
|odax_objective_name|`utf8`|
|optimization_goal|`int64`|
|optimization_goal_name|`utf8`|
|placement_breakdown|`json`|
|plan_name|`utf8`|
|plan_type|`utf8`|
|prediction_mode|`int64`|
|prediction_progress|`int64`|
|reference_id|`utf8`|
|reservation_status|`int64`|
|start_time|`timestamp[us, tz=UTC]`|
|status|`int64`|
|story_event_type|`int64`|
|target_cpm|`int64`|
|target_spec|`json`|
|time_created|`timestamp[us, tz=UTC]`|
|time_updated|`timestamp[us, tz=UTC]`|
|timezone_id|`int64`|
|timezone_name|`utf8`|
|topline_id|`int64`|
|video_view_length_constraint|`int64`|
|viewtag|`utf8`|