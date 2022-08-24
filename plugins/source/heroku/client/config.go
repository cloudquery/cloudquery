package client

type Config struct {
	Token string   `yaml:"token"`
	Teams []string `yaml:"teams"`
}

func NewConfig() *Config {
	return &Config{}
}

func (Config) Example() string {
	return `
		token: <Token HERE>
		teams: ["your_company"]
`
}
