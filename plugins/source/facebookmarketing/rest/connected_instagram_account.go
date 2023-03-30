package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type ConnnectedInstagramAccount struct {
	Biography                     string         `json:"biography"`
	BusinessDiscovery             map[string]any `json:"business_discovery"`
	FollowersCount                *int64         `json:"followers_count"`
	FollowsCount                  *int64         `json:"follows_count"`
	Id                            string         `json:"id"`
	IgId                          *int64         `json:"ig_id"`
	MediaCount                    *int64         `json:"media_count"`
	MentionedComment              map[string]any `json:"mentioned_comment"`
	MentionedMedia                map[string]any `json:"mentioned_media"`
	Name                          string         `json:"name"`
	OwnerBusiness                 map[string]any `json:"owner_business"`
	ProfilePictureUrl             string         `json:"profile_picture_url"`
	ShoppingProductTagEligibility *bool          `json:"shopping_product_tag_eligibility"`
	ShoppingReviewStatus          string         `json:"shopping_review_status"`
	Username                      string         `json:"username"`
	Website                       string         `json:"website"`
}

type ConnnectedInstagramAccountsResponseStruct struct {
	Data   []ConnnectedInstagramAccount `json:"data"`
	Paging *Paging                      `json:"paging"`
}

func (facebookClient *FacebookClient) ListConnectedInstagramAccounts(ctx context.Context, page string) (items []ConnnectedInstagramAccount, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(ConnnectedInstagramAccount{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "connected_instagram_accounts")
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

	var responseStruct ConnnectedInstagramAccountsResponseStruct
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
