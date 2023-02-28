package client

type Spec struct {
	Subscriptions []string `json:"subscriptions"`
	CloudName     string   `json:"cloud_name"`
}
