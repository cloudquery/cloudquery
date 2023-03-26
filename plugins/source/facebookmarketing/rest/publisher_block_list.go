package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type PublisherBlockList struct {
	// AppPublishers             any    `json:"app_publishers"`
	BusinessOwnerId           string `json:"business_owner_id"`
	Id                        string `json:"id"`
	IsAutoBlockingOn          *bool  `json:"is_auto_blocking_on"`
	IsEligibleAtCampaignLevel *bool  `json:"is_eligible_at_campaign_level"`
	LastUpdateTime            string `json:"last_update_time" datetime:"true"`
	LastUpdateUser            string `json:"last_update_user"`
	Name                      string `json:"name"`
	OwnerAdAccountId          string `json:"owner_ad_account_id"`
	// WebPublishers any `json:"web_publishers"`
}

type PublisherBlockListsResponseStruct struct {
	Data   []PublisherBlockList `json:"data"`
	Paging *Paging              `json:"paging"`
}

func (facebookClient *FacebookClient) ListPublisherBlockLists(ctx context.Context, page string) (items []PublisherBlockList, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(PublisherBlockList{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "publisher_block_lists")
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

	var responseStruct PublisherBlockListsResponseStruct
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
