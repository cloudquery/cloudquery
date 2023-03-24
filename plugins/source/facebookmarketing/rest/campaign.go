package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Campaign struct {
	AccountId                        string         `json:"account_id"`
	AdStrategyGroupId                string         `json:"ad_strategy_group_id"`
	AdStrategyId                     string         `json:"ad_strategy_id"`
	Adlabels                         []any          `json:"adlabels"`
	BidStrategy                      string         `json:"bid_strategy"`
	BoostedObjectId                  string         `json:"boosted_object_id"`
	BrandLiftStudies                 []any          `json:"brand_lift_studies"`
	BudgetRebalanceFlag              *bool          `json:"budget_rebalance_flag"`
	BudgetRemaining                  string         `json:"budget_remaining"`
	BuyingType                       string         `json:"buying_type"`
	CanCreateBrandLiftStudy          *bool          `json:"can_create_brand_lift_study"`
	CanUseSpendCap                   *bool          `json:"can_use_spend_cap"`
	ConfiguredStatus                 string         `json:"configured_status"`
	CreatedTime                      string         `json:"created_time" datetime:"true"`
	DailyBudget                      string         `json:"daily_budget"`
	EffectiveStatus                  string         `json:"effective_status"`
	HasSecondarySkadnetworkReporting *bool          `json:"has_secondary_skadnetwork_reporting"`
	Id                               string         `json:"id"`
	IsSkadnetworkAttribution         *bool          `json:"is_skadnetwork_attribution"`
	IssuesInfo                       []any          `json:"issues_info"`
	LastBudgetTogglingTime           string         `json:"last_budget_toggling_time" datetime:"true"`
	LifetimeBudget                   string         `json:"lifetime_budget"`
	Name                             string         `json:"name"`
	Objective                        string         `json:"objective"`
	PacingType                       []string       `json:"pacing_type"`
	PrimaryAttribution               string         `json:"primary_attribution"`
	PromotedObject                   map[string]any `json:"promoted_object"`
	Recommendations                  []any          `json:"recommendations"`
	SmartPromotionType               string         `json:"smart_promotion_type"`
	SourceCampaignId                 string         `json:"source_campaign_id"`
	SpecialAdCategories              []string       `json:"special_ad_categories"`
	SpecialAdCategory                string         `json:"special_ad_category"`
	SpecialAdCategoryCountry         []string       `json:"special_ad_category_country"`
	SpendCap                         string         `json:"spend_cap"`
	StartTime                        string         `json:"start_time" datetime:"true"`
	Status                           string         `json:"status"`
	StopTime                         string         `json:"stop_time" datetime:"true"`
	ToplineId                        string         `json:"topline_id"`
	UpdatedTime                      string         `json:"updated_time" datetime:"true"`
}

type CampaignsResponseStruct struct {
	Data   []Campaign `json:"data"`
	Paging *Paging    `json:"paging"`
}

func (facebookClient *FacebookClient) ListCampaigns(ctx context.Context, page string) (items []Campaign, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(Campaign{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "campaigns")
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

	var responseStruct CampaignsResponseStruct
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
