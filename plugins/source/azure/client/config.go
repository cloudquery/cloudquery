package client

// Provider Configuration

type Config struct {
	Subscriptions []string `hcl:"subscriptions,optional"`
}

func (Config) Example() string {
	return `
Optional. if you not specified, cloudquery tries to access all subscriptions available to tenant
subscriptions:
  - "<YOUR_SUBSCRIPTION_ID_HERE>"
`
}
