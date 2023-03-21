package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Customaudience struct {
	AccountId                             string         `json:"account_id"`
	ApproximateCountLowerBound            *int64         `json:"approximate_count_lower_bound"`
	ApproximateCountUpperBound            *int64         `json:"approximate_count_upper_bound"`
	CustomerFileSource                    string         `json:"customer_file_source"`
	DataSource                            map[string]any `json:"data_source"`
	DataSourceTypes                       string         `json:"data_source_types"`
	DatafileCustomAudienceUploadingStatus string         `json:"datafile_custom_audience_uploading_status"`
	DeleteTime                            *int64         `json:"delete_time"`
	DeliveryStatus                        map[string]any `json:"delivery_status"`
	Description                           string         `json:"description"`
	ExcludedCustomAudiences               []any          `json:"excluded_custom_audiences"`
	// ExternalEventSource map[string]any `json:"external_event_source"`
	HouseholdAudience            *int64         `json:"household_audience"`
	Id                           string         `json:"id"`
	IncludedCustomAudiences      []any          `json:"included_custom_audiences"`
	IsHousehold                  *bool          `json:"is_household"`
	IsSnapshot                   *bool          `json:"is_snapshot"`
	IsValueBased                 *bool          `json:"is_value_based"`
	LookalikeAudienceIds         []string       `json:"lookalike_audience_ids"`
	LookalikeSpec                map[string]any `json:"lookalike_spec"`
	Name                         string         `json:"name"`
	OperationStatus              map[string]any `json:"operation_status"`
	OptOutLink                   string         `json:"opt_out_link"`
	OwnerBusiness                map[string]any `json:"owner_business"`
	PageDeletionMarkedDeleteTime *int64         `json:"page_deletion_marked_delete_time"`
	PermissionForActions         map[string]any `json:"permission_for_actions"`
	PixelId                      string         `json:"pixel_id"`
	RegulatedAudienceSpec        map[string]any `json:"regulated_audience_spec"`
	RetentionDays                *int64         `json:"retention_days"`
	RevSharePolicyId             *int64         `json:"rev_share_policy_id"`
	Rule                         string         `json:"rule"`
	RuleAggregation              string         `json:"rule_aggregation"`
	RuleV2                       string         `json:"rule_v2"`
	SeedAudience                 *int64         `json:"seed_audience"`
	SharingStatus                map[string]any `json:"sharing_status"`
	Subtype                      string         `json:"subtype"`
	TimeContentUpdated           *int64         `json:"time_content_updated"`
	TimeCreated                  *int64         `json:"time_created"`
	TimeUpdated                  *int64         `json:"time_updated"`
}

type CustomaudiencesResponseStruct struct {
	Data   []Customaudience `json:"data"`
	Paging *Paging          `json:"paging"`
}

func (facebookClient *FacebookClient) ListCustomaudiences(ctx context.Context, page string) (items []Customaudience, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(Customaudience{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "customaudiences")
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

	var responseStruct CustomaudiencesResponseStruct
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
