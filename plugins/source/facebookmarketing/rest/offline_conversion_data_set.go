package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type OfflineConversionDataSet struct {
	AutomaticMatchingFields    []string       `json:"automatic_matching_fields"`
	Business                   map[string]any `json:"business"`
	CanProxy                   *bool          `json:"can_proxy"`
	Config                     string         `json:"config"`
	CreationTime               string         `json:"creation_time" datetime:"true"`
	Creator                    map[string]any `json:"creator"`
	DataUseSetting             string         `json:"data_use_setting"`
	Description                string         `json:"description"`
	DuplicateEntries           *int64         `json:"duplicate_entries"`
	EnableAutoAssignToAccounts *bool          `json:"enable_auto_assign_to_accounts"`
	EnableAutomaticMatching    *bool          `json:"enable_automatic_matching"`
	EventStats                 string         `json:"event_stats"`
	EventTimeMax               *int64         `json:"event_time_max"`
	EventTimeMin               *int64         `json:"event_time_min"`
	FirstPartyCookieStatus     string         `json:"first_party_cookie_status"`
	Id                         string         `json:"id"`
	IsConsolidatedContainer    *bool          `json:"is_consolidated_container"`
	IsCreatedByBusiness        *bool          `json:"is_created_by_business"`
	IsCrm                      *bool          `json:"is_crm"`
	IsMtaUse                   *bool          `json:"is_mta_use"`
	IsRestrictedUse            *bool          `json:"is_restricted_use"`
	IsUnavailable              *bool          `json:"is_unavailable"`
	LastFiredTime              string         `json:"last_fired_time" datetime:"true"`
	LastUploadApp              string         `json:"last_upload_app"`
	LastUploadAppChangedTime   *int64         `json:"last_upload_app_changed_time"`
	MatchRateApprox            *int64         `json:"match_rate_approx"`
	MatchedEntries             *int64         `json:"matched_entries"`
	Name                       string         `json:"name"`
	OwnerAdAccount             map[string]any `json:"owner_ad_account"`
	OwnerBusiness              map[string]any `json:"owner_business"`
	Usage                      map[string]any `json:"usage"`
	ValidEntries               *int64         `json:"valid_entries"`
}

type OfflineConversionDataSetsResponseStruct struct {
	Data   []OfflineConversionDataSet `json:"data"`
	Paging *Paging                    `json:"paging"`
}

func (facebookClient *FacebookClient) ListOfflineConversionDataSets(ctx context.Context, page string) (items []OfflineConversionDataSet, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(OfflineConversionDataSet{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "offline_conversion_data_sets")
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

	var responseStruct OfflineConversionDataSetsResponseStruct
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
