package client

type Config struct {
	AccessToken string   `yaml:"access_token"`
	Orgs        []string `yaml:"orgs"`
}

func NewConfig() *Config {
	return &Config{}
}

func (Config) Example() string {
	return `
		access_token: <Access Token HERE>
		orgs: ["cloudquery"]
`
}
