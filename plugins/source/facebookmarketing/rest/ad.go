package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Ad struct {
	AccountId                         string         `json:"account_id"`
	AdReviewFeedback                  map[string]any `json:"ad_review_feedback"`
	Adlabels                          []any          `json:"adlabels"`
	Adset                             map[string]any `json:"adset"`
	AdsetId                           string         `json:"adset_id"`
	BidAmount                         *int64         `json:"bid_amount"`
	BidInfo                           map[string]any `json:"bid_info"`
	BidType                           string         `json:"bid_type"`
	Campaign                          map[string]any `json:"campaign"`
	CampaignId                        string         `json:"campaign_id"`
	ConfiguredStatus                  string         `json:"configured_status"`
	ConversionDomain                  string         `json:"conversion_domain"`
	ConversionSpecs                   []any          `json:"conversion_specs"`
	CreatedTime                       string         `json:"created_time" datetime:"true"`
	Creative                          map[string]any `json:"creative"`
	DemolinkHash                      string         `json:"demolink_hash"`
	DisplaySequence                   *int64         `json:"display_sequence"`
	EffectiveStatus                   string         `json:"effective_status"`
	EngagementAudience                *bool          `json:"engagement_audience"`
	FailedDeliveryChecks              []any          `json:"failed_delivery_checks"`
	Id                                string         `json:"id"`
	IssuesInfo                        []any          `json:"issues_info"`
	LastUpdatedByAppId                string         `json:"last_updated_by_app_id"`
	Name                              string         `json:"name"`
	PreviewShareableLink              string         `json:"preview_shareable_link"`
	Priority                          *int64         `json:"priority"`
	Recommendations                   []any          `json:"recommendations"`
	SourceAdId                        string         `json:"source_ad_id"`
	Status                            string         `json:"status"`
	Targeting                         map[string]any `json:"targeting"`
	TrackingAndConversionWithDefaults map[string]any `json:"tracking_and_conversion_with_defaults"`
	TrackingSpecs                     []any          `json:"tracking_specs"`
	UpdatedTime                       string         `json:"updated_time" datetime:"true"`
}

type AdsResponseStruct struct {
	Data   []Ad    `json:"data"`
	Paging *Paging `json:"paging"`
}

func (facebookClient *FacebookClient) ListAds(ctx context.Context, page string) (items []Ad, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(Ad{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "ads")
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

	var responseStruct AdsResponseStruct
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
