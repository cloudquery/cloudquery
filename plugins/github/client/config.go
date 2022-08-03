package client

type Config struct {
	AccessToken string   `yaml:"access_token"`
	Orgs        []string `yaml:"orgs"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c Config) Example() string {
	return ""
}
