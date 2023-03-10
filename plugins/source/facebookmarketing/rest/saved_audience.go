package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type SavedAudience struct {
	AccountId                    map[string]any `json:"account_id"`
	ApproximateCountLowerBound   *int64         `json:"approximate_count_lower_bound"`
	ApproximateCountUpperBound   *int64         `json:"approximate_count_upper_bound"`
	DeleteTime                   *int64         `json:"delete_time"`
	Description                  string         `json:"description"`
	ExtraInfo                    string         `json:"extra_info"`
	Id                           string         `json:"id"`
	Name                         string         `json:"name"`
	OperationStatus              map[string]any `json:"operation_status"`
	OwnerBusiness                map[string]any `json:"owner_business"`
	PageDeletionMarkedDeleteTime *int64         `json:"page_deletion_marked_delete_time"`
	PermissionForActions         map[string]any `json:"permission_for_actions"`
	RunStatus                    string         `json:"run_status"`
	SentenceLines                any            `json:"sentence_lines"`
	Targeting                    map[string]any `json:"targeting"`
	TimeCreated                  string         `json:"time_created" datetime:"true"`
	TimeUpdated                  string         `json:"time_updated" datetime:"true"`
}

type SavedAudiencesResponseStruct struct {
	Data   []SavedAudience `json:"data"`
	Paging *Paging         `json:"paging"`
}

func (facebookClient *FacebookClient) ListSavedAudiences(ctx context.Context, page string) (items []SavedAudience, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(SavedAudience{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "saved_audiences")
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

	var responseStruct SavedAudiencesResponseStruct
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
