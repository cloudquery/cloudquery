package client

type Spec struct {
	RegionID    string `json:"region_id,omitempty"`
	AccessKey   string `json:"access_key,omitempty"`
	SecretKey   string `json:"secret_key,omitempty"`
	BillHistory int    `json:"bill_history,omitempty"`
}
