package client

type Spec struct {
	Accounts []AccountSpec `json:"accounts,omitempty"`
}

type AccountSpec struct {
	Regions   []string `json:"regions,omitempty"`
	AccessKey string   `json:"access_key,omitempty"`
	SecretKey string   `json:"secret_key,omitempty"`
}
