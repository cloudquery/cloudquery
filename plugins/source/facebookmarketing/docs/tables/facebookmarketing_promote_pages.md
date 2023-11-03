# Table: facebookmarketing_promote_pages

This table shows data for Facebook Marketing Promote Pages.

https://developers.facebook.com/docs/graph-api/reference/page#Reading

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|about|`utf8`|
|access_token|`utf8`|
|ad_campaign|`json`|
|affiliation|`utf8`|
|app_id|`utf8`|
|artists_we_like|`utf8`|
|attire|`utf8`|
|awards|`utf8`|
|band_interests|`utf8`|
|band_members|`utf8`|
|best_page|`json`|
|bio|`utf8`|
|birthday|`utf8`|
|booking_agent|`utf8`|
|built|`utf8`|
|can_checkin|`bool`|
|can_post|`bool`|
|category|`utf8`|
|category_list|`json`|
|checkins|`int64`|
|company_overview|`utf8`|
|connected_instagram_account|`json`|
|connected_page_backed_instagram_account|`json`|
|contact_address|`json`|
|copyright_whitelisted_ig_partners|`list<item: utf8, nullable>`|
|country_page_likes|`int64`|
|cover|`json`|
|culinary_team|`utf8`|
|current_location|`utf8`|
|delivery_and_pickup_option_info|`list<item: utf8, nullable>`|
|description|`utf8`|
|description_html|`utf8`|
|differently_open_offerings|`json`|
|directed_by|`utf8`|
|display_subtext|`utf8`|
|displayed_message_response_time|`utf8`|
|emails|`list<item: utf8, nullable>`|
|engagement|`json`|
|fan_count|`int64`|
|featured_video|`json`|
|features|`utf8`|
|followers_count|`int64`|
|food_styles|`list<item: utf8, nullable>`|
|founded|`utf8`|
|general_info|`utf8`|
|general_manager|`utf8`|
|genre|`utf8`|
|global_brand_page_name|`utf8`|
|global_brand_root_id|`utf8`|
|has_added_app|`bool`|
|has_transitioned_to_new_page_experience|`bool`|
|has_whatsapp_business_number|`bool`|
|has_whatsapp_number|`bool`|
|hometown|`utf8`|
|hours|`json`|
|id (PK)|`utf8`|
|impressum|`utf8`|
|influences|`utf8`|
|instagram_business_account|`json`|
|instant_articles_review_status|`utf8`|
|is_always_open|`bool`|
|is_chain|`bool`|
|is_community_page|`bool`|
|is_eligible_for_branded_content|`bool`|
|is_messenger_bot_get_started_enabled|`bool`|
|is_messenger_platform_bot|`bool`|
|is_owned|`bool`|
|is_permanently_closed|`bool`|
|is_published|`bool`|
|is_unclaimed|`bool`|
|is_verified|`bool`|
|is_webhooks_subscribed|`bool`|
|leadgen_tos_acceptance_time|`timestamp[us, tz=UTC]`|
|leadgen_tos_accepted|`bool`|
|leadgen_tos_accepting_user|`json`|
|link|`utf8`|
|location|`json`|
|members|`utf8`|
|merchant_id|`utf8`|
|merchant_review_status|`utf8`|
|messaging_feature_status|`json`|
|messenger_ads_default_icebreakers|`list<item: utf8, nullable>`|
|messenger_ads_default_page_welcome_message|`json`|
|messenger_ads_default_quick_replies|`list<item: utf8, nullable>`|
|messenger_ads_quick_replies_type|`utf8`|
|mini_shop_storefront|`json`|
|mission|`utf8`|
|mpg|`utf8`|
|name|`utf8`|
|name_with_location_descriptor|`utf8`|
|network|`utf8`|
|new_like_count|`int64`|
|offer_eligible|`bool`|
|overall_star_rating|`float64`|
|owner_business|`json`|
|page_token|`utf8`|
|parent_page|`json`|
|parking|`json`|
|payment_options|`json`|
|personal_info|`utf8`|
|personal_interests|`utf8`|
|pharma_safety_info|`utf8`|
|phone|`utf8`|
|pickup_options|`list<item: utf8, nullable>`|
|place_type|`utf8`|
|plot_outline|`utf8`|
|preferred_audience|`json`|
|press_contact|`utf8`|
|price_range|`utf8`|
|privacy_info_url|`utf8`|
|produced_by|`utf8`|
|products|`utf8`|
|promotion_eligible|`bool`|
|promotion_ineligible_reason|`utf8`|
|public_transit|`utf8`|
|rating_count|`int64`|
|recipient|`utf8`|
|record_label|`utf8`|
|release_date|`utf8`|
|restaurant_services|`json`|
|restaurant_specialties|`json`|
|schedule|`utf8`|
|screenplay_by|`utf8`|
|season|`utf8`|
|single_line_address|`utf8`|
|starring|`utf8`|
|start_info|`json`|
|store_code|`utf8`|
|store_location_descriptor|`utf8`|
|store_number|`int64`|
|studio|`utf8`|
|supports_donate_button_in_live_video|`bool`|
|supports_instant_articles|`bool`|
|talking_about_count|`int64`|
|temporary_status|`utf8`|
|unread_message_count|`int64`|
|unread_notif_count|`int64`|
|unseen_message_count|`int64`|
|username|`utf8`|
|verification_status|`utf8`|
|voip_info|`json`|
|website|`utf8`|
|were_here_count|`int64`|
|whatsapp_number|`utf8`|
|written_by|`utf8`|