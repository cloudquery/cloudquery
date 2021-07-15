package client

// Provider Configuration

type Config struct {
	// here goes top level configuration for your provider
	// This object will be pass filled in depending on user's configuration
	ExampleConfig  bool      `yaml:"example_config"`

	// resources that user asked to fetch
	// each resource can have optional additional configurations
	Resources  []struct {
		Name  string
		Other map[string]interface{} `yaml:",inline"`
	}
}

