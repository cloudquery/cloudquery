package cmd

import "github.com/rudderlabs/analytics-go"

const (
	rudderStackWritekey = "2Fz1JvndsXs5WjHQ03fSYNv4gcp"
	rudderStackEndpoint = "https://cloudquerypgm.dataplane.rudderstack.com"
)

type Analytics struct {
	client analytics.Client
}

func initAnalytics() *Analytics {
	c := Analytics{
		client: analytics.New(rudderStackWritekey, rudderStackEndpoint),
	}
	return &c
}

func (a *Analytics) Enqueue(event analytics.Message) error {
	if a.client != nil {
		return a.client.Enqueue(event)
	}
	return nil
}

func (a *Analytics) Close() error {
	if a.client != nil {
		return a.client.Close()
	}
	return nil
}
