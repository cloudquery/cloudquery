package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Adrule struct {
	AccountId      string         `json:"account_id"`
	CreatedBy      map[string]any `json:"created_by"`
	CreatedTime    string         `json:"created_time" datetime:"true"`
	EvaluationSpec map[string]any `json:"evaluation_spec"`
	ExecutionSpec  map[string]any `json:"execution_spec"`
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	ScheduleSpec   map[string]any `json:"schedule_spec"`
	Status         string         `json:"status"`
	UpdatedTime    string         `json:"updated_time" datetime:"true"`
}

type AdrulesResponseStruct struct {
	Data   []Adrule `json:"data"`
	Paging *Paging  `json:"paging"`
}

func (facebookClient *FacebookClient) ListAdrules(ctx context.Context, page string) (items []Adrule, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(Adrule{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "adrules_library")
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

	var responseStruct AdrulesResponseStruct
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
