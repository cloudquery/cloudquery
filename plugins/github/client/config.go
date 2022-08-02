package client

type Config struct {
	// here goes top level configuration for your provider
	// This object will be pass filled in depending on user's configuration
	// CHANGEME
	AccessToken string `yaml:"access_token"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c Config) Example() string {
	return ""
}
