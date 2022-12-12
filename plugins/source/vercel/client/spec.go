package client

type Spec struct {
	AccessToken string   `json:"access_token,omitempty"`
	TeamIDs     []string `json:"team_ids,omitempty"`

	EndpointURL string `json:"endpoint_url,omitempty"`
	Timeout     int64  `json:"timeout_secs,omitempty"`
}
