package client

type Client struct {
  logger zerolog.Logger
}

func (c *Client) Logger() *zerolog.Logger {
  return &c.logger
}

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
  c := Client{
    logger: logger,
  }
  return &c, fmt.Errorf("not implemented")
}