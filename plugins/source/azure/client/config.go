package client

type Config struct {
	Subscriptions []string `hcl:"subscriptions,optional"`
}
