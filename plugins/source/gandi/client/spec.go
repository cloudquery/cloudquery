package client

type Spec struct {
	APIKey     string `json:"api_key,omitempty"`
	SharingID  string `json:"sharing_id,omitempty"`
	GandiDebug bool   `json:"gandi_debug,omitempty"`

	EndpointURL string `json:"endpoint_url,omitempty"`
	Timeout     int64  `json:"timeout_secs,omitempty"`
}
