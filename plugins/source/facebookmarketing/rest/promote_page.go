package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type PromotePage struct {
	About                                 string         `json:"about"`
	AccessToken                           string         `json:"access_token"`
	AdCampaign                            map[string]any `json:"ad_campaign"`
	Affiliation                           string         `json:"affiliation"`
	AppId                                 string         `json:"app_id"`
	ArtistsWeLike                         string         `json:"artists_we_like"`
	Attire                                string         `json:"attire"`
	Awards                                string         `json:"awards"`
	BandInterests                         string         `json:"band_interests"`
	BandMembers                           string         `json:"band_members"`
	BestPage                              map[string]any `json:"best_page"`
	Bio                                   string         `json:"bio"`
	Birthday                              string         `json:"birthday"`
	BookingAgent                          string         `json:"booking_agent"`
	Built                                 string         `json:"built"`
	Business                              any            `json:"business"`
	CanCheckin                            *bool          `json:"can_checkin"`
	CanPost                               *bool          `json:"can_post"`
	Category                              string         `json:"category"`
	CategoryList                          []any          `json:"category_list"`
	Checkins                              *int64         `json:"checkins"`
	CompanyOverview                       string         `json:"company_overview"`
	ConnectedInstagramAccount             map[string]any `json:"connected_instagram_account"`
	ConnectedPageBackedInstagramAccount   map[string]any `json:"connected_page_backed_instagram_account"`
	ContactAddress                        map[string]any `json:"contact_address"`
	CopyrightWhitelistedIgPartners        []string       `json:"copyright_whitelisted_ig_partners"`
	CountryPageLikes                      *int64         `json:"country_page_likes"`
	Cover                                 map[string]any `json:"cover"`
	CulinaryTeam                          string         `json:"culinary_team"`
	CurrentLocation                       string         `json:"current_location"`
	DeliveryAndPickupOptionInfo           []string       `json:"delivery_and_pickup_option_info"`
	Description                           string         `json:"description"`
	DescriptionHtml                       string         `json:"description_html"`
	DifferentlyOpenOfferings              map[string]any `json:"differently_open_offerings"`
	DirectedBy                            string         `json:"directed_by"`
	DisplaySubtext                        string         `json:"display_subtext"`
	DisplayedMessageResponseTime          string         `json:"displayed_message_response_time"`
	Emails                                []string       `json:"emails"`
	Engagement                            map[string]any `json:"engagement"`
	FanCount                              *int64         `json:"fan_count"`
	FeaturedVideo                         map[string]any `json:"featured_video"`
	Features                              string         `json:"features"`
	FollowersCount                        *int64         `json:"followers_count"`
	FoodStyles                            []string       `json:"food_styles"`
	Founded                               string         `json:"founded"`
	GeneralInfo                           string         `json:"general_info"`
	GeneralManager                        string         `json:"general_manager"`
	Genre                                 string         `json:"genre"`
	GlobalBrandPageName                   string         `json:"global_brand_page_name"`
	GlobalBrandRootId                     string         `json:"global_brand_root_id"`
	HasAddedApp                           *bool          `json:"has_added_app"`
	HasTransitionedToNewPageExperience    *bool          `json:"has_transitioned_to_new_page_experience"`
	HasWhatsappBusinessNumber             *bool          `json:"has_whatsapp_business_number"`
	HasWhatsappNumber                     *bool          `json:"has_whatsapp_number"`
	Hometown                              string         `json:"hometown"`
	Hours                                 map[string]any `json:"hours"`
	Id                                    string         `json:"id"`
	Impressum                             string         `json:"impressum"`
	Influences                            string         `json:"influences"`
	InstagramBusinessAccount              map[string]any `json:"instagram_business_account"`
	InstantArticlesReviewStatus           string         `json:"instant_articles_review_status"`
	IsAlwaysOpen                          *bool          `json:"is_always_open"`
	IsChain                               *bool          `json:"is_chain"`
	IsCommunityPage                       *bool          `json:"is_community_page"`
	IsEligibleForBrandedContent           *bool          `json:"is_eligible_for_branded_content"`
	IsMessengerBotGetStartedEnabled       *bool          `json:"is_messenger_bot_get_started_enabled"`
	IsMessengerPlatformBot                *bool          `json:"is_messenger_platform_bot"`
	IsOwned                               *bool          `json:"is_owned"`
	IsPermanentlyClosed                   *bool          `json:"is_permanently_closed"`
	IsPublished                           *bool          `json:"is_published"`
	IsUnclaimed                           *bool          `json:"is_unclaimed"`
	IsVerified                            *bool          `json:"is_verified"`
	IsWebhooksSubscribed                  *bool          `json:"is_webhooks_subscribed"`
	Keywords                              any            `json:"keywords"`
	LeadgenTosAcceptanceTime              string         `json:"leadgen_tos_acceptance_time" datetime:"true"`
	LeadgenTosAccepted                    *bool          `json:"leadgen_tos_accepted"`
	LeadgenTosAcceptingUser               map[string]any `json:"leadgen_tos_accepting_user"`
	Link                                  string         `json:"link"`
	Location                              map[string]any `json:"location"`
	Members                               string         `json:"members"`
	MerchantId                            string         `json:"merchant_id"`
	MerchantReviewStatus                  string         `json:"merchant_review_status"`
	MessagingFeatureStatus                map[string]any `json:"messaging_feature_status"`
	MessengerAdsDefaultIcebreakers        []string       `json:"messenger_ads_default_icebreakers"`
	MessengerAdsDefaultPageWelcomeMessage map[string]any `json:"messenger_ads_default_page_welcome_message"`
	MessengerAdsDefaultQuickReplies       []string       `json:"messenger_ads_default_quick_replies"`
	MessengerAdsQuickRepliesType          string         `json:"messenger_ads_quick_replies_type"`
	MiniShopStorefront                    map[string]any `json:"mini_shop_storefront"`
	Mission                               string         `json:"mission"`
	Mpg                                   string         `json:"mpg"`
	Name                                  string         `json:"name"`
	NameWithLocationDescriptor            string         `json:"name_with_location_descriptor"`
	Network                               string         `json:"network"`
	NewLikeCount                          *int64         `json:"new_like_count"`
	OfferEligible                         *bool          `json:"offer_eligible"`
	OverallStarRating                     *float64       `json:"overall_star_rating"`
	OwnerBusiness                         map[string]any `json:"owner_business"`
	PageToken                             string         `json:"page_token"`
	ParentPage                            map[string]any `json:"parent_page"`
	Parking                               map[string]any `json:"parking"`
	PaymentOptions                        map[string]any `json:"payment_options"`
	PersonalInfo                          string         `json:"personal_info"`
	PersonalInterests                     string         `json:"personal_interests"`
	PharmaSafetyInfo                      string         `json:"pharma_safety_info"`
	Phone                                 string         `json:"phone"`
	PickupOptions                         []string       `json:"pickup_options"`
	PlaceType                             string         `json:"place_type"`
	PlotOutline                           string         `json:"plot_outline"`
	PreferredAudience                     map[string]any `json:"preferred_audience"`
	PressContact                          string         `json:"press_contact"`
	PriceRange                            string         `json:"price_range"`
	PrivacyInfoUrl                        string         `json:"privacy_info_url"`
	ProducedBy                            string         `json:"produced_by"`
	Products                              string         `json:"products"`
	PromotionEligible                     *bool          `json:"promotion_eligible"`
	PromotionIneligibleReason             string         `json:"promotion_ineligible_reason"`
	PublicTransit                         string         `json:"public_transit"`
	RatingCount                           *int64         `json:"rating_count"`
	Recipient                             string         `json:"recipient"`
	RecordLabel                           string         `json:"record_label"`
	ReleaseDate                           string         `json:"release_date"`
	RestaurantServices                    map[string]any `json:"restaurant_services"`
	RestaurantSpecialties                 map[string]any `json:"restaurant_specialties"`
	Schedule                              string         `json:"schedule"`
	ScreenplayBy                          string         `json:"screenplay_by"`
	Season                                string         `json:"season"`
	SingleLineAddress                     string         `json:"single_line_address"`
	Starring                              string         `json:"starring"`
	StartInfo                             map[string]any `json:"start_info"`
	StoreCode                             string         `json:"store_code"`
	StoreLocationDescriptor               string         `json:"store_location_descriptor"`
	StoreNumber                           *int64         `json:"store_number"`
	Studio                                string         `json:"studio"`
	SupportsDonateButtonInLiveVideo       *bool          `json:"supports_donate_button_in_live_video"`
	SupportsInstantArticles               *bool          `json:"supports_instant_articles"`
	TalkingAboutCount                     *int64         `json:"talking_about_count"`
	TemporaryStatus                       string         `json:"temporary_status"`
	UnreadMessageCount                    *int64         `json:"unread_message_count"`
	UnreadNotifCount                      *int64         `json:"unread_notif_count"`
	UnseenMessageCount                    *int64         `json:"unseen_message_count"`
	Username                              string         `json:"username"`
	VerificationStatus                    string         `json:"verification_status"`
	VoipInfo                              map[string]any `json:"voip_info"`
	Website                               string         `json:"website"`
	WereHereCount                         *int64         `json:"were_here_count"`
	WhatsappNumber                        string         `json:"whatsapp_number"`
	WrittenBy                             string         `json:"written_by"`
}

type PromotePagesResponseStruct struct {
	Data   []PromotePage `json:"data"`
	Paging *Paging       `json:"paging"`
}

func (facebookClient *FacebookClient) ListPromotePages(ctx context.Context, page string) (items []PromotePage, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(PromotePage{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "promote_pages")
	if err != nil {
		return nil, "", err
	}

	u := url.URL{
		Scheme:   "https",
		Host:     "graph.facebook.com",
		Path:     path,
		RawQuery: query.Encode(),
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String() /*body*/, nil)
	if err != nil {
		return nil, "", sanitizeUrlError(err)
	}

	response, err := facebookClient.httpClient.Do(request)
	if err != nil {
		return nil, "", sanitizeUrlError(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, "", httpErrorToGolangError(response)
	}

	var responseStruct PromotePagesResponseStruct
	err = json.NewDecoder(response.Body).Decode(&responseStruct)

	if err != nil {
		return nil, "", err
	}

	if responseStruct.Paging != nil && responseStruct.Paging.Next != "" {
		if responseStruct.Paging.Cursors != nil && responseStruct.Paging.Cursors.After != "" {
			return responseStruct.Data, responseStruct.Paging.Cursors.After, nil
		}
	}

	return responseStruct.Data, "", nil
}
