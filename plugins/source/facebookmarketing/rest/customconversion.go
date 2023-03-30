package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Customconversion struct {
	AccountId                string         `json:"account_id"`
	AggregationRule          string         `json:"aggregation_rule"`
	Business                 map[string]any `json:"business"`
	CreationTime             string         `json:"creation_time" datetime:"true"`
	CustomEventType          string         `json:"custom_event_type"`
	DataSources              []any          `json:"data_sources"`
	DefaultConversionValue   *int64         `json:"default_conversion_value"`
	Description              string         `json:"description"`
	EventSourceType          string         `json:"event_source_type"`
	FirstFiredTime           string         `json:"first_fired_time" datetime:"true"`
	Id                       string         `json:"id"`
	IsArchived               *bool          `json:"is_archived"`
	IsUnavailable            *bool          `json:"is_unavailable"`
	LastFiredTime            string         `json:"last_fired_time" datetime:"true"`
	Name                     string         `json:"name"`
	OfflineConversionDataSet map[string]any `json:"offline_conversion_data_set"`
	Pixel                    map[string]any `json:"pixel"`
	RetentionDays            *int64         `json:"retention_days"`
	Rule                     string         `json:"rule"`
}

type CustomconversionsResponseStruct struct {
	Data   []Customconversion `json:"data"`
	Paging *Paging            `json:"paging"`
}

func (facebookClient *FacebookClient) ListCustomconversions(ctx context.Context, page string) (items []Customconversion, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(Customconversion{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "customconversions")
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

	var responseStruct CustomconversionsResponseStruct
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
