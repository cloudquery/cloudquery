package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Adaccount struct {
	AccountId     string `json:"account_id"`
	AccountStatus *int64 `json:"account_status"`
	// AdAccountPromotableObjects map[string]any `json:"ad_account_promotable_objects"`
	Age                     *float64       `json:"age"`
	AgencyClientDeclaration map[string]any `json:"agency_client_declaration"`
	AmountSpent             string         `json:"amount_spent"`
	AttributionSpec         []any          `json:"attribution_spec"`
	Balance                 string         `json:"balance"`
	Business                map[string]any `json:"business"`
	BusinessCity            string         `json:"business_city"`
	BusinessCountryCode     string         `json:"business_country_code"`
	BusinessName            string         `json:"business_name"`
	BusinessState           string         `json:"business_state"`
	BusinessStreet          string         `json:"business_street"`
	BusinessStreet2         string         `json:"business_street2"`
	BusinessZip             string         `json:"business_zip"`
	CanCreateBrandLiftStudy *bool          `json:"can_create_brand_lift_study"`
	Capabilities            []string       `json:"capabilities"`
	CreatedTime             string         `json:"created_time" datetime:"true"`
	Currency                string         `json:"currency"`
	// CustomAudienceInfo map[string]any `json:"custom_audience_info"`
	DisableReason     *int64 `json:"disable_reason"`
	EndAdvertiser     string `json:"end_advertiser"`
	EndAdvertiserName string `json:"end_advertiser_name"`
	// ExistingCustomers          []string       `json:"existing_customers"`
	ExtendedCreditInvoiceGroup map[string]any `json:"extended_credit_invoice_group"`
	FailedDeliveryChecks       []any          `json:"failed_delivery_checks"`
	FbEntity                   *int64         `json:"fb_entity"`
	FundingSource              string         `json:"funding_source"`
	FundingSourceDetails       map[string]any `json:"funding_source_details"`
	HasAdvertiserOptedInOdax   *bool          `json:"has_advertiser_opted_in_odax"`
	HasMigratedPermissions     *bool          `json:"has_migrated_permissions"`
	// HasPageAuthorizedAdaccount        *bool          `json:"has_page_authorized_adaccount"`
	Id                                string `json:"id"`
	IoNumber                          string `json:"io_number"`
	IsAttributionSpecSystemDefault    *bool  `json:"is_attribution_spec_system_default"`
	IsDirectDealsEnabled              *bool  `json:"is_direct_deals_enabled"`
	IsIn3dsAuthorizationEnabledMarket *bool  `json:"is_in_3ds_authorization_enabled_market"`
	IsNotificationsEnabled            *bool  `json:"is_notifications_enabled"`
	IsPersonal                        *int64 `json:"is_personal"`
	IsPrepayAccount                   *bool  `json:"is_prepay_account"`
	IsTaxIdRequired                   *bool  `json:"is_tax_id_required"`
	// LiableAddress                     map[string]any `json:"liable_address"`
	LineNumbers              []int64 `json:"line_numbers"`
	MediaAgency              string  `json:"media_agency"`
	MinCampaignGroupSpendCap string  `json:"min_campaign_group_spend_cap"`
	MinDailyBudget           *int64  `json:"min_daily_budget"`
	Name                     string  `json:"name"`
	OffsitePixelsTosAccepted *bool   `json:"offsite_pixels_tos_accepted"`
	Owner                    string  `json:"owner"`
	// OwnerBusiness                     map[string]any `json:"owner_business"`
	Partner string         `json:"partner"`
	RfSpec  map[string]any `json:"rf_spec"`
	// SendBillToAddress      map[string]any `json:"send_bill_to_address"`
	// ShowCheckoutExperience *bool          `json:"show_checkout_experience"`
	// SoldToAddress          map[string]any `json:"sold_to_address"`
	SpendCap               string         `json:"spend_cap"`
	TaxId                  string         `json:"tax_id"`
	TaxIdStatus            *int64         `json:"tax_id_status"`
	TaxIdType              string         `json:"tax_id_type"`
	TimezoneId             *int64         `json:"timezone_id"`
	TimezoneName           string         `json:"timezone_name"`
	TimezoneOffsetHoursUtc *float64       `json:"timezone_offset_hours_utc"`
	TosAccepted            map[string]any `json:"tos_accepted"`
	UserTasks              []string       `json:"user_tasks"`
	UserTosAccepted        map[string]any `json:"user_tos_accepted"`
	// ViewableBusiness       map[string]any `json:"viewable_business"`
}

// type AdaccountsResponseStruct struct {
// 	Data   []Adaccount `json:"data"`
// 	Paging *Paging     `json:"paging"`
// }

func (facebookClient *FacebookClient) GetAdaccount(ctx context.Context) (*Adaccount, error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(Adaccount{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId)
	if err != nil {
		return nil, err
	}

	u := url.URL{
		Scheme:   "https",
		Host:     "graph.facebook.com",
		Path:     path,
		RawQuery: query.Encode(),
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String() /* body */, nil)
	if err != nil {
		return nil, sanitizeUrlError(err)
	}

	response, err := facebookClient.httpClient.Do(request)
	if err != nil {
		return nil, sanitizeUrlError(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, httpErrorToGolangError(response)
	}

	var adaccount Adaccount
	err = json.NewDecoder(response.Body).Decode(&adaccount)

	if err != nil {
		return nil, err
	}

	return &adaccount, nil
}
