package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Adcreative struct {
	AccountId                      string         `json:"account_id"`
	ActorId                        string         `json:"actor_id"`
	Adlabels                       []any          `json:"adlabels"`
	ApplinkTreatment               string         `json:"applink_treatment"`
	AssetFeedSpec                  map[string]any `json:"asset_feed_spec"`
	AuthorizationCategory          string         `json:"authorization_category"`
	AutoUpdate                     *bool          `json:"auto_update"`
	Body                           string         `json:"body"`
	BrandedContentSponsorPageId    string         `json:"branded_content_sponsor_page_id"`
	BundleFolderId                 string         `json:"bundle_folder_id"`
	CallToActionType               string         `json:"call_to_action_type"`
	CategorizationCriteria         string         `json:"categorization_criteria"`
	CategoryMediaSource            string         `json:"category_media_source"`
	CollaborativeAdsLsbImageBankId string         `json:"collaborative_ads_lsb_image_bank_id"`
	DegreesOfFreedomSpec           map[string]any `json:"degrees_of_freedom_spec"`
	DestinationSetId               string         `json:"destination_set_id"`
	DynamicAdVoice                 string         `json:"dynamic_ad_voice"`
	EffectiveAuthorizationCategory string         `json:"effective_authorization_category"`
	EffectiveInstagramMediaId      string         `json:"effective_instagram_media_id"`
	EffectiveInstagramStoryId      string         `json:"effective_instagram_story_id"`
	EffectiveObjectStoryId         string         `json:"effective_object_story_id"`
	EnableDirectInstall            *bool          `json:"enable_direct_install"`
	EnableLaunchInstantApp         *bool          `json:"enable_launch_instant_app"`
	Id                             string         `json:"id"`
	ImageCrops                     map[string]any `json:"image_crops"`
	ImageHash                      string         `json:"image_hash"`
	ImageUrl                       string         `json:"image_url"`
	InstagramActorId               string         `json:"instagram_actor_id"`
	InstagramPermalinkUrl          string         `json:"instagram_permalink_url"`
	InstagramStoryId               string         `json:"instagram_story_id"`
	InstagramUserId                string         `json:"instagram_user_id"`
	InteractiveComponentsSpec      map[string]any `json:"interactive_components_spec"`
	LinkDeepLinkUrl                string         `json:"link_deep_link_url"`
	LinkDestinationDisplayUrl      string         `json:"link_destination_display_url"`
	LinkOgId                       string         `json:"link_og_id"`
	LinkUrl                        string         `json:"link_url"`
	MessengerSponsoredMessage      string         `json:"messenger_sponsored_message"`
	Name                           string         `json:"name"`
	ObjectId                       string         `json:"object_id"`
	ObjectStoreUrl                 string         `json:"object_store_url"`
	ObjectStoryId                  string         `json:"object_story_id"`
	ObjectStorySpec                map[string]any `json:"object_story_spec"`
	ObjectType                     string         `json:"object_type"`
	ObjectUrl                      string         `json:"object_url"`
	OmnichannelLinkSpec            map[string]any `json:"omnichannel_link_spec"`
	PlacePageSetId                 string         `json:"place_page_set_id"`
	PlatformCustomizations         map[string]any `json:"platform_customizations"`
	PlayableAssetId                string         `json:"playable_asset_id"`
	PortraitCustomizations         map[string]any `json:"portrait_customizations"`
	ProductSetId                   string         `json:"product_set_id"`
	RecommenderSettings            map[string]any `json:"recommender_settings"`
	SourceInstagramMediaId         string         `json:"source_instagram_media_id"`
	Status                         string         `json:"status"`
	TemplateUrl                    string         `json:"template_url"`
	TemplateUrlSpec                map[string]any `json:"template_url_spec"`
	ThumbnailId                    string         `json:"thumbnail_id"`
	ThumbnailUrl                   string         `json:"thumbnail_url"`
	Title                          string         `json:"title"`
	UrlTags                        string         `json:"url_tags"`
	UsePageActorOverride           *bool          `json:"use_page_actor_override"`
	VideoId                        string         `json:"video_id"`
}

type AdcreativesResponseStruct struct {
	Data   []Adcreative `json:"data"`
	Paging *Paging      `json:"paging"`
}

func (facebookClient *FacebookClient) ListAdcreatives(ctx context.Context, page string) (items []Adcreative, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(Adcreative{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "adcreatives")
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

	var responseStruct AdcreativesResponseStruct
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
