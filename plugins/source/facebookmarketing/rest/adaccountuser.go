package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Adaccountuser struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Tasks []string `json:"tasks"`
}

type AdaccountusersResponseStruct struct {
	Data   []Adaccountuser `json:"data"`
	Paging *Paging         `json:"paging"`
}

func (facebookClient *FacebookClient) ListAdaccountusers(ctx context.Context, page string) (items []Adaccountuser, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(Adaccountuser{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "users")
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

	var responseStruct AdaccountusersResponseStruct
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
