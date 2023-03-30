package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Business struct {
	BlockOfflineAnalytics *bool `json:"block_offline_analytics"`
	// CollaborativeAdsManagedPartnerBusinessInfo map[string]any `json:"collaborative_ads_managed_partner_business_info"`
	CollaborativeAdsManagedPartnerEligibility map[string]any `json:"collaborative_ads_managed_partner_eligibility"`
	CollaborativeAdsPartnerPremiumOptions     map[string]any `json:"collaborative_ads_partner_premium_options"`
	CreatedBy                                 any            `json:"created_by"`
	CreatedTime                               string         `json:"created_time" datetime:"true"`
	ExtendedUpdatedTime                       string         `json:"extended_updated_time" datetime:"true"`
	Id                                        string         `json:"id"`
	IsHidden                                  *bool          `json:"is_hidden"`
	Link                                      string         `json:"link"`
	Name                                      string         `json:"name"`
	PaymentAccountId                          string         `json:"payment_account_id"`
	PrimaryPage                               map[string]any `json:"primary_page"`
	ProfilePictureUri                         string         `json:"profile_picture_uri"`
	TimezoneId                                *int64         `json:"timezone_id"`
	TwoFactorType                             string         `json:"two_factor_type"`
	UpdatedBy                                 any            `json:"updated_by"`
	UpdatedTime                               string         `json:"updated_time" datetime:"true"`
	VerificationStatus                        string         `json:"verification_status"`
	Vertical                                  string         `json:"vertical"`
	VerticalId                                *int64         `json:"vertical_id"`
}

type BusinesssResponseStruct struct {
	Data   []Business `json:"data"`
	Paging *Paging    `json:"paging"`
}

func (facebookClient *FacebookClient) ListBusinesss(ctx context.Context, page string) (items []Business, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(Business{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "agencies")
	if err != nil {
		return nil, "", err
	}

	u := url.URL{
		Scheme:   "https",
		Host:     "graph.facebook.com",
		Path:     path,
		RawQuery: query.Encode(),
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String() /* body */, nil)
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

	var responseStruct BusinesssResponseStruct
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
