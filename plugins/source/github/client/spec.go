package client

type Spec struct {
	AccessToken string   `yaml:"access_token"`
	Orgs        []string `yaml:"orgs"`
}
