package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Adimage struct {
	AccountId                       string   `json:"account_id"`
	CreatedTime                     string   `json:"created_time" datetime:"true"`
	Creatives                       []string `json:"creatives"`
	Hash                            string   `json:"hash"`
	Height                          *int64   `json:"height"`
	Id                              string   `json:"id"`
	IsAssociatedCreativesInAdgroups *bool    `json:"is_associated_creatives_in_adgroups"`
	Name                            string   `json:"name"`
	OriginalHeight                  *int64   `json:"original_height"`
	OriginalWidth                   *int64   `json:"original_width"`
	// OwnerBusiness                   map[string]any `json:"owner_business"` nonexistant field
	PermalinkUrl string `json:"permalink_url"`
	Status       string `json:"status"`
	UpdatedTime  string `json:"updated_time" datetime:"true"`
	Url          string `json:"url"`
	Url128       string `json:"url_128"`
	Width        *int64 `json:"width"`
}

type AdimagesResponseStruct struct {
	Data   []Adimage `json:"data"`
	Paging *Paging   `json:"paging"`
}

func (facebookClient *FacebookClient) ListAdimages(ctx context.Context, page string) (items []Adimage, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(Adimage{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "adimages")
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

	var responseStruct AdimagesResponseStruct
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
