package client

type Config struct {
	Token string `yaml:"token"`
}

func NewConfig() *Config {
	return &Config{}
}

func (Config) Example() string {
	return `
		token: <Token HERE>
`
}
