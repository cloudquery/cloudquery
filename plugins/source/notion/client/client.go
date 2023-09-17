package client

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	Spec   Spec
	Notion *NotionClient
}

func (c *Client) ID() string {
	return "notion"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (Client, error) {

	bearerToken := os.Getenv("NOTION_SECRET_KEY")
	notionVersion := os.Getenv("NOTION_VERSION")
	c, err := NewNotionClient(bearerToken, notionVersion)
	if err != nil {
		return Client{}, err
	}

	return Client{
		logger: logger,
		Spec:   *s,
		Notion: c,
	}, nil
}
