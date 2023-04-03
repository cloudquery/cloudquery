package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type AdStudy struct {
	Business                  map[string]any `json:"business"`
	CanceledTime              string         `json:"canceled_time" datetime:"true"`
	ClientBusiness            map[string]any `json:"client_business"`
	CooldownStartTime         string         `json:"cooldown_start_time" datetime:"true"`
	CreatedBy                 map[string]any `json:"created_by"`
	CreatedTime               string         `json:"created_time" datetime:"true"`
	Description               string         `json:"description"`
	EndTime                   string         `json:"end_time" datetime:"true"`
	Id                        string         `json:"id"`
	MeasurementContact        map[string]any `json:"measurement_contact"`
	Name                      string         `json:"name"`
	ObservationEndTime        string         `json:"observation_end_time" datetime:"true"`
	ResultsFirstAvailableDate string         `json:"results_first_available_date"`
	SalesContact              map[string]any `json:"sales_contact"`
	StartTime                 string         `json:"start_time" datetime:"true"`
	Type                      string         `json:"type"`
	UpdatedBy                 map[string]any `json:"updated_by"`
	UpdatedTime               string         `json:"updated_time" datetime:"true"`
}

type AdStudysResponseStruct struct {
	Data   []AdStudy `json:"data"`
	Paging *Paging   `json:"paging"`
}

func (facebookClient *FacebookClient) ListAdStudies(ctx context.Context, page string) (items []AdStudy, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(AdStudy{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "ad_studies")
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

	var responseStruct AdStudysResponseStruct
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
