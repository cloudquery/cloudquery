package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type BroadTargetingCategories struct {
	CategoryDescription    string   `json:"category_description"`
	Id                     string   `json:"id"`
	Name                   string   `json:"name"`
	ParentCategory         string   `json:"parent_category"`
	Path                   []string `json:"path"`
	SizeLowerBound         *int64   `json:"size_lower_bound"`
	SizeUpperBound         *int64   `json:"size_upper_bound"`
	Source                 string   `json:"source"`
	Type                   *int64   `json:"type"`
	TypeName               string   `json:"type_name"`
	UntranslatedName       string   `json:"untranslated_name"`
	UntranslatedParentName string   `json:"untranslated_parent_name"`
}

type BroadTargetingCategoriessResponseStruct struct {
	Data   []BroadTargetingCategories `json:"data"`
	Paging *Paging                    `json:"paging"`
}

func (facebookClient *FacebookClient) ListBroadTargetingCategories(ctx context.Context, page string) (items []BroadTargetingCategories, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(BroadTargetingCategories{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "broadtargetingcategories")
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

	var responseStruct BroadTargetingCategoriessResponseStruct
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
