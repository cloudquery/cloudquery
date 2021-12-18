package client

// Provider Configuration

type Config struct {
	// here goes top level configuration for your provider
	// This object will be pass filled in depending on user's configuration
	// CHANGEME
	ExampleConfig string `hcl:"example_config"`

	// resources that user asked to fetch
	// each resource can have optional additional configurations
	Resources []struct {
		Name  string
		Other map[string]interface{} `hcl:",inline"`
	}
}

func (c Config) Example() string {
	return `configuration {
	// CHANGEME:
	//Here you define your default/example documentation.
	//That is generated with cloudquery init YourProviderName
	// Optional or required parameters
	// debug = false
	// api_key = ""	
}
`
}
