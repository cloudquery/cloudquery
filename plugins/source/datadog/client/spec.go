package client

type Spec struct {
	Accounts    []Account `json:"accounts"`
	Concurrency int       `json:"concurrency"`
	Site        string    `json:"site"`
}

type Account struct {
	Name   string `json:"name"`
	APIKey string `json:"api_key"`
	AppKey string `json:"app_key"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = 10000
	}
}
