package client

import (
	"context"

	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	Spec   Spec
	Notion *NotionClient
}

func (Client) ID() string {
	return "notion"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (Client, error) {
	s.SetDefaults() // Sets the default value of Notion Version.
	bearerToken := s.BearerToken
	notionVersion := s.NotionVersion
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
