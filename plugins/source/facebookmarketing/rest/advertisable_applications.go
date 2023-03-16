package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Application struct {
	AamRules                           string   `json:"aam_rules"`
	AnAdSpaceLimit                     *int64   `json:"an_ad_space_limit"`
	AnPlatforms                        []string `json:"an_platforms"`
	AndroidKeyHash                     []string `json:"android_key_hash"`
	AndroidSdkErrorCategories          any      `json:"android_sdk_error_categories"`
	AppDomains                         []string `json:"app_domains"`
	AppEventsConfig                    any      `json:"app_events_config"`
	AppEventsFeatureBitmask            *int64   `json:"app_events_feature_bitmask"`
	AppEventsSessionTimeout            *int64   `json:"app_events_session_timeout"`
	AppInstallTracked                  *bool    `json:"app_install_tracked"`
	AppName                            string   `json:"app_name"`
	AppSignalsBindingIos               any      `json:"app_signals_binding_ios"`
	AppType                            *int64   `json:"app_type"`
	AuthDialogDataHelpUrl              string   `json:"auth_dialog_data_help_url"`
	AuthDialogHeadline                 string   `json:"auth_dialog_headline"`
	AuthDialogPermsExplanation         string   `json:"auth_dialog_perms_explanation"`
	AuthReferralDefaultActivityPrivacy string   `json:"auth_referral_default_activity_privacy"`
	AuthReferralEnabled                *int64   `json:"auth_referral_enabled"`
	AuthReferralExtendedPerms          []string `json:"auth_referral_extended_perms"`
	AuthReferralFriendPerms            []string `json:"auth_referral_friend_perms"`
	AuthReferralResponseType           string   `json:"auth_referral_response_type"`
	AuthReferralUserPerms              []string `json:"auth_referral_user_perms"`
	AutoEventMappingAndroid            any      `json:"auto_event_mapping_android"`
	AutoEventMappingIos                any      `json:"auto_event_mapping_ios"`
	AutoEventSetupEnabled              *bool    `json:"auto_event_setup_enabled"`
	// Business                           map[string]any `json:"business"`
	CanvasFluidHeight *bool  `json:"canvas_fluid_height"`
	CanvasFluidWidth  *int64 `json:"canvas_fluid_width"`
	CanvasUrl         string `json:"canvas_url"`
	Category          string `json:"category"`
	ClientConfig      any    `json:"client_config"`
	Company           string `json:"company"`
	ConfiguredIosSso  *bool  `json:"configured_ios_sso"`
	ContactEmail      string `json:"contact_email"`
	CreatedTime       string `json:"created_time" datetime:"true"`
	CreatorUid        string `json:"creator_uid"`
	// DailyActiveUsers     string `json:"daily_active_users"`
	// DailyActiveUsersRank *int64 `json:"daily_active_users_rank"`
	DeauthCallbackUrl              string         `json:"deauth_callback_url"`
	DefaultShareMode               string         `json:"default_share_mode"`
	Description                    string         `json:"description"`
	FinancialId                    string         `json:"financial_id"`
	Gdpv4ChromeCustomTabsEnabled   *bool          `json:"gdpv4_chrome_custom_tabs_enabled"`
	Gdpv4Enabled                   *bool          `json:"gdpv4_enabled"`
	Gdpv4NuxContent                string         `json:"gdpv4_nux_content"`
	Gdpv4NuxEnabled                *bool          `json:"gdpv4_nux_enabled"`
	HasMessengerProduct            *bool          `json:"has_messenger_product"`
	HostingUrl                     string         `json:"hosting_url"`
	IconUrl                        string         `json:"icon_url"`
	Id                             string         `json:"id"`
	IosBundleId                    []string       `json:"ios_bundle_id"`
	IosSdkDialogFlows              any            `json:"ios_sdk_dialog_flows"`
	IosSdkErrorCategories          any            `json:"ios_sdk_error_categories"`
	IosSfvcAttr                    *bool          `json:"ios_sfvc_attr"`
	IosSupportsNativeProxyAuthFlow *bool          `json:"ios_supports_native_proxy_auth_flow"`
	IosSupportsSystemAuth          *bool          `json:"ios_supports_system_auth"`
	IpadAppStoreId                 string         `json:"ipad_app_store_id"`
	IphoneAppStoreId               string         `json:"iphone_app_store_id"`
	LatestSdkVersion               any            `json:"latest_sdk_version"`
	Link                           string         `json:"link"`
	LoggingToken                   string         `json:"logging_token"`
	LogoUrl                        string         `json:"logo_url"`
	Migrations                     map[string]any `json:"migrations"`
	MobileProfileSectionUrl        string         `json:"mobile_profile_section_url"`
	MobileWebUrl                   string         `json:"mobile_web_url"`
	// MonthlyActiveUsers               string         `json:"monthly_active_users"`
	// MonthlyActiveUsersRank           *int64         `json:"monthly_active_users_rank"`
	Name            string `json:"name"`
	Namespace       string `json:"namespace"`
	ObjectStoreUrls any    `json:"object_store_urls"`
	// OwnerBusiness                    map[string]any `json:"owner_business"`
	PageTabDefaultName               string   `json:"page_tab_default_name"`
	PageTabUrl                       string   `json:"page_tab_url"`
	PhotoUrl                         string   `json:"photo_url"`
	PrivacyPolicyUrl                 string   `json:"privacy_policy_url"`
	ProfileSectionUrl                string   `json:"profile_section_url"`
	PropertyId                       string   `json:"property_id"`
	RealTimeModeDevices              []string `json:"real_time_mode_devices"`
	Restrictions                     any      `json:"restrictions"`
	RestrictiveDataFilterParams      string   `json:"restrictive_data_filter_params"`
	RestrictiveDataFilterRules       string   `json:"restrictive_data_filter_rules"`
	SdkUpdateMessage                 string   `json:"sdk_update_message"`
	SeamlessLogin                    *int64   `json:"seamless_login"`
	SecureCanvasUrl                  string   `json:"secure_canvas_url"`
	SecurePageTabUrl                 string   `json:"secure_page_tab_url"`
	ServerIpWhitelist                string   `json:"server_ip_whitelist"`
	SmartLoginBookmarkIconUrl        string   `json:"smart_login_bookmark_icon_url"`
	SmartLoginMenuIconUrl            string   `json:"smart_login_menu_icon_url"`
	SocialDiscovery                  *int64   `json:"social_discovery"`
	Subcategory                      string   `json:"subcategory"`
	SuggestedEventsSetting           string   `json:"suggested_events_setting"`
	SupportedPlatforms               []string `json:"supported_platforms"`
	SupportsApprequestsFastAppSwitch any      `json:"supports_apprequests_fast_app_switch"`
	SupportsAttribution              *bool    `json:"supports_attribution"`
	SupportsImplicitSdkLogging       *bool    `json:"supports_implicit_sdk_logging"`
	SuppressNativeIosGdp             *bool    `json:"suppress_native_ios_gdp"`
	TermsOfServiceUrl                string   `json:"terms_of_service_url"`
	UrlSchemeSuffix                  string   `json:"url_scheme_suffix"`
	UserSupportEmail                 string   `json:"user_support_email"`
	UserSupportUrl                   string   `json:"user_support_url"`
	WebsiteUrl                       string   `json:"website_url"`
	// WeeklyActiveUsers string `json:"weekly_active_users"`
}

type AdvertisableApplicationssResponseStruct struct {
	Data   []Application `json:"data"`
	Paging *Paging       `json:"paging"`
}

func (facebookClient *FacebookClient) ListAdvertisableApplications(ctx context.Context, page string) (items []Application, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(Application{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "advertisable_applications")
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

	var responseStruct AdvertisableApplicationssResponseStruct
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
