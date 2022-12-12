package client

type Spec struct {
	AccessToken string `json:"access_token,omitempty"`
	TeamID      string `json:"team_id,omitempty"`

	EndpointURL string `json:"endpoint_url,omitempty"`
	Timeout     int64  `json:"timeout_secs,omitempty"`
}
