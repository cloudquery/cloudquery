package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Advideo struct {
	AdBreaks []int64 `json:"ad_breaks"`
	// AdminCreator map[string]any `json:"admin_creator"` // nonexistant
	// AudioIsrc                 any            `json:"audio_isrc"` copyrighted - skipping
	BackdatedTime            string   `json:"backdated_time" datetime:"true"`
	BackdatedTimeGranularity string   `json:"backdated_time_granularity"`
	ContentCategory          string   `json:"content_category"`
	ContentTags              []string `json:"content_tags"`
	// Copyright                 map[string]any `json:"copyright"`
	// CopyrightMonitoringStatus string         `json:"copyright_monitoring_status"`
	CreatedTime            string         `json:"created_time" datetime:"true"`
	CustomLabels           []string       `json:"custom_labels"`
	Description            string         `json:"description"`
	EmbedHtml              string         `json:"embed_html"`
	Embeddable             *bool          `json:"embeddable"`
	Event                  map[string]any `json:"event"`
	Expiration             any            `json:"expiration"`
	Format                 any            `json:"format"`
	From                   any            `json:"from"`
	Icon                   string         `json:"icon"`
	Id                     string         `json:"id"`
	IsCrosspostVideo       *bool          `json:"is_crosspost_video"`
	IsCrosspostingEligible *bool          `json:"is_crossposting_eligible"`
	IsEpisode              *bool          `json:"is_episode"`
	IsInstagramEligible    *bool          `json:"is_instagram_eligible"`
	// IsReferenceOnly        *bool          `json:"is_reference_only"` copyrighted
	Length            *float64 `json:"length"`
	LiveAudienceCount *int64   `json:"live_audience_count"`
	LiveStatus        string   `json:"live_status"`
	// MusicVideoCopyright      map[string]any `json:"music_video_copyright"`
	PermalinkUrl             string         `json:"permalink_url"`
	Picture                  string         `json:"picture"`
	Place                    map[string]any `json:"place"`
	PostViews                *int64         `json:"post_views"`
	PremiereLivingRoomStatus string         `json:"premiere_living_room_status"`
	Privacy                  map[string]any `json:"privacy"`
	Published                *bool          `json:"published"`
	ScheduledPublishTime     string         `json:"scheduled_publish_time" datetime:"true"`
	Source                   string         `json:"source"`
	Spherical                *bool          `json:"spherical"`
	Status                   map[string]any `json:"status"`
	Title                    string         `json:"title"`
	UniversalVideoId         string         `json:"universal_video_id"`
	UpdatedTime              string         `json:"updated_time" datetime:"true"`
	Views                    *int64         `json:"views"`
}

type AdvideosResponseStruct struct {
	Data   []Advideo `json:"data"`
	Paging *Paging   `json:"paging"`
}

func (facebookClient *FacebookClient) ListAdvideos(ctx context.Context, page string) (items []Advideo, nextPage string, err error) {
	query := url.Values{}

	query.Set("fields", strings.Join(getAllFieldJsonTags(Advideo{}), ","))

	query.Set("access_token", facebookClient.AccessToken)

	if page != "" {
		query.Set("after", page)
	}

	path, err := url.JoinPath("v16.0", "act_"+facebookClient.AdAccountId, "advideos")
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

	var responseStruct AdvideosResponseStruct
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
