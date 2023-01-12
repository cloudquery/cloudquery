package client

type Spec struct {
	AccessToken string   `json:"access_token,omitempty"`
	TeamIDs     []string `json:"team_ids,omitempty"`

	EndpointURL string `json:"endpoint_url,omitempty"`
	Timeout     int64  `json:"timeout_secs,omitempty"`
	PageSize    int64  `json:"page_size,omitempty"`
	MaxRetries  int64  `json:"max_retries,omitempty"`
	MaxWait     int64  `json:"max_wait_secs,omitempty"`
}
