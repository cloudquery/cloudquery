package client

type Client struct {
	Tailnet string
	APIKey  string
	Clients map[string]Interface
}

func (c *Client) WithTailnet(tailnet string) *Client {
	client := *c
	client.Tailnet = tailnet
	return &client
}
