package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type ReachFrequencyPrediction struct {
	AccountId                       *int64         `json:"account_id"`
	ActivityStatus                  map[string]any `json:"activity_status"`
	AdFormats                       []any          `json:"ad_formats"`
	AuctionEntryOptionIndex         *int64         `json:"auction_entry_option_index"`
	AudienceSizeLowerBound          *int64         `json:"audience_size_lower_bound"`
	AudienceSizeUpperBound          *int64         `json:"audience_size_upper_bound"`
	BusinessId                      *int64         `json:"business_id"`
	BuyingType                      string         `json:"buying_type"`
	CampaignGroupId                 *int64         `json:"campaign_group_id"`
	CampaignId                      string         `json:"campaign_id"`
	CampaignTimeStart               string         `json:"campaign_time_start" datetime:"true"`
	CampaignTimeStop                string         `json:"campaign_time_stop" datetime:"true"`
	Currency                        string         `json:"currency"`
	CurveBudgetReach                map[string]any `json:"curve_budget_reach"`
	CurveReach                      []int64        `json:"curve_reach"`
	DailyGrpCurve                   []any          `json:"daily_grp_curve"`
	DailyImpressionCurve            []any          `json:"daily_impression_curve"`
	DailyImpressionCurveMap         any            `json:"daily_impression_curve_map"`
	DayPartingSchedule              []any          `json:"day_parting_schedule"`
	DestinationId                   string         `json:"destination_id"`
	EndTime                         string         `json:"end_time" datetime:"true"`
	ExpirationTime                  string         `json:"expiration_time" datetime:"true"`
	ExternalBudget                  *int64         `json:"external_budget"`
	ExternalImpression              *int64         `json:"external_impression"`
	ExternalMaximumBudget           *int64         `json:"external_maximum_budget"`
	ExternalMaximumImpression       string         `json:"external_maximum_impression"`
	ExternalMaximumReach            *int64         `json:"external_maximum_reach"`
	ExternalMinimumBudget           *int64         `json:"external_minimum_budget"`
	ExternalMinimumImpression       *int64         `json:"external_minimum_impression"`
	ExternalMinimumReach            *int64         `json:"external_minimum_reach"`
	ExternalReach                   *int64         `json:"external_reach"`
	FeedRatio0000                   *int64         `json:"feed_ratio_0000"`
	FrequencyCap                    *int64         `json:"frequency_cap"`
	FrequencyDistributionMap        any            `json:"frequency_distribution_map"`
	FrequencyDistributionMapAgg     any            `json:"frequency_distribution_map_agg"`
	GrpAudienceSize                 *float64       `json:"grp_audience_size"`
	GrpAvgProbabilityMap            string         `json:"grp_avg_probability_map"`
	GrpCountryAudienceSize          *float64       `json:"grp_country_audience_size"`
	GrpCurve                        []any          `json:"grp_curve"`
	GrpDmasAudienceSize             *float64       `json:"grp_dmas_audience_size"`
	GrpFilteringThreshold00         *int64         `json:"grp_filtering_threshold_00"`
	GrpPoints                       *float64       `json:"grp_points"`
	GrpRatio                        *float64       `json:"grp_ratio"`
	GrpReachRatio                   *float64       `json:"grp_reach_ratio"`
	GrpStatus                       string         `json:"grp_status"`
	HoldoutPercentage               *int64         `json:"holdout_percentage"`
	Id                              string         `json:"id"`
	ImpressionCurve                 []int64        `json:"impression_curve"`
	InstagramDestinationId          string         `json:"instagram_destination_id"`
	InstreamPackages                []string       `json:"instream_packages"`
	IntervalFrequencyCap            *int64         `json:"interval_frequency_cap"`
	IntervalFrequencyCapResetPeriod *int64         `json:"interval_frequency_cap_reset_period"`
	IsBonusMedia                    *int64         `json:"is_bonus_media"`
	IsConversionGoal                *int64         `json:"is_conversion_goal"`
	IsHigherAverageFrequency        *bool          `json:"is_higher_average_frequency"`
	IsIo                            *bool          `json:"is_io"`
	IsReservedBuying                *int64         `json:"is_reserved_buying"`
	IsTrp                           *bool          `json:"is_trp"`
	Name                            string         `json:"name"`
	Objective                       *int64         `json:"objective"`
	ObjectiveName                   string         `json:"objective_name"`
	OdaxObjective                   *int64         `json:"odax_objective"`
	OdaxObjectiveName               string         `json:"odax_objective_name"`
	OptimizationGoal                *int64         `json:"optimization_goal"`
	OptimizationGoalName            string         `json:"optimization_goal_name"`
	PausePeriods                    any            `json:"pause_periods"`
	PlacementBreakdown              map[string]any `json:"placement_breakdown"`
	PlacementBreakdownMap           any            `json:"placement_breakdown_map"`
	PlanName                        string         `json:"plan_name"`
	PlanType                        string         `json:"plan_type"`
	PredictionMode                  *int64         `json:"prediction_mode"`
	PredictionProgress              *int64         `json:"prediction_progress"`
	ReferenceId                     string         `json:"reference_id"`
	ReservationStatus               *int64         `json:"reservation_status"`
	StartTime                       string         `json:"start_time" datetime:"true"`
	Status                          *int64         `json:"status"`
	StoryEventType                  *int64         `json:"story_event_type"`
	TargetCpm                       *int64         `json:"target_cpm"`
	TargetSpec                      map[string]any `json:"target_spec"`
	TimeCreated                     string         `json:"time_created" datetime:"true"`
	TimeUpdated                     string         `json:"time_updated" datetime:"true"`
	TimezoneId                      *int64         `json:"timezone_id"`
	TimezoneName                    string         `json:"timezone_name"`
	ToplineId                       *int64         `json:"topline_id"`
	VideoViewLengthConstraint       *int64         `json:"video_view_length_constraint"`
	Viewtag                         string         `json:"viewtag"`
}

type ReachFrequencyPredictionsResponseStruct struct {
	Data   []ReachFrequencyPrediction `json:"data"`
	Paging *Paging                    `json:"paging"`
}

func (facebookClient *FacebookClient) ListReachFrequencyPredictions(ctx context.Context, page string) (items []ReachFrequencyPrediction, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(ReachFrequencyPrediction{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "reachfrequencypredictions")
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

	var responseStruct ReachFrequencyPredictionsResponseStruct
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
