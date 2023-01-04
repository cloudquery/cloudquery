package client

type Spec struct {
	FastlyAPIKey string   `json:"fastly_api_key"`
	Services     []string `json:"services"`
}
