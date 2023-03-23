package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Adset struct {
	AccountId           string         `json:"account_id"`
	Adlabels            []any          `json:"adlabels"`
	AdsetSchedule       []any          `json:"adset_schedule"`
	AssetFeedId         string         `json:"asset_feed_id"`
	AttributionSpec     []any          `json:"attribution_spec"`
	BidAdjustments      map[string]any `json:"bid_adjustments"`
	BidAmount           *int64         `json:"bid_amount"`
	BidConstraints      map[string]any `json:"bid_constraints"`
	BidInfo             map[string]any `json:"bid_info"`
	BidStrategy         string         `json:"bid_strategy"`
	BillingEvent        string         `json:"billing_event"`
	BudgetRemaining     string         `json:"budget_remaining"`
	Campaign            map[string]any `json:"campaign"`
	CampaignAttribution string         `json:"campaign_attribution"`
	CampaignId          string         `json:"campaign_id"`
	ConfiguredStatus    string         `json:"configured_status"`
	CreatedTime         string         `json:"created_time" datetime:"true"`
	CreativeSequence    []string       `json:"creative_sequence"`
	DailyBudget         string         `json:"daily_budget"`
	DailyMinSpendTarget string         `json:"daily_min_spend_target"`
	DailySpendCap       string         `json:"daily_spend_cap"`
	DestinationType     string         `json:"destination_type"`
	EffectiveStatus     string         `json:"effective_status"`
	EndTime             string         `json:"end_time" datetime:"true"`
	// ExistingCustomerBudgetPercentage *int64 `json:"existing_customer_budget_percentage"` https://github.com/facebook/facebook-java-business-sdk/issues/325
	FrequencyControlSpecs []any `json:"frequency_control_specs"`
	// FullFunnelExplorationMode string `json:"full_funnel_exploration_mode"` https://github.com/facebook/facebook-java-business-sdk/issues/325
	Id string `json:"id"`
	// InstagramActorId string `json:"instagram_actor_id"` https://github.com/facebook/facebook-java-business-sdk/issues/325
	IsDynamicCreative           *bool          `json:"is_dynamic_creative"`
	IssuesInfo                  []any          `json:"issues_info"`
	LearningStageInfo           map[string]any `json:"learning_stage_info"`
	LifetimeBudget              string         `json:"lifetime_budget"`
	LifetimeImps                *int64         `json:"lifetime_imps"`
	LifetimeMinSpendTarget      string         `json:"lifetime_min_spend_target"`
	LifetimeSpendCap            string         `json:"lifetime_spend_cap"`
	MultiOptimizationGoalWeight string         `json:"multi_optimization_goal_weight"`
	Name                        string         `json:"name"`
	// OptimizationGoal string `json:"optimization_goal"` https://github.com/facebook/facebook-java-business-sdk/issues/325
	OptimizationSubEvent     string         `json:"optimization_sub_event"`
	PacingType               []string       `json:"pacing_type"`
	PromotedObject           map[string]any `json:"promoted_object"`
	Recommendations          []any          `json:"recommendations"`
	RecurringBudgetSemantics *bool          `json:"recurring_budget_semantics"`
	ReviewFeedback           string         `json:"review_feedback"`
	RfPredictionId           string         `json:"rf_prediction_id"`
	SourceAdsetId            string         `json:"source_adset_id"`
	StartTime                string         `json:"start_time" datetime:"true"`
	Status                   string         `json:"status"`
	Targeting                map[string]any `json:"targeting"`
	// TargetingOptimizationTypes   map[string]any `json:"targeting_optimization_types"` wrong type declared in facebook-business-sdk.
	TimeBasedAdRotationIdBlocks  []any   `json:"time_based_ad_rotation_id_blocks"`
	TimeBasedAdRotationIntervals []int64 `json:"time_based_ad_rotation_intervals"`
	UpdatedTime                  string  `json:"updated_time" datetime:"true"`
	UseNewAppClick               *bool   `json:"use_new_app_click"`
}

type AdsetsResponseStruct struct {
	Data   []Adset `json:"data"`
	Paging *Paging `json:"paging"`
}

func (facebookClient *FacebookClient) ListAdsets(ctx context.Context, page string) (items []Adset, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(Adset{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "adsets")
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

	var responseStruct AdsetsResponseStruct
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
